package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"database/sql/driver"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
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
	if err := c.s.RequestComStmtPrepare(query); err != nil {
		return nil, errors.Trace(err)
	} else {
		return &mysqlStmt{c.s}, nil
	}
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

//Exec method implement the Execer interface in /database/sql/driver
//If a Conn implements neither ExecerContext nor Execer Execer,
//the sql package's DB.Exec will first prepare a query,
//execute the statement, and then close the statement.
//Deprecated: Drivers should implement ExecerContext instead.
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
	if _, err := parseQuery(query, args); err != nil {
		return nil, errors.Trace(err)
	} else {
		return nil, nil
	}
}

//QueryContext method implement the QueryerContext interface in /database/sql/driver
func (c *mysqlConn) QueryContext(ctx context.Context, query string, nargs []driver.NamedValue) (driver.Rows, error) {
	if args, err := parseNamedValue(nargs); err != nil {
		return nil, errors.Trace(err)
	} else {
		//get the rows
		rowMsg := make(chan driver.Rows)
		errMsg := make(chan error)
		go func() {
			if rows, err := c.Query(query, args); err != nil {
				errMsg <- errors.Trace(err)
			} else {
				rowMsg <- rows
			}
		}()

		//get the message
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case err := <-errMsg:
			return nil, errors.Trace(err)
		case row := <-rowMsg:
			return row, nil
		}
	}

}

//ResetSession method implement the SessionResetter interface in /database/sql/driver
func (c *mysqlConn) ResetSession(ctx context.Context) error {
	return errors.Errorf("client:ResetSession is not supported")
}

//parseNamedValue return the driverValue
//we ingore the name param in the NamedValue
//the nargs and args should have the same order
//Undo:supported the name param in the future
func parseNamedValue(nargs []driver.NamedValue) ([]driver.Value, error) {
	args := make([]driver.Value, len(nargs))
	for i := range nargs {
		args[i] = nargs[i].Value
	}
	return args, nil
}

//parseQuery assin the value to query
//we hopen the query will use ? to represent a value
//for example:
// select * from test where id = ? and name = ?
func parseQuery(query string, args []driver.Value) (string, error) {
	//basic test
	if strings.Count(query, "?") != len(args) {
		return "", errors.Errorf("client:no enough ? in the query")
	}

	//init
	buffer := bytes.NewBuffer([]byte{})
	reader := binary.NewBuffer(hack.Slice(query))
	for i := range args {
		//write query slice
		if q, err := reader.ReadSlice('?'); err != nil {
			return "", errors.Trace(err)
		} else {
			buffer.Write(q)
		}
		//write value
		if v, err := parseValue(args[i]); err != nil {
			return "", errors.Trace(err)
		} else {
			buffer.WriteString(v)
		}
	}

	return buffer.String(), nil

}

//parseValue convert driver.Value to string
func parseValue(arg driver.Value) (string, error) {
	switch val := arg.(type) {
	case int:
		return strconv.Itoa(val), nil
	case int8:
		return strconv.FormatInt(int64(val), 10), nil
	case int16:
		return strconv.FormatInt(int64(val), 10), nil
	case int32:
		return strconv.FormatInt(int64(val), 10), nil
	case int64:
		return strconv.FormatInt(int64(val), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(val), 10), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(val), nil
	case time.Time:
		return val.Format(protocol.DefaultTimeFormat), nil
	case string:
		return strconv.Quote(val), nil
	case []byte:
		return hack.String(val), nil
	default:
		return "", errors.Errorf("client:not supported value type")
	}
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
//Exec executes a query that doesn't return rows, such as an INSERT or UPDATE.
//Deprecated: Drivers should implement StmtExecContext instead (or additionally).
func (st *mysqlStmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errors.Errorf("client:Exec method for Stmt is not supported anymore")
}

//Query implement the Query method
//Query executes a query that may return rows, such as a SELECT
// Deprecated: Drivers should implement StmtQueryContext instead (or additionally).
func (st *mysqlStmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, errors.Errorf("client:Query method for Stmt is not supported")
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
