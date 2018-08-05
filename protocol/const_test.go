package protocol

import (
	. "gopkg.in/check.v1"
)

type MyConstSuite struct{}

var _ = Suite(&MyConstSuite{})

func (s *MyConstSuite) TestClientConst(c *C) {
	data := []uint32{
		CLIENT_LONG_PASSWORD,
		CLIENT_FOUND_ROWS,
		CLIENT_LONG_FLAG,
		CLIENT_CONNECT_WITH_DB,
		CLIENT_NO_SCHEMA,
		CLIENT_COMPRESS,
		CLIENT_ODBC,
		CLIENT_LOCAL_FILES,
		CLIENT_IGNORE_SPACE,
		CLIENT_PROTOCOL_41,
		CLIENT_INTERACTIVE,
		CLIENT_SSL,
		CLIENT_IGNORE_SIGPIPE,
		CLIENT_TRANSACTIONS,
		CLIENT_RESERVED,
		CLIENT_SECURE_CONNECTION,
		CLIENT_MULTI_STATEMENTS,
		CLIENT_MULTI_RESULTS,
		CLIENT_PS_MULTI_RESULTS,
		CLIENT_PLUGIN_AUTH,
		CLIENT_CONNECT_ATTRS,
		CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA,
		CLIENT_CAN_HANDLE_EXPIRED_PASSWORDS,
		CLIENT_SSL_VERIFY_SERVER_CERT,
		CLIENT_REMEMBER_OPTIONS,
	}
	res := []uint32{
		0x01,
		0x02,
		0x04,
		0x08,
		0x10,
		0x20,
		0x40,
		0x80,
		0x100,
		0x200,
		0x400,
		0x800,
		0x1000,
		0x2000,
		0x4000,
		0x8000,
		0x10000,
		0x20000,
		0x40000,
		0x80000,
		0x100000,
		0x200000,
		0x400000,
		0x40000000,
		0x80000000,
	}

	for i := range res {
		c.Assert(data[i], Equals, res[i])
	}
}

//func (s *MyConstSuite) TestClientAllConst(c *C) {
//	data := []uint32{
//		CLIENT_ALL_FLAGS,
//		CLIENT_SSL,
//		^(CLIENT_SSL),
//		CLIENT_COMPRESS,
//		^(CLIENT_COMPRESS),
//		CLIENT_SSL_VERIFY_SERVER_CERT,
//		^(CLIENT_SSL_VERIFY_SERVER_CERT),
//		CLIENT_BASIC_FLAGS,
//	}
//	res := []uint32{
//		0xc07FFFFF,
//		0x800,
//		0xFFFFF7FF,
//		0x20,
//		0xFFFFFFdF,
//		0x40000000,
//		0xBFFFFFFF,
//		0x807FF7DF,
//	}
//
//	for i := range res {
//		c.Assert(data[i], Equals, res[i])
//	}
//}
