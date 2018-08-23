package protocol

//The max length of a payload
const PayLoadMaxLen uint32 = 1 << 24

//Default of max_allowed_packet defined by the MySQL Server (2^30)
const MaxAllowedSize uint32 = 1073741824

//default DEFAULT_CHARSET is utf8
const DefaultCharacter uint8 = 0x21

//Default capability flags, including:
//see details in mysql-router
//mysql-router-8.0/src/mysql_protocol/include/mysqlrouter/mysql_protocol/handshake_packet.h line 68
const DefaultClientCapabilities uint32 = CLIENT_LONG_PASSWORD |
	CLIENT_LONG_FLAG |
	CLIENT_PROTOCOL_41 |
	CLIENT_TRANSACTIONS |
	CLIENT_SECURE_CONNECTION |
	CLIENT_MULTI_RESULTS |
	CLIENT_PLUGIN_AUTH |
	CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA |
	CLIENT_CONNECT_ATTRS

const DefaultNull byte = 0xfb