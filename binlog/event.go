package binlog

import (
	"github.com/juju/errors"
)

//Header represent the event header
type Header struct {
	Timestamp []byte
	EventType []byte
	ServerId  []byte
	EventSize []byte
	LogPos    []byte
	Flags     []byte
}

//header only have 19 bytes
func NewHeader(b []byte) *Header {
	return &Header{
		Timestamp: b[:4],
		EventType: b[4:6],
		ServerId:  b[6:10],
		EventSize: b[10:14],
		LogPos:    b[14:18],
		Flags:     b[18:19],
	}
}

//represent the start_event_v3
type StartEvent struct{}

func NewStartEvent(b []byte) *StartEvent {
	return &StartEvent{}
}

func (e *StartEvent) Parse() error {
	return nil
}

//represent the query_event
type QueryEvent struct{}

func NewQueryEvent(b []byte) *QueryEvent {
	return &QueryEvent{}
}

func (e *QueryEvent) Parse() error {
	return nil
}

//represent the stop_event
type StopEvent struct{}

func NewStopEvent(b []byte) *StopEvent {
	return &StopEvent{}
}

func (e *StopEvent) Parse() error {
	return nil
}

//represent the rotate_event
type RotateEvent struct{}

func NewRotateEvent(b []byte) *RotateEvent {
	return &RotateEvent{}
}

func (e *RotateEvent) Parse() error {
	return nil
}

//represent the intvar_event
type IntvarEvent struct{}

func NewIntvarEvent(b []byte) *IntvarEvent {
	return &IntvarEvent{}
}

func (e *IntvarEvent) Parse() error {
	return nil
}

//represent the Load_event
type LoadEvent struct{}

func NewLoadEvent(b []byte) *LoadEvent {
	return &LoadEvent{}
}

func (e *LoadEvent) Parse() error {
	return nil
}

//represent the Slave_event
type SlaveEvent struct{}

func NewSlaveEvent(b []byte) *SlaveEvent {
	return &SlaveEvent{}
}

func (e *SlaveEvent) Parse() error {
	return nil
}

//represent the Create_File_event
type CreateFileEvent struct{}

func NewCreateFileEvent(b []byte) *CreateFileEvent {
	return &CreateFileEvent{}
}

func (e *CreateFileEvent) Parse() error {
	return nil
}

//represent the Append_Block_Event
type AppendBlockEvent struct{}

func NewAppendBlockEvent(b []byte) *AppendBlockEvent {
	return &AppendBlockEvent{}
}

func (e *AppendBlockEvent) Parse() error {
	return nil
}

//represent the Exec_Load_Event
type ExecLoadEvent struct{}

func NewExecLoadEvent(b []byte) *ExecLoadEvent {
	return &ExecLoadEvent{}
}

func (e *ExecLoadEvent) Parse() error {
	return nil
}

//represent the Delete_File_Event
type DeleteFileEvent struct{}

func NewDeleteFileEvent(b []byte) *DeleteFileEvent {
	return &DeleteFileEvent{}
}

func (e *DeleteFileEvent) Parse() error {
	return nil
}

//represent the New_Load_Event
type LoadEventV2 struct{}

func NewLoadEventV2(b []byte) *LoadEventV2 {
	return &LoadEventV2{}
}

func (e *LoadEventV2) Parse() error {
	return nil
}

//represent the Rand_Event
type RandEvent struct{}

func NewRandEvent(b []byte) *RandEvent {
	return &RandEvent{}
}

func (e *RandEvent) Parse() error {
	return nil
}

//represent the User_Var_Event
type UserVarEvent struct{}

func NewUserVarEvent(b []byte) *UserVarEvent {
	return &UserVarEvent{}
}

func (e *UserVarEvent) Parse() error {
	return nil
}

//represent the Format_Description_Event
type FormatDescriptionEvent struct{}

func NewFormatDescriptionEvent(b []byte) *FormatDescriptionEvent {
	return &FormatDescriptionEvent{}
}

func (e *FormatDescriptionEvent) Parse() error {
	return nil
}

//represent the Xid_Event
type XidEvent struct{}

func NewXidEvent(b []byte) *XidEvent {
	return &XidEvent{}
}

func (e *XidEvent) Parse() error {
	return nil
}

//represent the Begin_Load_Query_Event
type BeginLoadQueryEvent struct{}

func NewBeginLoadQueryEvent(b []byte) *BeginLoadQueryEvent {
	return &BeginLoadQueryEvent{}
}

func (e *BeginLoadQueryEvent) Parse() error {
	return nil
}

//represent the Execute_Load_Query_Event
type ExecuteLoadQueryEvent struct{}

