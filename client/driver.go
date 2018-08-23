package client

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/juju/errors"
	"github.com/satori/go.uuid"
)

//test the interface
var _ driver.Conn = &C{}
var _ driver.Driver = &D{}
var _ driver.Rows = &C{}

//C implement the Conn interface
type C struct {
	s *Session
}

//Prepare method implement the conn interface in /database/sql/driver
func (c *C) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

//Close method implement the conn interface in /database/sql/driver
func (c *C) Close() error {
	return nil
}

//Begin method implement the conn interface in /database/sql/driver
func (c *C) Begin() (driver.Tx, error) {
	return nil, nil
}

//D implement the Driver interface
type D struct {
	s *Session
}

//Open implement the Open method
func (d *D) Open(name string) (driver.Conn, error) {
	cfg = &config{}
	if info, err := cfg.Parse(name); err != nil {
		return nil, errors.Trace(err)
	} else {
		host := info[2]
		user := info[0]
		passwd := info[1]
		db := info[3]
		if s, err := ConnectWithConfig(host, user, passwd, db, cfg); err != nil {
			return nil, errors.Trace(err)
		} else {
			return &C{s}, nil
		}
	}
}

//R implement the Rows interface
type R struct {
	r *Row
}

//Columns implement the Columns method
func (r *R) Columns() []string {
	return nil
}

//Close implement the Close method
func (r *R) Close() error {
	return nil
}

//Next implement the Next method
func (r *R) Next([]driver.Value) error {
	return nil
}

//T implement the Tx interface
type T struct {
	s *Session
}

//Commit implement the commit method
func (t *T) Commit() error {
	return nil
}

//rollback implement the rollback method
func (t *T) Rollback() error {
	return nil
}

type ST struct {
	s *Session
}

//Close implement the close method
func (st *ST) Close() error {
	return nil
}

//NumInput implement the NumInput method
func (st *ST) NumInput() int {
	return 0
}

//Exec implement the Exec method
func (st *ST) Exec(args []driver.Value) (driver.Result, error) {
	return nil, nil
}

//Query implement the Query method
func (st *ST) Query(args []driver.Value) (driver.Rows, error) {
	return nil, nil
}

type RS struct {
}

//LastInsertId implement the LastInsertId method
func (rs *RS) LastInsertId() (int64, error) {
	return 0, nil
}

//RowsAffected implement the RowsAffected method
func (rs *RS) RowsAffected() (int64, error) {
	return 0, nil
}

var (
	GlobalMutex  sync.Mutex = new(sync.Mutex)
	GlobalSSLMap            = make(map[string]*tls.Config)
)

