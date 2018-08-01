package protocol

import (
	"bytes"
	"io"

	"github.com/juju/errors"
)

func max(a uint64, b uint64) uint64 {
	if a >= b {
		return a
	} else {
		return b
	}
}

//https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::SSLRequest
func (c *Session) ReadHandshakeV10() error {
	//read inithandshake packet
	buffer, err := c.ReadPacket()
	if err != nil {
		return errors.Trace(err)
	}
	payload := NewPayLoad(buffer[0:])

	//read protocol version 1 bytes
	protocol_version := payload.ReadInt1()
	if protocol_version != 0x0a {
		return errors.Errorf("invalid protocol version: %d ", protocol_version)
	}

	//read server version string.null
	_ = payload.ReadStringWithNull()

	//read connect id 4 bytes
	c.ConnectionId = payload.ReadInt4()

	//read auth-plugin-data-part-1 8 bytes
	c.AuthPluginData = payload.ReadStringWithFixLen(8)

	//skip filler 1 bytes
	_ = payload.ReadInt1()

	//read low capability flags (lower 2 bytes)
	lflag := payload.ReadInt2()

	//if no more data in the packet
	if payload.IsEOF() {
		return nil
	}

	//read character set 1 bytes
	c.Character = payload.ReadInt1()

	//read status
	_ = payload.ReadInt2()

	//read upper capability flags (upper 2 bytes)
	uflag := payload.ReadInt2()

	//set Capabilities (4 bytes)
	c.Capabilities = uint32(uflag)<<16 | uint32(lflag)

	if c.Capabilities&CLIENT_PLUGIN_AUTH == 0 {
		return errors.Errorf("server does not support CLIENT_PLUGIN_AUTH")
	}

	//length of auth-plugin-data 1bytes
	auth_plugin_data_len := payload.ReadInt1()
	//skip reserved 10 bytes
	payload.ReadZero(10)

	//I find it does not work
	//https://dev.mysql.com/doc/internals/en/secure-password-authentication.html shows
	//client-side expects a 20-byte random challenge
	//if we use $len=MAX(13, length of auth-plugin-data - 8 as the length of auth_plugin_data_part2
	//the length is more than 21 instead of 20
	//so I decide to read only 12 bytes data as the challenge
	length := max(13, uint64(auth_plugin_data_len)-8)
	c.AuthPluginData = append(c.AuthPluginData, payload.ReadStringWithFixLen(12)...)
	//skip
	payload.ReadZero(length - 12)

	c.PluginName = payload.ReadStringWithNull()

	if string(c.PluginName) != "mysql_native_password" {
		return errors.Errorf("server does not use mysql_native_password as Authentication method!")
	}

	return nil
}

//https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::SSLRequest
func (c *Session) WriteHandshakeResponse41() error {
	//write a header four bytes
	length := 4

	//add
	capabilities := DefaultClientCapabilities | CLIENT_PLUGIN_AUTH
	length += 4

	//set max-packet size character set  reserved (all [0])
	length += 4 + 1 + 23

	//username+0x0
	length += len(c.Username) + 1

	//calculate auth-response
	passwd := CalAuthResponse(c.AuthPluginData, c.Passwd)
	if len(passwd) != 20 {
		return errors.Errorf("invalid scrambled password,looking forward 20 bytes instead of %d bytes", len(passwd))
	}
	length += 1 + 20

	//database
	if len(c.Database) != 0 {
		capabilities = capabilities & ^CLIENT_CONNECT_WITH_DB
	}
	length += len(c.Database) + 1

	//auth plugin name
	length += len(c.PluginName) + 1

	//create a payload
	buffer := make([]byte, length)
	payload := NewPayLoad(buffer)

	//write header
	payload.WriteZero(4)
	//write capability flags
	payload.WriteInt4(capabilities)
	//max-packet size
	payload.WriteZero(4)
	//character set
	payload.WriteInt1(DEFAULT_CHARSET)
	//reserved
	payload.WriteZero(23)
	//username
	payload.WriteStringWithNull([]byte(c.Username))
	//length of auth-response
	payload.WriteInt1(uint8(len(passwd)))
	//auth-response
	payload.WriteStringWithFixLen(passwd)
	//database
	if len(c.Database) != 0 {
		payload.WriteStringWithNull([]byte(c.Database))
	} else {
		payload.WriteZero(0)
	}
	//auth plugin name
	payload.WriteStringWithNull(c.PluginName)

	return c.WritePacket(payload.Bytes())

}

//ReadOkPacket implements to read ok packet
func (c *Session) ReadOkPacket() error {
	//read inithandshake packet
	buffer, err := c.ReadPacket()
	if err != nil {
		return errors.Trace(err)
	}

	payload := NewPayLoad(buffer[:])

	if header := payload.ReadInt1(); header == 0x00 || header == 0xfe {
		return nil
	} else if header == 0xff {
		//read error code
		errorCode := payload.ReadInt2()
		//skip sql_state_marker and sql_state
		payload.ReadZero(6)
		//read error message
		errorMessage := payload.ReadStringWithEOF()
		//close the connection
		c.Close()
		return errors.Errorf("Error[%d]:%s", errorCode, string(errorMessage))
	} else {
		return errors.Errorf("Not a recognized ok packet.")
	}
}
