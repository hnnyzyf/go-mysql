package client

import (
	"net"

	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

//session represent a connection to server from client
type session struct {
	//the real conn to server
	c *net.TCPConn

	//connection inforamtion
	username string
	password string
	database string

	//The sequence-id is incremented with each packet and may wrap around.
	//It starts at 0 and is reset to 0 when a new command begins in the Command Phase.
	sid uint8

	//20-bytes random data from serve
	authPluginData []byte

	//The capability flags are used by the client and server
	//to indicate which features they support and want to use.
	capabilities uint32

	//auth represent the authentication method
	method protocol.Method
}

//Connect create a tcp connection to server
func Connect(addr string, user string, passwd string, db string) (*session, error) {
	//the client connecting to the server
	raddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, errors.Trace(err)
	}

	//connect to raddr
	c, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return nil, errors.Trace(err)
	}

	conn := &session{c, username, passwd, db}

	//the server responds with the Initial Handshake Packet
	if err = conn.ReadHandshakeV10(); err != nil {
		return nil, errors.Trace(err)
	}

	//the client sends the Handshake Response Packet
	if err = conn.WriteHandshakeResponse41(); err != nil {
		return nil, errors.Trace(err)
	}

	//read ok packet
	if err := conn.ReadOkPacket(); err != nil {
		return nil, errors.Trace(err)
	}

	return conn, nil

}

//ReadHandshakeV10 Read the init handshake packet
func (s *session) ReadHandshakeV10() error {
	//Get packet from conn
	p, err := protocol.ReadPacket(s.c)
	if err != nil {
		return errors.Trace(err)
	}

	//check sequenced id
	sid := p.ReadId()
	if sid != s.sid {
		return errors.Errorf("invalid sequence id %d,expect %d", sequenceId, s.sid)
	}

	//get payload from packet
	payload := p.ReadPayload()

	//read protocol version 1 bytes
	protocol_version := payload.ReadInt1()
	if protocol_version != 0x0a {
		return errors.Errorf("invalid protocol version: %d ", protocol_version)
	}

	//read server version string.null
	_ = payload.ReadStringWithNull()

	//read connect id 4 bytes
	connectionId := payload.ReadInt4()

	//read auth-plugin-data-part-1 8 bytes
	authPluginData := payload.ReadStringWithFixLen(8)

	//skip filler 1 bytes
	_ = payload.ReadInt1()

	//read low capability flags (lower 2 bytes)
	lflag := payload.ReadInt2()

	//if no more data in the packet
	if payload.IsEOF() {
		return nil
	}

	//read character set 1 bytes
	character := payload.ReadInt1()

	//read status
	_ = payload.ReadInt2()

	//read upper capability flags (upper 2 bytes)
	uflag := payload.ReadInt2()

	//set Capabilities (4 bytes)
	s.capabilities = uint32(uflag)<<16 | uint32(lflag)
	if s.capabilities&protocol.CLIENT_PLUGIN_AUTH == 0 {
		return errors.Errorf("server does not support CLIENT_PLUGIN_AUTH")
	}

	//length of auth-plugin-data 1bytes
	authPluginDataLen := payload.ReadInt1()

	//skip reserved 10 bytes
	payload.ReadZero(10)

	//I find it does not work
	//https://dev.mysql.com/doc/internals/en/secure-password-authentication.html shows
	//client-side expects a 20-byte random challenge
	//if we use $len=MAX(13, length of auth-plugin-data - 8 as the length of auth_plugin_data_part2
	//the length is more than 21 instead of 20
	//so I decide to read only 12 bytes data as the challenge
	length := max(13, uint64(authPluginDataLen)-8)
	s.authPluginData = append(authPluginData, payload.ReadStringWithFixLen(12)...)

	//skip
	payload.ReadZero(length - 12)

	//read pluginName and create authentication method
	pluginName = payload.ReadStringWithNull()
	method, err := protocol.NewMethod(hack.String(pluginName))
	if err != nil {
		return errors.Trace(err)
	}
	s.method = method

	return nil
}
