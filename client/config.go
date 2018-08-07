package client

import (
	"crypto/tls"
	"time"

	"github.com/hnnyzyf/go-mysql/protocol"
)

var defaultConfig = &Config{
	Timeout: 0,

	Charset:      "utf8",
	Capabilities: protocol.DefaultClientCapabilities,

	IsSSL: false,

	AutoCommit:     true,
	MutilStatement: false,
}

type Config struct {
	//Dial Timeout
	Timeout time.Duration

	//capabilities flag use in client
	Capabilities uint32
	//charset
	Charset string

	//wether use ssl to connct
	IsSSL     bool
	TlsConfig *tls.Config

	//mutil statement
	MutilStatement bool
	//autocommit
	AutoCommit bool
}
