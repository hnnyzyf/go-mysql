package client

import (
	"net"

	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/hnnyzyf/go-mysql/util/math"
	"github.com/juju/errors"
)

var errHandShake10 = errors.Errorf("Client：Not a valid HandShake10 packets")
var errResponsePacket = errors.Errorf("Client:Not a valid response Packet")
var errERRPacket = errors.Errorf("Client:Not a valid Err packet")

//session represent a connection to server from client
type session struct {
	//the real conn to server
	c *net.TCPConn

	//connection inforamtion
	username string
	password string
	database string

	//some options
	cfg *Config

	//The sequence-id is incremented with each packet and may wrap around.
	//It starts at 0 and is reset to 0 when a new command begins in the Command Phase.
	sid uint8

	//auth represent the authentication method
	method protocol.Method

	//20-bytes random data from serve
	authPluginData []byte

	//the name of authentication method()
	pluginName []byte

	//The capability flags are used by the client and server
	//to indicate which features they support and want to use.
	capabilities uint32

	//the character of server
	character uint8
}

//Connect create a tcp connection to server
func Connect(host string, user string, passwd string, db string, cfg *Config) (*session, error) {
	//the client connecting to the server
	raddr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return nil, errors.Trace(err)
	}

	//connect to raddr
	c, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return nil, errors.Trace(err)
	}

	conn := &session{
		c:        c,
		username: user,
		password: passwd,
		database: db,
		cfg:      cfg,
	}

	//the server responds with the Initial Handshake Packet
	if err = conn.ReadHandshakeV10(); err != nil {
		return nil, errors.Trace(err)
	}

	//the client sends the Handshake Response Packet
	if err = conn.WriteHandshakeResponse41(); err != nil {
		return nil, errors.Trace(err)
	}

	//read response packet and decide what to do according to the type of packet
	// 		OK-packet:return success
	//		ERR-packet:return false
	//		EOF-packet:return false
	packet, header, err := conn.ReadResponsePacket()
	if err != nil {
		return nil, errors.Trace(err)
	}

	//check response packet
	switch header {
	case protocol.OK_Packet:
		return conn, nil
	case protocol.ERR_Packet:
		if errMsg, err := parseErrPacket(packet); err != nil {
			return nil, errors.Trace(err)
		} else {
			return nil, errors.Errorf(errMsg.msg)
		}
	default:
		return nil, errors.Errorf("Fail to accept correct generic response packet from server!")
	}

	return conn, nil

}

//ReadHandshakeV10 Read the init handshake packet
func (s *session) ReadHandshakeV10() error {
	//Get packet from conn
	packet, err := protocol.ReadPacket(s.c, s.sid)
	if err != nil {
		return errors.Trace(err)
	}

	//get payload from packet
	payload, err := packet.ReadPayload()
	if err != nil {
		return err
	}

	//read protocol version 1 bytes
	if protocol_version, err := payload.ReadInt1(); err != nil {
		return errHandShake10
	} else {
		if protocol_version != 0x0a {
			return errors.Errorf("Client:invalid protocol version: %d ", protocol_version)
		}
	}
	//read server version string.null
	if _, err := payload.ReadStringWithNull(); err != nil {
		return errHandShake10
	}

	//read connect id 4 bytes
	if _, err := payload.ReadInt4(); err != nil {
		return errHandShake10
	}

	//read auth-plugin-data-part-1 8 bytes
	authPluginData, err := payload.ReadStringWithFixLen(8)
	if err != nil {
		return errHandShake10
	}

	//skip filler 1 bytes
	if _, err := payload.ReadInt1(); err != nil {
		return errHandShake10
	}

	//read low capability flags (lower 2 bytes)
	lflag, err := payload.ReadInt2()
	if err != nil {
		return errHandShake10
	}

	//if no more data in the packet
	if payload.IsEOF() {
		return nil
	}

	//read character set 1 bytes
	s.character, err = payload.ReadInt1()
	if err != nil {
		return errHandShake10
	}
	//read status
	if _, err := payload.ReadInt2(); err != nil {
		return errHandShake10
	}

	//read upper capability flags (upper 2 bytes)
	uflag, err := payload.ReadInt2()
	if err != nil {
		return errHandShake10
	}

	//set Capabilities (4 bytes)
	s.capabilities = uint32(uflag)<<16 | uint32(lflag)
	if s.capabilities&protocol.CLIENT_PLUGIN_AUTH == 0 {
		return errors.Errorf("Server does not support CLIENT_PLUGIN_AUTH!")
	}

	if s.capabilities&protocol.CLIENT_PROTOCOL_41 == 0 {
		return errors.Errorf("The server seems too old and does not support CLIENT_PROTOCOL_41!")
	}

	if s.capabilities&protocol.CLIENT_SECURE_CONNECTION == 0 {
		return errors.Errorf("The server does not support CLIENT_SECURE_CONNECTION!")
	}

	//length of auth-plugin-data 1bytes
	authPluginDataLen, err := payload.ReadInt1()
	if err != nil {
		return errHandShake10
	}

	//skip reserved 10 bytes
	if err := payload.ReadZero(10); err != nil {
		return errHandShake10
	}

	//I find it does not work
	//https://dev.mysql.com/doc/internals/en/secure-password-authentication.html shows
	//client-side expects a 20-byte random challenge
	//if we use $len=MAX(13, length of auth-plugin-data - 8 as the length of auth_plugin_data_part2
	//the length is more than 21 instead of 20
	//so I decide to read only 12 bytes data as the challenge
	//maybe I should have a look at Mysql source code!?
	length := math.MaxInt(13, int(authPluginDataLen)-8)
	lauthPluginData, err := payload.ReadStringWithFixLen(12)
	if err != nil {
		return errHandShake10
	}
	s.authPluginData = append(authPluginData, lauthPluginData...)

	//skip
	if err := payload.ReadZero(length - 12); err != nil {
		return errHandShake10
	}

	//read pluginName and create authentication method
	s.pluginName, err = payload.ReadStringWithNull()
	if err != nil {
		return errHandShake10
	}

	method, err := protocol.NewMethod(hack.String(s.pluginName))
	if err != nil {
		return errors.Trace(err)
	}
	s.method = method

	return nil
}

