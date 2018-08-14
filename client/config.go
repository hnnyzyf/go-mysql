package client

import (
	"crypto/tls"
	"time"

	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

var defaultConfig = &config{
	timeout: 0,

	charset:      "utf8",
	capabilities: protocol.DefaultClientCapabilities,

	isSSL:     false,
	tlsConfig: nil,

	autoCommit:     true,
	mutilStatement: false,

	attrs: defaultConnectionAttrs,
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

	//the connection attris
	attrs Attr
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

		attrs: defaultConnectionAttrs,
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

//SetConnectionAttrs set the connectionAttrs
func (c *config) SetConnectionAttrs(attrs map[string]string) {
	c.attrs = Attr(attrs)
}

//Attr represent the connection attrs
type Attr map[string]string

var AllClientVersion = "5.7.18"

var defaultConnectionAttrs = Attr{
	"_os":             "linux",
	"_client_name":    "go-mysql-client",
	"_pid":            "6666",
	"_client_version": AllClientVersion,
	"_platform":       "x86_64",
	"program_name":    "go-mysql-client",
}

func (a Attr) size() int {
	length := 0

	if os, ok := a["_os"]; ok {
		length += 1 + len("_os") + binary.LengthOfInteger(uint64(len(os))) + len(os)
	}

	if client, ok := a["_client_name"]; ok {
		length += 1 + len("_client_name") + binary.LengthOfInteger(uint64(len(client))) + len(client)
	}

	if pid, ok := a["_pid"]; ok {
		length += 1 + len("_pid") + binary.LengthOfInteger(uint64(len(pid))) + len(pid)
	}

	if version, ok := a["_client_version"]; ok {
		length += 1 + len("_client_version") + binary.LengthOfInteger(uint64(len(version))) + len(version)
	}

	if platform, ok := a["_platform"]; ok {
		length += 1 + len("_platform") + binary.LengthOfInteger(uint64(len(platform))) + len(platform)
	}

	if program_name, ok := a["program_name"]; ok {
		length += 1 + len("program_name") + binary.LengthOfInteger(uint64(len(program_name))) + len(program_name)
	}

	return length
}

//marshal return the serialization results
func (a Attr) marshal() ([]byte, error) {
	//init size
	size := a.size()
	length := binary.LengthOfInteger(uint64(size))
	payload := binary.NewBuffer(make([]byte, size+length))

	//no attribute to use,just return error
	if size == 0 {
		return []byte{}, nil
	}

	//write length
	if _, err := payload.WriteLengthEncodedInteger(uint64(size)); err != nil {
		return nil, errors.Trace(err)
	}

	//write os
	if os, ok := a["_os"]; ok {
		//write key
		if _, err := payload.WriteLengthEncodedInteger(uint64(len("_os"))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice("_os")); err != nil {
			return nil, errors.Trace(err)
		}

		//write length
		if _, err := payload.WriteLengthEncodedInteger(uint64(len(os))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice(os)); err != nil {
			return nil, errors.Trace(err)
		}
	}

	if client, ok := a["_client_name"]; ok {
		//write key
		if _, err := payload.WriteLengthEncodedInteger(uint64(len("_client_name"))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice("_client_name")); err != nil {
			return nil, errors.Trace(err)
		}

		//write length
		if _, err := payload.WriteLengthEncodedInteger(uint64(len(client))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice(client)); err != nil {
			return nil, errors.Trace(err)
		}
	}

	if pid, ok := a["_pid"]; ok {
		//write key
		if _, err := payload.WriteLengthEncodedInteger(uint64(len("_pid"))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice("_pid")); err != nil {
			return nil, errors.Trace(err)
		}

		//write length
		if _, err := payload.WriteLengthEncodedInteger(uint64(len(pid))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice(pid)); err != nil {
			return nil, errors.Trace(err)
		}
	}

	if version, ok := a["_client_version"]; ok {
		//write key
		if _, err := payload.WriteLengthEncodedInteger(uint64(len("_client_version"))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice("_client_version")); err != nil {
			return nil, errors.Trace(err)
		}

		//write length
		if _, err := payload.WriteLengthEncodedInteger(uint64(len(version))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice(version)); err != nil {
			return nil, errors.Trace(err)
		}
	}

	if platform, ok := a["_platform"]; ok {
		//write key
		if _, err := payload.WriteLengthEncodedInteger(uint64(len("_platform"))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice("_platform")); err != nil {
			return nil, errors.Trace(err)
		}

		//write length
		if _, err := payload.WriteLengthEncodedInteger(uint64(len(platform))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice(platform)); err != nil {
			return nil, errors.Trace(err)
		}
	}

	if program_name, ok := a["program_name"]; ok {
		//write key
		if _, err := payload.WriteLengthEncodedInteger(uint64(len("program_name"))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice("program_name")); err != nil {
			return nil, errors.Trace(err)
		}

		//write length
		if _, err := payload.WriteLengthEncodedInteger(uint64(len(program_name))); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteStringWithFixLen(hack.Slice(program_name)); err != nil {
			return nil, errors.Trace(err)
		}
	}

	return payload.Bytes(), nil
}
