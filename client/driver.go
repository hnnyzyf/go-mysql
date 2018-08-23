package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"database/sql/driver"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
	"github.com/satori/go.uuid"
)

//test the interface
var (
	_ driver.Conn               = &mysqlConn{}
	_ driver.ConnBeginTx        = &mysqlConn{}
	_ driver.ConnPrepareContext = &mysqlConn{}
	_ driver.Execer             = &mysqlConn{}
	_ driver.ExecerContext      = &mysqlConn{}
	_ driver.Pinger             = &mysqlConn{}
	_ driver.Queryer            = &mysqlConn{}
	_ driver.QueryerContext     = &mysqlConn{}
	//unsupported
	//_ driver.SessionResetter    = &mysqlConn{}

	_ driver.Driver = &mysqlDriver{}
	//unsupported
	//_ driver.DriverContext = &mysqlDriver{}

	_ driver.Rows                           = &mysqlRows{}
	_ driver.RowsColumnTypeDatabaseTypeName = &mysqlRows{}
	_ driver.RowsColumnTypeLength           = &mysqlRows{}
	_ driver.RowsColumnTypeNullable         = &mysqlRows{}
	_ driver.RowsColumnTypePrecisionScale   = &mysqlRows{}
	_ driver.RowsColumnTypeScanType         = &mysqlRows{}
	_ driver.RowsNextResultSet              = &mysqlRows{}

	_ driver.Stmt             = &mysqlStmt{}
	_ driver.StmtExecContext  = &mysqlStmt{}
	_ driver.StmtQueryContext = &mysqlStmt{}

	_ driver.Tx = &mysqlTx{}
)

//mysqlConn implement the Conn interface
type mysqlConn struct {
	s *Session
}

