package client

import (
	"crypto/tls"
	"time"

	"github.com/hnnyzyf/go-mysql/protocol"
)

var defaultConfig = &config{
	timeout: 0,

	charset:      "utf8",
	capabilities: protocol.DefaultClientCapabilities,

	isSSL:     false,
	tlsConfig: nil,

	autoCommit:     true,
	mutilStatement: false,
}

type config struct {
	//Dial Timeout
	timeout time.Duration

	//capabilities flag use in client
	//You have better not change the Capabilities
	capabilities uint32
	//charset
	charset string

	//wether use ssl to connct
	isSSL     bool
	tlsConfig *tls.Config

	//mutil statement
	mutilStatement bool
	//autocommit
	autoCommit bool
}

func NewConfig() *config {
	return &config{
		timeout: 0,

		charset:      "utf8",
		capabilities: protocol.DefaultClientCapabilities,

		isSSL:     false,
		tlsConfig: nil,

		autoCommit:     true,
		mutilStatement: false,
	}
}

//SetTimeout add the timeout for dail
func (c *config) SetTimeout(t time.Duration) {
	c.timeout = t
}

//SetCapability add the capability flag
func (c *config) SetCapability(capability uint32) {
	c.capabilities = capability
}

//SetCharSet add the charset
func (c *config) SetCharSet(s string) {
	c.charset = s
}

//SetSSL add the ssl
func (c *config) SetSSL(ok bool, cfg *tls.Config) {
	c.isSSL = ok
	c.tlsConfig = cfg
}

//SetMutilStatement set the mutistatement
func (c *config) SetMutilStatement(ok bool) {
	c.mutilStatement = ok
}

//SetAutoCommit set the autocommit
func (c *config) SetAutoCommit(ok bool) {
	c.autoCommit = ok
}