func NewExecuteLoadQueryEvent(b []byte) *ExecuteLoadQueryEvent {
	return &ExecuteLoadQueryEvent{}
}

func (e *ExecuteLoadQueryEvent) Parse() error {
	return nil
}

//represent the Table_Map_Event
type TableMapEvent struct{}

func NewTableMapEvent(b []byte) *TableMapEvent {
	return &TableMapEvent{}
}

func (e *TableMapEvent) Parse() error {
	return nil
}

//represent the Write_Row_EventV0
type WriteRowEventV0 struct{}

func NewWriteRowEventV0(b []byte) *WriteRowEventV0 {
	return &WriteRowEventV0{}
}

func (e *WriteRowEventV0) Parse() error {
	return nil
}

//represent the Update_Row_EventV0
type UpdateRowEventV0 struct{}

func NewUpdateRowEventV0(b []byte) *UpdateRowEventV0 {
	return &UpdateRowEventV0{}
}

func (e *UpdateRowEventV0) Parse() error {
	return nil
}

//represent the Delete_Row_EventV0
type DeleteRowEventV0 struct{}

func NewDeleteRowEventV0(b []byte) *DeleteRowEventV0 {
	return &DeleteRowEventV0{}
}

func (e *DeleteRowEventV0) Parse() error {
	return nil
}

//represent the Write_Row_EventV1
type WriteRowEventV1 struct{}

func NewWriteRowEventV1(b []byte) *WriteRowEventV1 {
	return &WriteRowEventV1{}
}

func (e *WriteRowEventV1) Parse() error {
	return nil
}

//represent the Update_Row_EventV1
type UpdateRowEventV1 struct{}

func NewUpdateRowEventV1(b []byte) *UpdateRowEventV1 {
	return &UpdateRowEventV1{}
}

func (e *UpdateRowEventV1) Parse() error {
	return nil
}

//represent the Delete_Row_EventV1
type DeleteRowEventV1 struct{}

func NewDeleteRowEventV1(b []byte) *DeleteRowEventV1 {
	return &DeleteRowEventV1{}
}

func (e *DeleteRowEventV1) Parse() error {
	return nil
}

//represent the Incident_Event
type IncidentEvent struct{}

func NewIncidentEvent(b []byte) *IncidentEvent {
	return &IncidentEvent{}
}

func (e *IncidentEvent) Parse() error {
	return nil
}

//represent the HeartBeat_Event
type HeartBeatEvent struct{}

func NewHeartBeatEvent(b []byte) *HeartBeatEvent {
	return &HeartBeatEvent{}
}

func (e *HeartBeatEvent) Parse() error {
	return nil
}

//represent the Ignorable_Event
type IgnorableEvent struct{}

func NewIgnorableEvent(b []byte) *IgnorableEvent {
	return &IgnorableEvent{}
}

func (e *IgnorableEvent) Parse() error {
	return nil
}

//represent the Rows_Query_Event
type RowsQueryEvent struct{}

func NewRowsQueryEvent(b []byte) *RowsQueryEvent {
	return &RowsQueryEvent{}
}

func (e *RowsQueryEvent) Parse() error {
	return nil
}

//represent the Write_Row_EventV2
type WriteRowEventV2 struct{}

func NewWriteRowEventV2(b []byte) *WriteRowEventV2 {
	return &WriteRowEventV2{}
}

func (e *WriteRowEventV2) Parse() error {
	return nil
}

//represent the Update_Row_EventV2
type UpdateRowEventV2 struct{}

func NewUpdateRowEventV2(b []byte) *UpdateRowEventV2 {
	return &UpdateRowEventV2{}
}

func (e *UpdateRowEventV2) Parse() error {
	return nil
}

//represent the Delete_Row_EventV2
type DeleteRowEventV2 struct{}

func NewDeleteRowEventV2(b []byte) *DeleteRowEventV2 {
	return &DeleteRowEventV2{}
}

func (e *DeleteRowEventV2) Parse() error {
	return nil
}

//represent the Gtid_Event
type GtidEvent struct{}

func NewGtidEvent(b []byte) *GtidEvent {
	return &GtidEvent{}
}

func (e *GtidEvent) Parse() error {
	return nil
}

//represent the Anonymous_Gtid_Event
type AnonymousGtidEvent struct{}

func NewAnonymousGtidEvent(b []byte) *AnonymousGtidEvent {
	return &AnonymousGtidEvent{}
}

func (e *AnonymousGtidEvent) Parse() error {
	return nil
}

//represent the Previoud_Gtid_Event
type PrevioudGtidEvent struct{}

func NewPrevioudGtidEvent(b []byte) *PrevioudGtidEvent {
	return &PrevioudGtidEvent{}
}

func (e *PrevioudGtidEvent) Parse() error {
	return nil
}
