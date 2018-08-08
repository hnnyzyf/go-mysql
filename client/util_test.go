package client

import (
	. "gopkg.in/check.v1"

	"github.com/hnnyzyf/go-mysql/protocol"
)

type MytestSuite struct{}

var _ = Suite(&MytestSuite{})

func (s *MytestSuite) TestVersion(c *C) {
	data := []struct {
		version uint8
		res     bool
	}{
		{0x01, false},
		{0x0a, true},
		{0xff, false},
	}

	for i := range data {
		c.Assert(testProtocolVersion(data[i].version), Equals, data[i].res)
	}
}

type testfunction func(uint32) bool

func (s *MytestSuite) TestClientConst(c *C) {
	data := []struct {
		f   testfunction
		c   uint32
		res bool
	}{
		{testPluginAuth, protocol.CLIENT_PLUGIN_AUTH, true},
		{testProtocol41, protocol.CLIENT_PROTOCOL_41, true},
		{testSecureConnection, protocol.CLIENT_SECURE_CONNECTION, true},
		{testSSL, protocol.CLIENT_SSL, true},
		{testLenencData, protocol.CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA, true},
		{testConnectWithDB, protocol.CLIENT_CONNECT_WITH_DB, true},
		{testConnectAttrs, protocol.CLIENT_CONNECT_ATTRS, true},
	}

	for i := range data {
		c.Assert(data[i].f(data[i].c), Equals, data[i].res)
	}
}
