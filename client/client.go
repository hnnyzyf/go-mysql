package client

import (
	"crypto/tls"
	"net"

	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/charset"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/hnnyzyf/go-mysql/util/math"
	"github.com/juju/errors"
	"golang.org/x/text/encoding"
)

var errHandShake10 = errors.Errorf("Client:not a valid HandShake10 packets")
var errResponsePacket = errors.Errorf("Client:not a valid response Packet")
var errERRPacket = errors.Errorf("Client:not a valid Err packet")

//session represent a connection to server from client
type session struct {
	//the real conn to server,the
	conn net.Conn

	//connection inforamtion
	username string
	password string
	database string

	//some options
	cfg *config

	//The sequence-id is incremented with each packet and may wrap around.
	//It starts at 0 and is reset to 0 when a new command begins in the Command Phase.
	sid uint8

	//capabilities flag use in client
	capabilities uint32

	//20-bytes random data from serve
	authPluginData []byte

	//authentication method
	authMethod protocol.Method

	//charset
	charset uint8
	//charset method
	decoder *encoding.Decoder
	encoder *encoding.Encoder
}

//connect create a tcp connection to server
func connect(host string, user string, passwd string, db string, cfg *config) (*session, error) {
	//the client connecting to the server with timeout
	conn, err := net.DialTimeout("tcp", host, cfg.timeout)
	if err != nil {
		return nil, errors.Trace(err)
	}

	//init session
	s := &session{
		conn:     conn,
		username: user,
		password: passwd,
		database: db,
		cfg:      cfg,
	}

	//init configuration
	if err := s.init(); err != nil {
		return nil, errors.Trace(err)
	}

	//the server responds with the Initial Handshake Packet
	if err := s.readHandShakeV10(); err != nil {
		return nil, errors.Trace(err)
	}

	if cfg.isSSL {
		//the client use ssl to sends the SSL Request Packet
		if err := s.writeSSLRequeset(); err != nil {
			return nil, errors.Trace(err)
		}
	}

	//the client sends the Handshake Response Packet
	if err := s.writeHandshakeResponse41(); err != nil {
		return nil, errors.Trace(err)
	}

	//read response packet and decide what to do according to the type of packet
	// 		OK-packet:return success
	//		ERR-packet:return false
	//		EOF-packet:return false
	packet, header, err := s.readResponsePacket()
	if err != nil {
		return nil, errors.Trace(err)
	}

	//check response packet
	switch header {
	case protocol.OK_Packet:
		return s, nil
	case protocol.ERR_Packet:
		if errMsg, err := parseErrPacket(packet); err != nil {
			return nil, errors.Trace(err)
		} else {
			return nil, errors.Errorf(errMsg.msg)
		}
	default:
		return nil, errors.Errorf("Client:fail to accept correct generic response packet from server!")
	}

	return s, nil
}

//Connect is a wrapper of connect
func Connect(host string, user string, passwd string, db string) (*session, error) {
	return connect(host, user, passwd, db, defaultConfig)
}

//ConnectWithConfig is a wrapper of connect
func ConnectWithConfig(host string, user string, passwd string, db string, cfg *config) (*session, error) {
	return connect(host, user, passwd, db, cfg)
}

//init read configuration from config
func (s *session) init() error {
	cfg := s.cfg

	//set capabilities
	s.capabilities = cfg.capabilities | protocol.CLIENT_CONNECT_ATTRS

	if len(s.database) == 0 {
		s.capabilities &= (^protocol.CLIENT_CONNECT_WITH_DB)
	} else {
		s.capabilities |= protocol.CLIENT_CONNECT_WITH_DB
	}

	if cfg.isSSL {
		s.capabilities |= protocol.CLIENT_SSL
	}

	if cfg.mutilStatement {
		s.capabilities |= protocol.CLIENT_MULTI_STATEMENTS
	}

	//test ssl
	if cfg.isSSL && cfg.tlsConfig == nil {
		return errors.Errorf("Client:not provide tls configuraions.")
	}

	//set charset method
	if method, err := charset.GetMethod(cfg.charset); err != nil {
		return errors.Trace(err)
	} else {
		s.decoder = method.NewDecoder()
		s.encoder = method.NewEncoder()
	}

	//set charset id
	if id, err := charset.GetID(cfg.charset); err != nil {
		return errors.Trace(err)
	} else {
		s.charset = id
	}

	return nil
}

