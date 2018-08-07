package client

import (
	"github.com/hnnyzyf/go-mysql/protocol"
)

//testProtocolVersion tests whether  the server supprot v10 protocol
func testProtocolVersion(version uint8) bool {
	return version != 0x0A
}

//testPluginAuth tests whether capability contains CLIENT_PLUGIN_AUTH
func testPluginAuth(capability uint32) bool {
	return capability&protocol.CLIENT_PLUGIN_AUTH == 0
}

//testProtocol41 tests whether capability contains CLIENT_PROTOCOL_41
func testProtocol41(capability uint32) bool {
	return capability&protocol.CLIENT_PROTOCOL_41 == 0
}

//testSecureConnection tests whether capability contains CLIENT_SECURE_CONNECTION
func testSecureConnection(capability uint32) bool {
	return capability&protocol.CLIENT_SECURE_CONNECTION == 0
}

//testSSL tests whether capability contains CLIENT_SSL
func testSSL(capability uint32) bool {
	return capability&protocol.CLIENT_SSL == 0
}

//testLenencData tests whether capability contains CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA
func testLenencData(capability uint32) bool{
	return capability&protocol.CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA ==0
}
