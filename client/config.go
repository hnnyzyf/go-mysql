package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
	"github.com/satori/go.uuid"
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
		Timeout:             defaultConfig.Timeout,
		ReadTime:            defaultConfig.ReadTime,
		WriteTime:           defaultConfig.WriteTime,
		Capabilities:        defaultConfig.Capabilities,
		Charset:             defaultConfig.Charset,
		MaxAllowedPacket:    defaultConfig.MaxAllowedPacket,
		AllowSSL:            defaultConfig.AllowSSL,
		TlsConfig:           defaultConfig.TlsConfig,
		AllowMutilStatement: defaultConfig.AllowMutilStatement,
		AllowAutoCommit:     defaultConfig.AllowAutoCommit,
		Attrs:               NewAttr(),
	}
}

//Attr represent the connection attrs
type Attr map[string]string

var (
	GlobalClientVersion    = "5.7.18"
	defaultConnectionAttrs = Attr{
		"_os":             runtime.GOOS,
		"_client_name":    "go-mysql-client",
		"_pid":            "0000",
		"_client_version": "5.7.18",
		"_platform":       runtime.GOARCH,
		"program_name":    "go-mysql-client",
	}
)

//NewAttr create a new Attr and return the basic inforamtion
func NewAttr() Attr {
	a := make(map[string]string)
	a["_os"] = defaultConnectionAttrs["_os"]
	a["_client_name"] = defaultConnectionAttrs["_client_name"]
	a["_pid"] = defaultConnectionAttrs["_pid"]
	a["_client_version"] = defaultConnectionAttrs["_client_version"]
	a["_platform"] = defaultConnectionAttrs["_platform"]
	a["program_name"] = defaultConnectionAttrs["program_name"]
	return Attr(a)
}

//size return the length of attr
func (a Attr) size() int {
	length := 0
	for k, v := range a {
		param := k
		value := v
		length += binary.LengthOfString(hack.Slice(param))
		length += binary.LengthOfString(hack.Slice(value))
	}
	return length
}

