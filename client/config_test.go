package client

import (
	"github.com/hnnyzyf/go-mysql/protocol"
	. "gopkg.in/check.v1"
)

type MyConfigSuite struct{}

var _ = Suite(&MyConfigSuite{})

func (s *MyConfigSuite) TestConfig(c *C) {
	data := []*config{
		{0, protocol.DefaultClientCapabilities, "utf8", false, nil, true, false},
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
