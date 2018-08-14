package client

import (
	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/binary"
)

//testProtocolVersion tests whether  the server supprot v10 protocol
func testProtocolVersion(version uint8) bool {
	return version == 0x0A
}

//testPluginAuth tests whether capability contains CLIENT_PLUGIN_AUTH
func testPluginAuth(capability uint32) bool {
	return capability&protocol.CLIENT_PLUGIN_AUTH != 0
}

//testProtocol41 tests whether capability contains CLIENT_PROTOCOL_41
func testProtocol41(capability uint32) bool {
	return capability&protocol.CLIENT_PROTOCOL_41 != 0
}

//testSecureConnection tests whether capability contains CLIENT_SECURE_CONNECTION
func testSecureConnection(capability uint32) bool {
	return capability&protocol.CLIENT_SECURE_CONNECTION != 0
}

//testSSL tests whether capability contains CLIENT_SSL
func testSSL(capability uint32) bool {
	return capability&protocol.CLIENT_SSL != 0
}

//testLenencData tests whether capability contains CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA
func testLenencData(capability uint32) bool {
	return capability&protocol.CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA != 0
}

//testConnectWithDB tests whether capability contains CLIENT_CONNECT_WITH_DB
func testConnectWithDB(capability uint32) bool {
	return capability&protocol.CLIENT_CONNECT_WITH_DB != 0
}

//testConnectAttrs tests whether capability contains CLIENT_CONNECT_ATTRS
func testConnectAttrs(capability uint32) bool {
	return capability&protocol.CLIENT_CONNECT_ATTRS != 0
}

//testOkPacket tests whether current packet is ok packet
func testOkPacket(header uint8) bool {
	return header == protocol.OK_Packet
}

//testErrpacket test whether current packet is err packet
func testErrPacket(header uint8) bool {
	return header == protocol.ERR_Packet
}

//teseofpacket test whether current packet is eof packet
func testEofPacket(header uint8) bool {
	return header == protocol.EOF_Packet
}

//teseofpacket test whether current packet is eof packet
func tesResultPacket(header int) bool {
	return header == bianry.NotLengthEncodeInteger
}