//readHandShakeV10 Read the init handshake packet
func (s *session) readHandShakeV10() error {
	//Get packet from conn
	packet, err := protocol.ReadPacket(s.conn, s.sid)
	if err != nil {
		return errors.Trace(err)
	}

	//get payload from packet
	payload, err := packet.ReadPayload()
	if err != nil {
		return err
	}

	//read protocol version 1 bytes
	if protocolVersion, err := payload.ReadInt1(); err != nil {
		return errHandShake10
	} else {
		if !testProtocolVersion(protocolVersion) {
			return errors.Errorf("Client:not a valid protocol version: %d ", protocolVersion)
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
	} else {
		s.authPluginData = authPluginData
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
	if _, err := payload.ReadInt1(); err != nil {
		return errHandShake10
	}

	//read status flag
	if _, err := payload.ReadInt2(); err != nil {
		return errHandShake10
	}

	//read upper capability flags (upper 2 bytes)
	uflag, err := payload.ReadInt2()
	if err != nil {
		return errHandShake10
	}

	//set Capabilities (4 bytes)
	capabilities := uint32(uflag)<<16 | uint32(lflag)
	if s.cfg.isSSL && !testSSL(capabilities) {
		return errors.Errorf("Client:server does not support SSL connection!")
	}

	//The client should only announce the capabilities
	//in the Handshake Response Packet that it has in common with the server.
	s.capabilities &= capabilities

	//length of auth-plugin-data 1bytes
	var authPluginDataLen uint8
	if testPluginAuth(capabilities) {
		if length, err := payload.ReadInt1(); err != nil {
			return errHandShake10
		} else {
			authPluginDataLen = length
		}
	} else {
		if err := payload.ReadZero(1); err != nil {
			return errHandShake10
		} else {
			authPluginDataLen = 0
		}
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
	if testSecureConnection(capabilities) {
		length := math.MaxInt(13, int(authPluginDataLen)-8)
		lauthPluginData, err := payload.ReadStringWithFixLen(12)
		if err != nil {
			return errHandShake10
		}
		//skip
		if err := payload.ReadZero(length - 12); err != nil {
			return errHandShake10
		}
		s.authPluginData = append(s.authPluginData, lauthPluginData...)
	}

	//read pluginName and set authentication method
	var pluginName string
	if testPluginAuth(capabilities) {
		if name, err := payload.ReadStringWithNull(); err != nil {
			return errHandShake10
		} else {
			pluginName = hack.String(name)
		}
	} else {
		//https://dev.mysql.com/doc/internals/en/determining-authentication-method.html
		if !testProtocol41(capabilities) || !testSecureConnection(capabilities) {
			pluginName = "mysql_old_password"
		} else if testProtocol41(capabilities) && testSecureConnection(capabilities) {
			pluginName = "mysql_native_password"
		} else {
			return errors.Errorf("Client:server do not support any authentication method")
		}
	}
	if method, err := protocol.GetMethod(pluginName); err != nil {
		return errors.Trace(err)
	} else {
		s.authMethod = method
	}

	return nil
}

//writeSSLRequeset create ssl communication to server
func (s *session) writeSSLRequeset() error {
	length := 4 + 4 + 1 + 23

	//create a payload
	payload := binary.NewBuffer(make([]byte, length))

	//write capability
	if err := payload.WriteInt4(s.capabilities); err != nil {
		return errors.Trace(err)
	}

	//write max_packet size
	if err := payload.WriteInt4(protocol.PayLoadMaxLen); err != nil {
		return errors.Trace(err)
	}

	//write chracter set
	if err := payload.WriteInt1(s.charset); err != nil {
		return errors.Trace(err)
	}

	//create packet
	packet := protocol.NewPacket(length, s.sid, payload)

	//write packet
	if sid, err := protocol.WritePacket(s.conn, packet); err != nil {
		return errors.Trace(err)
	} else {
		s.sid = sid
	}

	//replace the origin connection to ssl connection
	s.conn = tls.Client(s.conn, s.cfg.tlsConfig)

	return nil
}

//https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::SSLRequest
//writeHandshakeResponse41 return a HandshakeResponse to server
func (s *session) writeHandshakeResponse41() error {
	//write a header four bytes
	length := 0
	//add capabilities
	length += 4

	//set max-packet size character set  reserved (all [0])
	length += 4 + 1 + 23

	//username+0x0 and convert database from utf8 to the charset
	var username []byte
	if name, err := s.decoder.Bytes(hack.Slice(s.username)); err != nil {
		return errors.Trace(err)
	} else {
		username = name
	}
	length += len(username) + 1

	//calculate auth-response
	password := s.authMethod.EncodeKey(s.authPluginData, s.password)
	length += len(password)
	if testLenencData(s.capabilities) {
		length += binary.LengthOfInteger(uint64(len(password)))
	} else if testSecureConnection(s.capabilities) {
		length += 1
	} else {
		length += 1
	}

	//set database name
	var database []byte
	if testConnectWithDB(s.capabilities) {
		//convert database from utf8 to the charset
		if db, err := s.decoder.Bytes(hack.Slice(s.database)); err != nil {
			return errors.Trace(err)
		} else {
			database = db
		}
		length += len(database) + 1
	}

	//auth plugin name
	if testPluginAuth(s.capabilities) {
		length += len(s.authMethod.Name()) + 1
	}

	//add attributes
	var key []byte
	var value []byte
	if testConnectAttrs(s.capabilities) {
		key = hack.Slice("AUTOCOMMIT")
		if s.cfg.autoCommit {
			value = hack.Slice("0")
		} else {
			value = hack.Slice("1")
		}
		length += binary.LengthOfInteger(uint64(len(key) + len(value)))
		length += len(key)
		length += len(value)
	}

	//create a payload
	payload := binary.NewBuffer(make([]byte, length))

	//write capability flags
	if err := payload.WriteInt4(s.capabilities); err != nil {
		return errors.Trace(err)
	}

	//max-packet size
	if err := payload.WriteInt4(protocol.PayLoadMaxLen); err != nil {
		return errors.Trace(err)
	}

	//character set
	if err := payload.WriteInt1(s.charset); err != nil {
		return errors.Trace(err)
	}

	//reserved
	if err := payload.WriteZero(23); err != nil {
		return errors.Trace(err)
	}

	//username
	if err := payload.WriteStringWithNull(username); err != nil {
		return errors.Trace(err)
	}

	//write auth-response
	if testLenencData(s.capabilities) {
		if _, err := payload.WriteLengthEncodedInteger(uint64(len(password))); err != nil {
			return errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(password); err != nil {
			return errors.Trace(err)
		}
	} else if testSecureConnection(s.capabilities) {
		if err := payload.WriteInt1(uint8(len(password))); err != nil {
			return errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(password); err != nil {
			return errors.Trace(err)
		}
	} else {
		if err := payload.WriteStringWithNull(password); err != nil {
			return errors.Trace(err)
		}
	}

	//database
	if testConnectWithDB(s.capabilities) {
		if err := payload.WriteStringWithNull(database); err != nil {
			return errors.Trace(err)
		}
	}

	//auth plugin name
	if testPluginAuth(s.capabilities) {
		if err := payload.WriteStringWithNull(hack.Slice(s.authMethod.Name())); err != nil {
			return errors.Trace(err)
		}
	}

	//set attribute
	if testConnectAttrs(s.capabilities) {
		//write length of all key-values
		n := binary.LengthOfInteger(uint64(len(key) + len(value)))
		if _, err := payload.WriteLengthEncodedInteger(uint64(n)); err != nil {
			return errors.Trace(err)
		}
		//write key
		if err := payload.WriteStringWithFixLen(key); err != nil {
			return errors.Trace(err)
		}
		//write value
		if err := payload.WriteStringWithFixLen(value); err != nil {
			return errors.Trace(err)
		}
	}

	//create packet
	packet := protocol.NewPacket(length, s.sid, payload)

	//write packet
	if sid, err := protocol.WritePacket(s.conn, packet); err != nil {
		return errors.Trace(err)
	} else {
		s.sid = sid
	}

	return nil
}

//readResponsePacket implements to read generic response packet
func (s *session) readResponsePacket() (*protocol.Packet, uint8, error) {
	//Get packet from conn
	packet, err := protocol.ReadPacket(s.conn, s.sid+1)
	if err != nil {
		return nil, protocol.Unknown_Packet, errors.Trace(err)
	} else {
		//id should be equal to response packet
		s.sid = packet.ReadId()
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
		return nil, protocol.Unknown_Packet, errors.Trace(err)
	}

	//OK: header = 0 and length of packet > 7
	//EOF: header = 0xfe and length of packet < 9
	if header == protocol.OK_Packet && length > 7-4 {
		return packet, header, nil
	} else if header == protocol.EOF_Packet && length < 9-4 {
		return packet, header, nil
	} else if header == protocol.ERR_Packet {
		return packet, header, nil
	} else {
		return nil, header, errResponsePacket
	}

}

func (s *session) Close() {
	s.conn.Close()
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
