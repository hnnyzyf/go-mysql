package client

import (
	"github.com/hnnyzyf/go-mysql/protocol"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

//RequestCommand write a command to server according to the
func (s *Session) RequestCommand(size int, buffer []byte) error {
	//check size and buffer
	if size != len(buffer) {
		return errors.Errorf("Client:expect %d bytes buffer,but the real size is %d bytes", size, len(buffer))
	}
	//create payload
	payload := binary.NewBuffer(buffer)
	//reset sid
	s.Reset()
	return s.WritePacket(size, payload.ToReader())
}

//RequestComSleep send a COM_SLEEP to server
func (s *Session) RequestComSleep() error {
	size := 1
	buffer := []byte{protocol.COM_SLEEP}
	return s.RequestCommand(size, buffer)
}

//RequestComQuit send a COM_QUIT to server
func (s *Session) RequestComQuit() error {
	size := 1
	buffer := []byte{protocol.COM_QUIT}
	return s.RequestCommand(size, buffer)
}

//RequestComInitDB send a COM_INIT_DB to server
func (s *Session) RequestComInitDB(db string) error {
	size := 1 + len(db)
	buffer := append([]byte{protocol.COM_INIT_DB}, hack.Slice(db)...)
	return s.RequestCommand(size, buffer)
}

//RequestComQuery send a ComQuery to server
func (s *Session) RequestComQuery(query string) error {
	size := len(query)
	buffer := append([]byte{protocol.COM_QUERY}, hack.Slice(query)...)
	return s.RequestCommand(size, buffer)
}

//RequestComFieldList send a COM_FIELD_LIST to server
func (s *Session) RequestComFieldList(...interface{}) error {
	return errors.Errorf("Client:COM_FIELD_LIST is deprecated and will be removed in a future version of MySQL")
}

//RequestComCreateDB send a COM_CREATE_DB to server
//this command seems unusefull,always return error(6143):Unknown command
func (s *Session) RequestComCreateDB(db string) error {
	size := 1 + len(db)
	buffer := append([]byte{protocol.COM_CREATE_DB}, hack.Slice(db)...)
	return s.RequestCommand(size, buffer)
}

//RequestComDropDB send a COM_DROP_DB to server
//this command seems unusefull,always return error(6143):Unknown command
func (s *Session) RequestComDropDB(db string) error {
	size := 1 + len(db)
	buffer := append([]byte{protocol.COM_DROP_DB}, hack.Slice(db)...)
	return s.RequestCommand(size, buffer)
}

//RequestComRefresh send a COM_REFRESH to server
func (s *Session) RequestComRefresh(...interface{}) error {
	return errors.Errorf("Client:COM_REFRESH is deprecated and will be removed in a future version of MySQL")
}

//RequestComShutDown send a COM_SHUTDOWN to server
func (s *Session) RequestComShutDown(...interface{}) error {
	return errors.Errorf("Client:COM_SHUTDOWN is deprecated and will be removed in a future version of MySQL")
}

//WriteShutDown send a COM_STATISTICS to server
func (s *Session) RequestComStatistics(...interface{}) error {
	return errors.Errorf("Client:COM_STATISTICS is deprecated and will be removed in a future version of MySQL")
}

//RequestComProcessInfo send a COM_PROCESS_INFO to server
func (s *Session) RequestComProcessInfo(...interface{}) error {
	return errors.Errorf("Client:COM_PROCESS_INFO is deprecated and will be removed in a future version of MySQL")
}

//RequestComConnect send a COM_CONNECT to server
func (s *Session) RequestComConnect() error {
	size := 1
	buffer := []byte{protocol.COM_CONNECT}
	return s.RequestCommand(size, buffer)
}

//RequestComProcessKill send a COM_PROCESS_KILL to server
func (s *Session) RequestComProcessKill(id uint32) error {
	return errors.Errorf("Client:COM_PROCESS_KILL is deprecated and will be removed in a future version of MySQL")
}

//RequestComDebug send a COM_DEBUG to server
func (s *Session) RequestComDebug() error {
	size := 1
	buffer := []byte{protocol.COM_DEBUG}
	return s.RequestCommand(size, buffer)
}

//RequestComPing send a COM_PING to server
func (s *Session) RequestComPing() error {
	size := 1
	buffer := []byte{protocol.COM_PING}
	return s.RequestCommand(size, buffer)
}

//RequestComTime send a COM_TIME to server
func (s *Session) RequestComTime() error {
	size := 1
	buffer := []byte{protocol.COM_TIME}
	return s.RequestCommand(size, buffer)
}

//RequestComDelayedInsert send a COM_DELAYED_INSERT to server
func (s *Session) RequestComDelayedInsert() error {
	size := 1
	buffer := []byte{protocol.COM_DELAYED_INSERT}
	return s.RequestCommand(size, buffer)
}

//RequestComChangeUser send a COM_CHANGE_USER to server
func (s *Session) RequestComChangeUser(...interface{}) error {
	return errors.Errorf("Clinet:COM_CHANGE_USER is not implemented")
}

//RequestComResetConnection send a COM_RESET_CONNECTION to server
func (s *Session) RequestComResetConnection() error {
	size := 1
	buffer := []byte{protocol.COM_RESET_CONNECTION}
	return s.RequestCommand(size, buffer)
}

//RequestComDeamon send a COM_DAEMON to server
func (s *Session) RequestComDeamon() error {
	size := 1
	buffer := []byte{protocol.COM_DAEMON}
	return s.RequestCommand(size, buffer)
}

//RequestComStmtPrepare send a COM_STMT_PREPARE to server
func (s *Session) RequestComStmtPrepare(query string) error {
	size := 1 + len(query)
	buffer := append([]byte{protocol.COM_STMT_PREPARE}, hack.Slice(query)...)
	return s.RequestCommand(size, buffer)
}

//RequestComStmtSendLongData send a COM_STMT_SEND_LONG_DATA to server
func (s *Session) RequestComStmtSendLongData(...interface{}) error {
	return errors.Errorf("Clinet:COM_STMT_SEND_LONG_DATA is not implemented")
}

//RequestComStmtExecute send a COM_STMT_EXECUTE to server
func (s *Session) RequestComStmtExecute(...interface{}) error {
	return errors.Errorf("Clinet:COM_STMT_EXECUTE is not implemented")
}

//RequestComStmtClose send a COM_STMT_CLOSE to server
func (s *Session) RequestComStmtClose(...interface{}) error {
	return errors.Errorf("Clinet:COM_STMT_CLOSE is not implemented")
}

//RequestComStmtReset send a COM_STMT_RESET to server
func (s *Session) RequestComStmtReset(...interface{}) error {
	return errors.Errorf("Clinet:COM_STMT_RESET is not implemented")
}

const (
	MYSQL_OPTION_MULTI_STATEMENTS_ON  uint16 = 0x00
	MYSQL_OPTION_MULTI_STATEMENTS_OFF uint16 = 0x01
)

//RequestComSetOption send a COM_SET_OPTION to server
func (s *Session) RequestComSetOption(mutil bool) error {
	buffer := []byte{protocol.COM_SET_OPTION, 0, 0}
	option := MYSQL_OPTION_MULTI_STATEMENTS_OFF
	if mutil {
		option = MYSQL_OPTION_MULTI_STATEMENTS_ON
	}

	if err := binary.WriteInt2(buffer[1:], option); err != nil {
		return errors.Trace(err)
	} else {
		return s.RequestCommand(3, buffer)
	}
}

//RequestComStmtFetch send a COM_STMT_FETCH to server
func (s *Session) RequestComStmtFetch(id uint32, num uint32) error {
	size := 1 + 4 + 4
	buffer := []byte{protocol.COM_STMT_FETCH, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if err := binary.WriteInt4(buffer[1:5], id); err != nil {
		return errors.Trace(err)
	}
	if err := binary.WriteInt4(buffer[5:9], num); err != nil {
		return errors.Trace(err)
	}
	return s.RequestCommand(size, buffer)
}
