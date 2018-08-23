package client

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"runtime"
	"time"

	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

type Config struct {
	//Dial Timeout
	Timeout time.Duration

	//I/O read timeout
	ReadTime time.Duration

	//I/O write timeout
	WriteTime time.Duration

	//capabilities flag use in client
	//You have better not change the Capabilities
	Capabilities uint32

	//charset
	Charset string

	//max_allowed_packet
	MaxAllowedPacket uint32

	//wether use ssl to connct
	AllowSSL  bool
	TlsConfig *tls.Config

	//mutil statement
	AllowMutilStatement bool

	//autocommit
	AllowAutoCommit bool

	//the connection attris
	Attrs Attr
}

var defaultConfig = &Config{
	Timeout:             0,
	ReadTime:            0,
	WriteTime:           0,
	Capabilities:        protocol.DefaultClientCapabilities,
	Charset:             "utf8",
	MaxAllowedPacket:    protocol.MaxAllowedSize,
	AllowSSL:            false,
	TlsConfig:           nil,
	AllowMutilStatement: false,
	AllowAutoCommit:     true,
	Attrs:               defaultConnectionAttrs,
}

//NewConfig create a config with the default value
func NewConfig() *Config {
	return &Config{
		Timeout:             0,
		ReadTime:            0,
		WriteTime:           0,
		Capabilities:        protocol.DefaultClientCapabilities,
		Charset:             "utf8",
		MaxAllowedPacket:    protocol.MaxAllowedSize,
		AllowSSL:            false,
		TlsConfig:           nil,
		AllowMutilStatement: false,
		AllowAutoCommit:     true,
		Attrs:               NewAttrs(),
	}
}

//Attr represent the connection attrs
type Attr map[string]string

var (
	GlobalClientVersion    = "5.7.18"
	defaultConnectionAttrs = Attr{
		"_os":             "linux",
		"_client_name":    "go-mysql-client",
		"_pid":            "0000",
		"_client_version": "5.7.18",
		"_platform":       "x86_64",
		"program_name":    "go-mysql-client",
	}
)

//NewAttr create a new Attr and return the basic inforamtion
func NewAttr() Attr {
	a := new(Attr)
	a["_os"] = runtime.GOOS
	a["_client_name"] = "go-mysql-client"
	a["_pid"] = "0000"
	a["_client_version"] = "5.7.18"
	a["_platform"] = runtime.GOARCH
	return a
}

//size return the length of attr
func (a Attr) size() int {
	length := 0
	for k, v := range a {
		param := k
		value := v
		length += binary.LengthOfString(param)
		length += binary.LengthOfString(value)
	}
	return length
}

//marshal return the serialization results
func (a Attr) marshal() ([]byte, error) {
	//init
	length := a.size()
	size := binary.LengthOfInteger(uint64(length)) + length
	payload := make([]byte, size)

	//write length
	if _, err := payload.WriteLengthEncodedInteger(length); err != nil {
		return nil, errors.Trace(err)
	}

	//write
	for k, v := range a {
		param := k
		value := v
		if err := payload.WriteLengthEncodedString(hack.Slice(param)); err != nil {
			return nil, errors.Trace(err)
		}
		if err := payload.WriteLengthEncodedString(hack.Slice(value)); err != nil {
			return nil, errors.Trace(err)
		}
	}

	return payload.Bytes(), nil
}