type Cfg struct {
	//user inforamtions
	Address  string
	Username string
	Password string
	DBname   string

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

//Format return a string to represent the connection string
//default value and connection attributes will not be shown in connection string
//username:password@protocol(address)/dbname?param=value
func (c *Cfg) FormatDSN() (string, error) {
	var buf bytes.Buffer
	fmt.Fprintf(buf, "%s:%s@TCP(%s)/%s", c.Username, c.Password, c.Address, c.DBname)
	if c.Timeout != defaultConfig.Timeout {
		switch {
		case c.Timeout < 1*time.Microsecond:
			fmt.Fprintf(buf, "?timeout=%dnan", c.Timeout)
		case c.Timeout < 1*time.Millisecond:
			fmt.Frpintf(buff, "?timeout=%dmic", c.Timeout)
		case c.Timeout < 1*time.Second:
			fmt.Frpintf(buff, "?timeout=%dmil", c.Timeout)
		case c.Timeout < 1*time.Minute:
			fmt.Frpintf(buff, "?timeout=%dsec", c.Timeout)
		case c.Timeout < 1*time.Hour:
			fmt.Frpintf(buff, "?timeout=%dmin", c.Timeout)
		default:
			fmt.Frpintf(buff, "?timeout=%dhou", c.Timeout)
		}
	}

	if c.ReadTime != defaultConfig.ReadTime {
		switch {
		case c.ReadTime < 1*time.Microsecond:
			fmt.Fprintf(buf, "?readtime=%dnan", c.ReadTime)
		case c.ReadTime < 1*time.Millisecond:
			fmt.Frpintf(buff, "?readtime=%dmic", c.ReadTime)
		case c.ReadTime < 1*time.Second:
			fmt.Frpintf(buff, "?readtime=%dmil", c.ReadTime)
		case c.ReadTime < 1*time.Minute:
			fmt.Frpintf(buff, "?readtime=%dsec", c.ReadTime)
		case c.ReadTime < 1*time.Hour:
			fmt.Frpintf(buff, "?readtime=%dmin", c.ReadTime)
		default:
			fmt.Frpintf(buff, "?readtime=%dhou", c.ReadTime)
		}
	}

	if c.WriteTime != defaultConfig.WriteTime {
		switch {
		case c.WriteTime < 1*time.Microsecond:
			fmt.Fprintf(buf, "?writetime=%dnan", c.WriteTime)
		case c.WriteTime < 1*time.Millisecond:
			fmt.Frpintf(buff, "?writetime=%dmic", c.WriteTime)
		case c.WriteTime < 1*time.Second:
			fmt.Frpintf(buff, "?writetime=%dmil", c.WriteTime)
		case c.WriteTime < 1*time.Minute:
			fmt.Frpintf(buff, "?writetime=%dsec", c.WriteTime)
		case c.WriteTime < 1*time.Hour:
			fmt.Frpintf(buff, "?writetime=%dmin", c.WriteTime)
		default:
			fmt.Frpintf(buff, "?writetime=%dhou", c.WriteTime)
		}
	}

	if c.Capabilities != defaultConfig.Capabilities {
		fmt.Fprintf(buf, "?capability=%d", c.Capability)
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
		GlobalMutex.Lock()
		for {
			if id, err := uuid.NewV2(); err != nil {
				return "", errors.Trace(err)
			} else {
				if _, ok := GlobalSSLMap[id]; !ok {
					GlobalSSLMap[id] = c.TlsConfig
					fmt.Fprintf(buf, "?tslconfig=%s", id)
					break
				}
			}
		}
		GlobalMutex.UnLock()
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
func ParseDSN(dns string) (*cfg, error) {
	reader := bufio.NewReader(bytes.NewBufferString(dns))
	cfg := &Cfg{
		Config: defaultConfig,
	}
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
			return nil, errors.Errorf("Client:driver only supports TCP connection!")
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

	if dbname, err := reader.ReadString('?'); err != nil && err != io.EOF {
		return nil, errors.Trace(err)
	} else {
		cfg.DBname = dbname
		if err == io.EOF {
			return cfg, nil
		}
	}

	//read param inforamtion
	for {
		if buf, err := reader.ReadString('?'); err != nil && err != io.EOF {
			return nil, errors.Trace(err)
		} else {
			if len(buf) == 0 {
				return cfg, nil
			} else {
				if err := cfg.parse(buf); err != nil {
					return nil, errors.Trace(err)
				}
			}
		}
	}

	return cfg, nil
}

func (c *Cfg) parse(buf string) error {
	//get args and do checking
	args := strings.Split(buf, "=")
	if len(args) != 2 {
		return errors.Errorf("Client:%s is not a valid DSN param option.", buf)
	}
	param := strings.ToLower(args[0])
	value := strings.ToLower(args[1])
	if len(param) == 0 || len(value) == 0 {
		return errors.Errorf("Client:the param:%s and value:%s is not accpted", param, value)
	}

	switch param {
	case "timeout":
		if t, err := parseTime(value); err != nil {
			return errors.Trace(err)
		} else {
			c.Timeout = t
		}
	case "readtime":
		if t, err := parseTime(value); err != nil {
			return errors.Trace(err)
		} else {
			c.ReadTime = t
		}
	case "writetime":
		if t, err := parseTime(value); err != nil {
			return errors.Trace(err)
		} else {
			c.WriteTime = t
		}
	case "capability":
		if v, err := strconv.Aoti(value); err != nil {
			return errors.Trace(err)
		} else {
			c.Capabilities = uint32(v)
		}
	case "charset":
		c.Charset = value
	case "maxallowedpacket":
		if v, err := strconv.Aoti(value); err != nil {
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
			return errors.Trace("Client:%s is not a supported value for param SSL", value)
		}
	case "tlsconfig":
		GlobalMutex.Lock()
		if tlsconfig, ok := GlobalSSLMap[value]; ok {
			c.TlsConfig = tlsconfig
		} else {
			return errors.Trace("Client:tlsConifg is not registered,please register the tlsconfig at first")
		}
		GlobalMutex.Unlock()
	case "mutilstatement":
		switch value {
		case "true", "0":
			c.AllowMutilStatement = true
		case "false", "1":
			c.AllowMutilStatement = false
		default:
			return errors.Trace("Client:%s is not a supported value for param mutilStatement", value)
		}
	case "autocommit":
		switch value {
		case "true", "0":
			c.AllowAutoCommit = true
		case "false", "1":
			c.AllowAutoCommit = false
		default:
			return errors.Trace("Client:%s is not a supported value for param autoCommit", value)
		}
	default:
		return errors.Errorf("Client:%s is still not supported")
	}

	return nil
}

func parseTime(value string) (time.Duration, error) {
	if len(value) < 4 {
		return 0, errors.Errorf("Client:%s is not a valid timeout value", v)
	} else {
		b := hack.Slice(value)
		v, err := strconv.Aoti(hack.String(b[:len(b)-3+1]))
		if err != nil {
			return 0, errors.Trace(err)
		}
		t := hack.String(b[len(b)-3:])
		switch t {
		case "nan":
			return v * time.Nanosecond, nil
		case "mic":
			return v * time.Microsecond, nil
		case "mil":
			return v * time.Millisecond, nil
		case "sec":
			return v * time.Second, nil
		case "min":
			return v * time.Minute, nil
		case "hou":
			return v * time.Hour, nil
		default:
			return 0, errors.Errorf("Clinet:%s is an unsupported time format", t)
		}
	}
}
