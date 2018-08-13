package binlog

import (
	"github.com/juju/errors"
)

//EventHandler represent the handler of Event
type EventHandler interface {
	//Accept a Visitor to the event and return error
	Accept() error
}

type Visitor interface {
	//Visitor visit a Event Handler
	Visit(EventHandler)
}

//NewEventHandler return a EventHandler according to its type
func NewEventHandler(event uint8, buffer []byte) (EventHandler, error) {
	switch event {
	case START_EVENT_V3:
		return NewStartEvent(buffer), nil
	case QUERY_EVENT:
		return NewQueryEvent(buffer), nil
	case STOP_EVENT:
		return NewStopEvent(buffer), nil
	case ROTATE_EVENT:
		return NewRotateEvent(buffer), nil
	case INTVAR_EVENT:
		return NewIntvarEvent(buffer), nil
	case LOAD_EVENT:
		return NewLoadEvent(buffer), nil
	case SLAVE_EVENT:
		return NewSlaveEvent(buffer), nil
	case CREATE_FILE_EVENT:
		return NewCreateFileEvent(buffer), nil
	case APPEND_BLOCK_EVENT:
		return NewAppendBlockEvent(buffer), nil
	case EXEC_LOAD_EVENT:
		return NewExecLoadEvent(buffer), nil
	case DELETE_FILE_EVENT:
		return NewDeleteFileEvent(buffer), nil
	case NEW_LOAD_EVENT:
		return NewLoadEventV2(buffer), nil
	case RAND_EVENT:
		return NewRandEvent(buffer), nil
	case USER_VAR_EVENT:
		return NewUserVarEvent(buffer), nil
	case FORMAT_DESCRIPTION_EVENT:
		return NewFormatDescriptionEvent(buffer), nil
	case XID_EVENT:
		return NewXidEvent(buffer), nil
	case BEGIN_LOAD_QUERY_EVENT:
		return NewBeginLoadQueryEvent(buffer), nil
	case EXECUTE_LOAD_QUERY_EVENT:
		return NewExecuteLoadQueryEvent(buffer), nil
	case TABLE_MAP_EVENT:
		return NewTableMapEvent(buffer), nil
	case WRITE_ROWS_EVENTv0:
		return NewWriteRowEventV0(buffer), nil
	case UPDATE_ROWS_EVENTv0:
		return NewUpdateRowEventV0(buffer), nil
	case DELETE_ROWS_EVENTv0:
		return NewDeleteRowEventV0(buffer), nil
	case WRITE_ROWS_EVENTv1:
		return NewWriteRowEventV1(buffer), nil
	case UPDATE_ROWS_EVENTv1:
		return NewUpdateRowEventV1(buffer), nil
	case DELETE_ROWS_EVENTv1:
		return NewDeleteRowEventV1(buffer), nil
	case INCIDENT_EVENT:
		return NewIncidentEvent(buffer), nil
	case HEARTBEAT_EVENT:
		return NewHeartBeatEvent(buffer), nil
	case IGNORABLE_EVENT:
		return NewIgnorableEvent(buffer), nil
	case ROWS_QUERY_EVENT:
		return NewRowsQueryEvent(buffer), nil
	case WRITE_ROWS_EVENTv2:
		return NewWriteRowEventV2(buffer), nil
	case UPDATE_ROWS_EVENTv2:
		return NewUpdateRowEventV2(buffer), nil
	case DELETE_ROWS_EVENTv2:
		return NewDeleteRowEventV2(buffer), nil
	case GTID_EVENT:
		return NewGtidEvent(buffer), nil
	case ANONYMOUS_GTID_EVENT:
		return NewAnonymousGtidEvent(buffer), nil
	case PREVIOUS_GTIDS_EVENT:
		return NewPrevioudGtidEvent(buffer), nil
	default:
		return nil, errors.Errorf("Binlog:unsupported binlog event type:%d", event)
	}
}

//ReadEvent init a event handler from buffer
func ReadEvent(buffer []byte) (EventHandler, error) {
	if len(buffer) < int(EventHeaderLen) {
		return nil, errors.Errorf("Binlog:event header is only %d bytes,expect %d bytes", len(buffer), EventHeaderLen)
	}

	//get header
	header := NewHeader(buffer)
	return NewEventHandler(header.Kind(), buffer[EventHeaderLen:])

}