//https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::SSLRequest
//WriteHandshakeResponse41 return a HandshakeResponse to server
func (s *session) WriteHandshakeResponse41() error {
	//write a header four bytes
	length := 0
	//add capabilities
	capabilities := protocol.DefaultClientCapabilities | protocol.CLIENT_PLUGIN_AUTH
	length += 4

	//set max-packet size character set  reserved (all [0])
	length += 4 + 1 + 23

	//username+0x0
	length += len(s.username) + 1

	//calculate auth-response
	ciphertext := s.method.EncodeKey(s.authPluginData, s.password)

	if len(ciphertext) != 20 {
		return errors.Errorf("invalid scrambled password,looking forward 20 bytes instead of %d bytes", len(ciphertext))
	}
	length += 1 + 20

	//database
	if len(s.database) != 0 {
		capabilities = capabilities & (^protocol.CLIENT_CONNECT_WITH_DB)
	}
	length += len(s.database) + 1

	//auth plugin name
	length += len(s.pluginName) + 1

	//create a payload
	payload := binary.NewBuffer(make([]byte, length))

	//write capability flags
	if err := payload.WriteInt4(capabilities); err != nil {
		return errors.Trace(err)
	}

	//max-packet size
	if err := payload.WriteZero(4); err != nil {
		return errors.Trace(err)
	}

	//character set
	if err := payload.WriteInt1(protocol.DefaultCharacter); err != nil {
		return errors.Trace(err)
	}

	//reserved
	if err := payload.WriteZero(23); err != nil {
		return errors.Trace(err)
	}

	//username
	if err := payload.WriteStringWithNull(hack.Slice(s.username)); err != nil {
		return errors.Trace(err)
	}

	//length of auth-response
	if err := payload.WriteInt1(uint8(len(ciphertext))); err != nil {
		return errors.Trace(err)
	}

	//auth-response
	if err := payload.WriteStringWithFixLen(ciphertext); err != nil {
		return errors.Trace(err)
	}

	//database
	if len(s.database) != 0 {
		if err := payload.WriteStringWithNull(hack.Slice(s.database)); err != nil {
			return errors.Trace(err)
		}
	} else {
		if err := payload.WriteZero(0); err != nil {
			return errors.Trace(err)
		}
	}

	//auth plugin name
	if err := payload.WriteStringWithNull(s.pluginName); err != nil {
		return errors.Trace(err)
	}

	//create packet
	packet := protocol.NewPacket(length, s.sid, payload)

	//write packet
	sid, err := protocol.WritePacket(s.c, packet)
	if err != nil {
		return errors.Trace(err)
	}

	//set sid
	s.sid = sid

	return nil
}

//ReadResponsePacket implements to read generic response packet
func (s *session) ReadResponsePacket() (*protocol.Packet, uint8, error) {
	//Get packet from conn
	packet, err := protocol.ReadPacket(s.c, s.sid)
	if err != nil {
		return nil, protocol.Unknown_Packet, errors.Trace(err)
	}

	//Read length
	length := packet.ReadLength()

	//Read payload
	payload, err := packet.ReadPayload()
	if err != nil {
		return nil, protocol.Unknown_Packet, errors.Trace(err)
	}

	//Read Header
	header, err := payload.ReadInt1()
	if err != nil {
		return nil, protocol.Unknown_Packet, errResponsePacket
	}

	if header == protocol.OK_Packet && length > 7 {
		return packet, header, nil
	} else if header == protocol.EOF_Packet && length < 9 {
		return packet, header, nil
	} else if header == protocol.ERR_Packet {
		return packet, header, nil
	} else {
		return nil, header, errResponsePacket
	}

}

func (s *session) Close() {
	s.c.Close()
}

type errMsg struct {
	code  uint16 //error-code
	state []byte //	SQL State
	msg   string //human readable error message
}

//parseErrPacket parse the payload of errPacket
func parseErrPacket(packet *protocol.Packet) (*errMsg, error) {
	payload, err := packet.ReadPayload()
	if err != nil {
		return nil, errors.Trace(err)
	}

	//header
	if _, err := payload.ReadInt1(); err != nil {
		return nil, errERRPacket
	}

	//code
	code, err := payload.ReadInt2()
	if err != nil {
		return nil, errERRPacket
	}

	//marker
	if _, err = payload.ReadStringWithFixLen(1); err != nil {
		return nil, errERRPacket
	}

	//SQL State
	state, err := payload.ReadStringWithFixLen(5)
	if err != nil {
		return nil, errERRPacket
	}

	//human readable error message
	msg, err := payload.ReadStringWithEOF()
	if err != nil {
		return nil, errERRPacket
	}

	return &errMsg{
		code:  code,
		state: state,
		msg:   hack.String(msg),
	}, nil
}
