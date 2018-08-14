package client

import (
	"github.com/hnnyzyf/go-mysql/protocol"
	. "gopkg.in/check.v1"
)

type MyConfigSuite struct{}

var _ = Suite(&MyConfigSuite{})

func (s *MyConfigSuite) TestConfig(c *C) {
	data := []*config{
		{0, protocol.DefaultClientCapabilities, "utf8", false, nil, true, false, defaultConnectionAttrs},
	}

	for i := range data {
		cfg := NewConfig()
		cfg.SetTimeout(data[i].timeout)
		cfg.SetCapability(data[i].capabilities)
		cfg.SetCharSet(data[i].charset)
		cfg.SetSSL(data[i].isSSL, data[i].tlsConfig)
		cfg.SetMutilStatement(data[i].mutilStatement)
		cfg.SetAutoCommit(data[i].autoCommit)

		c.Assert(data[i].timeout, Equals, cfg.timeout)
		c.Assert(data[i].charset, Equals, cfg.charset)
		c.Assert(data[i].capabilities, Equals, cfg.capabilities)
		c.Assert(data[i].isSSL, Equals, cfg.isSSL)
		c.Assert(data[i].tlsConfig, Equals, cfg.tlsConfig)
		c.Assert(data[i].mutilStatement, Equals, cfg.mutilStatement)
		c.Assert(data[i].autoCommit, Equals, cfg.autoCommit)
	}
}

func (s *MyConfigSuite) TestAttr(c *C) {
	attr := Attr{
		"_os":             "1",
		"_client_name":    "2",
		"_pid":            "3",
		"_client_version": "4",
		"_platform":       "5",
		"program_name":    "6",
	}

	size := len("_os"+"_client_name"+"_pid"+"_client_version"+"_platform"+"program_name") + 6 + 6 + 6
	c.Assert(attr.size(), Equals, size)

	_, err := attr.marshal()
	c.Assert(err, IsNil)

}
