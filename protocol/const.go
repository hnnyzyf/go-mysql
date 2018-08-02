package protocol

//read https://dev.mysql.com/doc/internals/en/capability-flags.html#packet-Protocol::CapabilityFlags
const (
	CLIENT_LONG_PASSWORD                  uint32 = 1 << 0
	CLIENT_FOUND_ROWS                     uint32 = 1 << 1
	CLIENT_LONG_FLAG                      uint32 = 1 << 2
	CLIENT_CONNECT_WITH_DB                uint32 = 1 << 3
	CLIENT_NO_SCHEMA                      uint32 = 1 << 4
	CLIENT_COMPRESS                       uint32 = 1 << 5
	CLIENT_ODBC                           uint32 = 1 << 6
	CLIENT_LOCAL_FILES                    uint32 = 1 << 7
	CLIENT_IGNORE_SPACE                   uint32 = 1 << 8
	CLIENT_PROTOCOL_41                    uint32 = 1 << 9
	CLIENT_INTERACTIVE                    uint32 = 1 << 10
	CLIENT_SSL                            uint32 = 1 << 11
	CLIENT_IGNORE_SIGPIPE                 uint32 = 1 << 12
	CLIENT_TRANSACTIONS                   uint32 = 1 << 13
	CLIENT_RESERVED                       uint32 = 1 << 14
	CLIENT_SECURE_CONNECTION              uint32 = 1 << 15
	CLIENT_MULTI_STATEMENTS               uint32 = 1 << 16
	CLIENT_MULTI_RESULTS                  uint32 = 1 << 17
	CLIENT_PS_MULTI_RESULTS               uint32 = 1 << 18
	CLIENT_PLUGIN_AUTH                    uint32 = 1 << 19
	CLIENT_CONNECT_ATTRS                  uint32 = 1 << 20
	CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA uint32 = 1 << 21
	CLIENT_CAN_HANDLE_EXPIRED_PASSWORDS   uint32 = 1 << 22

	CLIENT_SSL_VERIFY_SERVER_CERT uint32 = 1 << 30
	CLIENT_REMEMBER_OPTIONS       uint32 = 1 << 31

	///Gather all possible capabilites (flags) supported by the server
	CLIENT_ALL_FLAGS   uint32 = CLIENT_LONG_PASSWORD | CLIENT_FOUND_ROWS | CLIENT_LONG_FLAG | CLIENT_CONNECT_WITH_DB | CLIENT_NO_SCHEMA | CLIENT_COMPRESS | CLIENT_ODBC | CLIENT_LOCAL_FILES | CLIENT_IGNORE_SPACE | CLIENT_PROTOCOL_41 | CLIENT_INTERACTIVE | CLIENT_SSL | CLIENT_IGNORE_SIGPIPE | CLIENT_TRANSACTIONS | CLIENT_RESERVED | CLIENT_SECURE_CONNECTION | CLIENT_MULTI_STATEMENTS | CLIENT_MULTI_RESULTS | CLIENT_PS_MULTI_RESULTS | CLIENT_SSL_VERIFY_SERVER_CERT | CLIENT_REMEMBER_OPTIONS | CLIENT_PLUGIN_AUTH | CLIENT_CONNECT_ATTRS | CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA | CLIENT_CAN_HANDLE_EXPIRED_PASSWORDS
	CLIENT_BASIC_FLAGS uint32 = CLIENT_ALL_FLAGS & (^CLIENT_SSL) & (^CLIENT_COMPRESS) & (^CLIENT_SSL_VERIFY_SERVER_CERT)
)

//read https://dev.mysql.com/doc/internals/en/status-flags.html#packet-Protocol::StatusFlags
const (
	SERVER_STATUS_IN_TRANS uint16 = 1 << iota
	SERVER_STATUS_AUTOCOMMIT
	SERVER_MORE_RESULTS_EXISTS
	SERVER_STATUS_NO_GOOD_INDEX_USED
	SERVER_STATUS_NO_INDEX_USED
	SERVER_STATUS_CURSOR_EXISTS
	SERVER_STATUS_LAST_ROW_SENT
	SERVER_STATUS_DB_DROPPED
	SERVER_STATUS_NO_BACKSLASH_ESCAPES
	SERVER_STATUS_METADATA_CHANGED
	SERVER_QUERY_WAS_SLOW
	SERVER_PS_OUT_PARAMS
	SERVER_STATUS_IN_TRANS_READONLY
	SERVER_SESSION_STATE_CHANGE
)

//read https://dev.mysql.com/doc/internals/en/command-phase.html
const (
	COM_SLEEP uint8 = iota
	COM_QUIT
	COM_INIT_DB
	COM_QUERY
	COM_FIELD_LIST
	COM_CREATE_DB
	COM_DROP_DB
	COM_REFRESH
	COM_SHUTDOWN
	COM_STATISTICS
	COM_PROCESS_INFO
	COM_CONNECT
	COM_PROCESS_KILL
	COM_DEBUG
	COM_PING
	COM_TIME
	COM_DELAYED_INSERT
	COM_CHANGE_USER
	COM_BINLOG_DUMP
	COM_TABLE_DUMP
	COM_CONNECT_OUT
	COM_REGISTER_SLAVE
	COM_STMT_PREPARE
	COM_STMT_EXECUTE
	COM_STMT_SEND_LONG_DATA
	COM_STMT_CLOSE
	COM_STMT_RESET
	COM_SET_OPTION
	COM_STMT_FETCH
	COM_DAEMON
	COM_BINLOG_DUMP_GTID
	COM_RESET_CONNECTION
)

const (
	PayLoadMaxLen uint32 = 1 << 24

	//Default capability flags, including:
	//    CLIENT_LONG_PASSWD
	//    CLIENT_LONG_FLAG
	//    CLIENT_CONNECT_WITH_DB
	//    CLIENT_PROTOCOL_41
	//    CLIENT_TRANSACTIONS
	//    CLIENT_SECURE_CONNECTION
	//    CLIENT_MULTI_STATEMENTS
	//    CLIENT_MULTI_RESULTS
	//    CLIENT_LOCAL_FILES
	//see details in mysql-router
	DefaultClientCapabilities uint32 = 238221

	//Default of max_allowed_packet defined by the MySQL Server (2^30)
	MaxAllowedSize uint32 = 1073741824

	//default DEFAULT_CHARSET is utf8
	DEFAULT_CHARSET uint8 = 0x21
)