//marshal return the serialization results
func (a Attr) marshal() ([]byte, error) {
	//init
	length := a.size()
	size := binary.LengthOfInteger(uint64(length)) + length
	payload := binary.NewBuffer(make([]byte, size))

	//write length
	if _, err := payload.WriteLengthEncodedInteger(uint64(length)); err != nil {
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

type Cfg struct {
	//user inforamtions
	Address  string
	Username string
	Password string
	DBname   string

	//Uid represent the tls configuration file id
	Uid string

	//other configurations
	*Config
}

//NewCfg create a New cfg for
func NewCfg(address string, username string, password string, dbname string) *Cfg {
	return &Cfg{
		Address:  address,
		Username: username,
		Password: password,
		DBname:   dbname,
		Config:   NewConfig(),
	}
}

//NewCfg create a New cfg for
func NewCfgWithSSL(address string, username string, password string, dbname string, uid string) *Cfg {
	return &Cfg{
		Address:  address,
		Username: username,
		Password: password,
		DBname:   dbname,
		Uid:      uid,
		Config:   NewConfig(),
	}
}

//Format return a string to represent the connection string
//default value and connection attributes will not be shown in connection string
//username:password@protocol(address)/dbname?param=value
func (c *Cfg) FormatDSN() (string, error) {
	buf := bytes.NewBuffer([]byte{})
	fmt.Fprintf(buf, "%s:%s@TCP(%s)/%s", c.Username, c.Password, c.Address, c.DBname)

	if c.Uid != "" {
		if cfg, err := FindTlsConfig(c.Uid); err != nil {
			return "", errors.Trace(err)
		} else {
			if c.TlsConfig != nil {
				return "", errors.Errorf("Clinet:You can not use both cfg.Uid and cfg.TlsConfg,please choose only one param")
			} else {
				c.TlsConfig = cfg
			}
		}
	}

	if c.Timeout != defaultConfig.Timeout {
		fmt.Fprintf(buf, "?timeout=%s", c.Timeout.String())
	}

	if c.ReadTime != defaultConfig.ReadTime {
		fmt.Fprintf(buf, "?readtime=%s", c.ReadTime.String())
	}

	if c.WriteTime != defaultConfig.WriteTime {
		fmt.Fprintf(buf, "?writetime=%s", c.WriteTime.String())
	}

	if c.Capabilities != defaultConfig.Capabilities {
		fmt.Fprintf(buf, "?capability=%d", c.Capabilities)
	}

	if c.Charset != defaultConfig.Charset {
		fmt.Fprintf(buf, "?charset=%s", c.Charset)
	}

	if c.MaxAllowedPacket != defaultConfig.MaxAllowedPacket {
		fmt.Fprintf(buf, "?maxAllowedPacket=%d", c.MaxAllowedPacket)
	}

	if c.AllowSSL != defaultConfig.AllowSSL {
		if c.AllowSSL {
			fmt.Fprintf(buf, "?ssl=true")
		} else {
			fmt.Fprintf(buf, "?ssl=false")
		}
	}

	if c.TlsConfig != nil {
		uid := RegisterTlsConfig(c.TlsConfig)
		fmt.Fprintf(buf, "?tslconfig=%s", uid)
	}

	if c.AllowMutilStatement != defaultConfig.AllowMutilStatement {
		if c.AllowMutilStatement {
			fmt.Fprintf(buf, "?mutilStatement=true")
		} else {
			fmt.Fprintf(buf, "?mutilStatement=false")
		}
	}

	if c.AllowAutoCommit != defaultConfig.AllowAutoCommit {
		if c.AllowMutilStatement {
			fmt.Fprintf(buf, "?autoCommit=true")
		} else {
			fmt.Fprintf(buf, "?autoCommit=false")
		}
	}

	return buf.String(), nil
}

//Parse read from a string and return a configuration
//username:password@protocol(address)/dbname?param=value
func ParseDSN(dns string) (*Cfg, error) {
	buf := hack.Slice(dns)
	reader := binary.NewBuffer(buf)
	cfg := &Cfg{
		Config: NewConfig(),
	}
	//read user information
	//read user information
	if username, err := reader.ReadString(':'); err != nil {
		return nil, errors.Trace(err)
	} else {
		cfg.Username = username
	}

	if password, err := reader.ReadString('@'); err != nil {
		return nil, errors.Trace(err)
	} else {
		cfg.Password = password
	}

	if protocol, err := reader.ReadString('('); err != nil {
		return nil, errors.Trace(err)
	} else {
		if strings.ToUpper(protocol) != "TCP" {
			return nil, errors.Errorf("client:driver only supports TCP connection!")
		}
	}

	if address, err := reader.ReadString(')'); err != nil {
		return nil, errors.Trace(err)
	} else {
		cfg.Address = address
	}

	if _, err := reader.ReadString('/'); err != nil {
		return nil, errors.Trace(err)
	}

	if dbname, err := reader.ReadString('?'); err != nil {
		return nil, errors.Trace(err)
	} else {
		//if dsn is endwith ?,raise error
		if reader.IsEOF() && buf[len(buf)-1] == '?' {
			return nil, errors.Errorf("Clinet:DSN should not endwith ?")
		} else {
			cfg.DBname = dbname
		}
	}

	//read param inforamtion
	set := make(map[string]bool)
	for {
		if b, err := reader.ReadString('?'); err != nil && err != io.EOF {
			return nil, errors.Trace(err)
			//if we find there is nothing to read,just return
		} else if err == io.EOF {
			return cfg, nil
		} else {
			if err := cfg.parse(b, set); err != nil {
				return nil, errors.Trace(err)
			}
		}
	}
}

//parse parse a buffer and set the param and value
func (c *Cfg) parse(buf string, set map[string]bool) error {
	//get args and do checking
	args := strings.Split(buf, "=")
	if len(args) != 2 {
		return errors.Errorf("client:%s is not a valid DSN param option.", buf)
	}
	param := strings.ToLower(args[0])
	value := strings.ToLower(args[1])
	if len(param) == 0 || len(value) == 0 {
		return errors.Errorf("client:the param:%s and value:%s is not accpted", param, value)
	}

	//check whether the param has been set twice
	if _, ok := set[param]; ok {
		return errors.Errorf("client:set the param:%s twice", param)
	} else {
		set[param] = true
	}

	switch param {
	case "timeout":
		if t, err := time.ParseDuration(value); err != nil {
			return errors.Trace(err)
		} else {
			c.Timeout = t
		}
	case "readtime":
		if t, err := time.ParseDuration(value); err != nil {
			return errors.Trace(err)
		} else {
			c.ReadTime = t
		}
	case "writetime":
		if t, err := time.ParseDuration(value); err != nil {
			return errors.Trace(err)
		} else {
			c.WriteTime = t
		}
	case "capability":
		if v, err := strconv.Atoi(value); err != nil {
			return errors.Trace(err)
		} else {
			c.Capabilities = uint32(v)
		}
	case "charset":
		c.Charset = value
	case "maxallowedpacket":
		if v, err := strconv.Atoi(value); err != nil {
			return errors.Trace(err)
		} else {
			c.MaxAllowedPacket = uint32(v)
		}
	case "ssl":
		switch value {
		case "true", "0":
			c.AllowSSL = true
		case "false", "1":
			c.AllowSSL = false
		default:
			return errors.Errorf("client:%s is not a supported value for param SSL", value)
		}
	case "tlsconfig":
		if tlsconfig, err := FindTlsConfig(value); err != nil {
			return errors.Trace(err)
		} else {
			c.TlsConfig = tlsconfig
		}
	case "mutilstatement":
		switch value {
		case "true", "0":
			c.AllowMutilStatement = true
		case "false", "1":
			c.AllowMutilStatement = false
		default:
			return errors.Errorf("client:%s is not a supported value for param mutilStatement", value)
		}
	case "autocommit":
		switch value {
		case "true", "0":
			c.AllowAutoCommit = true
		case "false", "1":
			c.AllowAutoCommit = false
		default:
			return errors.Errorf("client:%s is not a supported value for param autoCommit", value)
		}
	case "compress":
		return errors.Errorf("client:compress is still not supported")
	default:
		return errors.Errorf("client:%s is still not supported")
	}

	return nil
}

//RegisterTlsConfig registe a tls configuration into global ssl map
func RegisterTlsConfig(tlsconfig *tls.Config) string {
	GlobalMutex.Lock()
	defer GlobalMutex.Unlock()
	for {
		uid := uuid.NewV1().String()
		if _, ok := GlobalSSLMap[uid]; !ok {
			GlobalSSLMap[uid] = tlsconfig
			return uid
		}
	}
}

//FindTlsConfig find a tls configuration according to the uuid
func FindTlsConfig(uid string) (*tls.Config, error) {
	GlobalMutex.Lock()
	defer GlobalMutex.Unlock()
	if cfg, ok := GlobalSSLMap[uid]; ok {
		return cfg, nil
	} else {
		return nil, errors.Errorf("client:tlsConifg:%s is not registered,please register the tlsconfig at first", uid)
	}
}