//Prepare method implement the conn interface in /database/sql/driver
func (c *mysqlConn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

//Close method implement the conn interface in /database/sql/driver
func (c *mysqlConn) Close() error {
	return nil
}

//Begin method implement the conn interface in /database/sql/driver
func (c *mysqlConn) Begin() (driver.Tx, error) {
	return nil, nil
}

//BeginTx method implement the ConnBeginTx interface in /database/sql/driver
func (c *mysqlConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	return nil, errors.Errorf("client:BeginTx is not supported")
}

//PrepareContext ¶ method implement the ConnPrepareContext interface in /database/sql/driver
func (c *mysqlConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	return nil, errors.Errorf("client:PrepareContext is not supported")
}

//Exec  method implement the Execer interface in /database/sql/driver
func (c *mysqlConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	return nil, errors.Errorf("client:Exec is not supported")
}

//ExecContext  method implement the Execer interface in /database/sql/driver
func (c *mysqlConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	return nil, errors.Errorf("client:ExecContext is not supported")
}

//Ping  method implement the Pinger  interface in /database/sql/driver
func (c *mysqlConn) Ping(ctx context.Context) error {
	return errors.Errorf("client:ExecContext is not supported")
}

//Query method implement the Queryer interface in /database/sql/driver
func (c *mysqlConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	return nil, errors.Errorf("client:Query is not supported")
}

//QueryContext method implement the QueryerContext interface in /database/sql/driver
func (c *mysqlConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	return nil, errors.Errorf("client:QueryContext is not supported")
}

//ResetSession method implement the SessionResetter interface in /database/sql/driver
func (c *mysqlConn) ResetSession(ctx context.Context) error {
	return errors.Errorf("client:ResetSession is not supported")
}

//D implement the Driver interface
type mysqlDriver struct {
	s *Session
}

//Open implement the Open method
func (d *mysqlDriver) Open(name string) (driver.Conn, error) {
	if cfg, err := ParseDSN(name); err != nil {
		return nil, errors.Trace(err)
	} else {
		if s, err := ConnectWithConfig(cfg.Address, cfg.Username, cfg.Password, cfg.DBname, cfg.Config); err != nil {
			return nil, errors.Trace(err)
		} else {
			return &mysqlConn{s}, nil
		}
	}
}

//OpenConnector method implement the DriverContext  interface in /database/sql/driver
//func (c *mysqlDriver) OpenConnector(name string) (driver.Connector, error) {
//return nil, errors.Errorf("client:OpenConnector is not supported ")
//

type mysqlResult struct {
}

//LastInsertId implement the LastInsertId method
func (rs *mysqlResult) LastInsertId() (int64, error) {
	return 0, nil
}

//RowsAffected implement the RowsAffected method
func (rs *mysqlResult) RowsAffected() (int64, error) {
	return 0, nil
}

var (
	GlobalMutex  = new(sync.Mutex)
	GlobalSSLMap = make(map[string]*tls.Config)
)

//R implement the Rows interface
type mysqlRows struct {
	r *response
}

//Columns implement the Columns method
func (r *mysqlRows) Columns() []string {
	columns := make([]string, len(r.r.def))
	for i := range r.r.def {
		def := r.r.def[i]
		columns[i] = hack.String(def.name)
	}
	return columns
}

//Close implement the Close method
func (r *mysqlRows) Close() error {
	return r.r.Terminate()
}

//Next implement the Next method
func (r *mysqlRows) Next(args []driver.Value) error {
	if rawBytes, err := r.r.Next(); err != nil && err != io.EOF {
		return errors.Trace(err)
	} else if err == io.EOF {
		return io.EOF
	} else {
		if len(rawBytes) != len(args) {
			return errors.Errorf("client:not provide enought arguments:%d(%d)", len(rawBytes), len(args))
		}
		for i := range rawBytes {
			raw := rawBytes[i]
			if testNullValue(raw) {
				args[i] = nil
			} else {
				args[i] = raw
			}
		}
		return nil
	}
}

//HasNextResultSet implement the RowsNextResultSet  interface
func (r *mysqlRows) HasNextResultSet() bool {
	return false
}

//NextResultSet implement the RowsNextResultSet  interface
func (r *mysqlRows) NextResultSet() error {
	return errors.Errorf("client:NextResultSet is not supported")
}

//ColumnTypeDatabaseTypeName implement the RowsColumnTypeDatabaseTypeName interface
func (r *mysqlRows) ColumnTypeDatabaseTypeName(index int) string {
	return ""
}

//ColumnTypeLength implement the RowsColumnTypeLength  interface
func (r *mysqlRows) ColumnTypeLength(index int) (length int64, ok bool) {
	return 0, false
}

//ColumnTypeLength implement the RowsColumnTypeLength  interface
func (r *mysqlRows) ColumnTypeNullable(index int) (nullable, ok bool) {
	return false, false
}

//ColumnTypePrecisionScale implement the RowsColumnTypePrecisionScale  interface
func (r *mysqlRows) ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool) {
	return 0, 0, false
}

//ColumnTypeScanType implement the RowsColumnTypeScanType   interface
func (r *mysqlRows) ColumnTypeScanType(index int) reflect.Type {
	return nil
}

type mysqlStmt struct {
	s *Session
}

//Close implement the close method
func (st *mysqlStmt) Close() error {
	return nil
}

//NumInput implement the NumInput method
func (st *mysqlStmt) NumInput() int {
	return 0
}

//Exec implement the Exec method
func (st *mysqlStmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errors.Errorf("client:Exec is not supported")
}

//Query implement the Query method
func (st *mysqlStmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, errors.Errorf("client:Query is not supported")
}

//ExecContext implement the StmtExecContext  interface
func (st *mysqlStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return nil, errors.Errorf("client:ExecContext is not supported")
}

//QueryContext implement the StmtQueryContext interface
func (st *mysqlStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return nil, errors.Errorf("client:QueryContext is not supported")
}

//T implement the Tx interface
type mysqlTx struct {
}

//Commit implement the commit method
func (t *mysqlTx) Commit() error {
	return errors.Errorf("client:Commit is not supported")
}

//rollback implement the rollback method
func (t *mysqlTx) Rollback() error {
	return errors.Errorf("client:Rollback is not supported")
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
