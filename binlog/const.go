package binlog

const (
	UNKNOWN_EVENT            uint8 = 0x00
	START_EVENT_V3           uint8 = 0x01
	QUERY_EVENT              uint8 = 0x02
	STOP_EVENT               uint8 = 0x03
	ROTATE_EVENT             uint8 = 0x04
	INTVAR_EVENT             uint8 = 0x05
	LOAD_EVENT               uint8 = 0x06
	SLAVE_EVENT              uint8 = 0x07
	CREATE_FILE_EVENT        uint8 = 0x08
	APPEND_BLOCK_EVENT       uint8 = 0x09
	EXEC_LOAD_EVENT          uint8 = 0x0a
	DELETE_FILE_EVENT        uint8 = 0x0b
	NEW_LOAD_EVENT           uint8 = 0x0c
	RAND_EVENT               uint8 = 0x0d
	USER_VAR_EVENT           uint8 = 0x0e
	FORMAT_DESCRIPTION_EVENT uint8 = 0x0f
	XID_EVENT                uint8 = 0x10
	BEGIN_LOAD_QUERY_EVENT   uint8 = 0x11
	EXECUTE_LOAD_QUERY_EVENT uint8 = 0x12
	TABLE_MAP_EVENT          uint8 = 0x13
	WRITE_ROWS_EVENTv0       uint8 = 0x14
	UPDATE_ROWS_EVENTv0      uint8 = 0x15
	DELETE_ROWS_EVENTv0      uint8 = 0x16
	WRITE_ROWS_EVENTv1       uint8 = 0x17
	UPDATE_ROWS_EVENTv1      uint8 = 0x18
	DELETE_ROWS_EVENTv1      uint8 = 0x19
	INCIDENT_EVENT           uint8 = 0x1a
	HEARTBEAT_EVENT          uint8 = 0x1b
	IGNORABLE_EVENT          uint8 = 0x1c
	ROWS_QUERY_EVENT         uint8 = 0x1d
	WRITE_ROWS_EVENTv2       uint8 = 0x1e
	UPDATE_ROWS_EVENTv2      uint8 = 0x1f
	DELETE_ROWS_EVENTv2      uint8 = 0x20
	GTID_EVENT               uint8 = 0x21
	ANONYMOUS_GTID_EVENT     uint8 = 0x22
	PREVIOUS_GTIDS_EVENT     uint8 = 0x23
)

const (
	LOG_EVENT_BINLOG_IN_USE_F            uint16 = 0x0001
	LOG_EVENT_FORCED_ROTATE_F            uint16 = 0x0002
	LOG_EVENT_THREAD_SPECIFIC_F          uint16 = 0x0004
	LOG_EVENT_SUPPRESS_USE_F             uint16 = 0x0008
	LOG_EVENT_UPDATE_TABLE_MAP_VERSION_F uint16 = 0x0010
	LOG_EVENT_ARTIFICIAL_F               uint16 = 0x0020
	LOG_EVENT_RELAY_LOG_F                uint16 = 0x0040
	LOG_EVENT_IGNORABLE_F                uint16 = 0x0080
	LOG_EVENT_NO_FILTER_F                uint16 = 0x0100
	LOG_EVENT_MTS_ISOLATE_F              uint16 = 0x0200
)

var MagicNumber string = string([]byte{0xfe, 0x62, 0x69, 0x6e})

const MagicNumberLen uint32 = 4

//we only support v4 binlog version
const EventHeaderLen uint32 = 19
