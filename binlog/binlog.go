package binlog

import (
	"io/ioutil"

	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

var errInvalidBinlog error = errors.Errorf("Binlog:not a valid binlog file")
var errInvalidVersion error = errors.Errorf("Binlog:only supprot binlog version:v4")

//we only support
type reader struct {
	//the buffer stores the all information in file
	b []byte
	//pos reprenst the position of next event
	pos uint32
}

//Create a Reader for binlog
func NewReader(path string) (*reader, error) {
	if b, err := ioutil.ReadFile(path); err != nil {
		return nil, errors.Annotate(err, "Binlog:fail to create reader,the error is ")
	} else {
		return &reader{b, 4}, nil
	}
}

//Parse a complete binlog file
func (r *reader) ParseFile() error {
	//check magic number
	if err := r.readMagicNumber(); err != nil {
		return errors.Trace(err)
	}

	//check binlog version
	if err := r.ReadFormatDescriptionEvent(); err != nil {
		return errors.Trace(err)
	}

	return nil
}

//readMagicNumber check the binlog file header
func (r *reader) readMagicNumber() error {
	//checkc the magic hader
	if magic := hack.String(r.b[:4]); magic != MagicNumber {
		return errInvalidBinlog
	}
	return nil
}

//By the time you read the first event from the log you don't know what binlog version the binlog has.
//To determine the version correctly it has to be checked if the first event is:
//
//	- FORMAT_DESCRIPTION_EVENT version = 4
//  - START_EVENT_V3(not support)
//		- if event-size == 13 + 56: version = 1
//		- if event-size == 19 + 56: version = 3
//		- otherwise: invalid
func (r *reader) ReadFormatDescriptionEvent() error {
	buffer := r.b[r.pos:]

	header := binary.NewBuffer(buffer[:EventHeaderLen+2])
	//read timestamp
	if err := header.Skip(4); err != nil {
		return errors.Trace(err)
	}

	//event type
	if event, err := header.ReadInt1(); err != nil {
		return errors.Trace(err)
	} else {
		if event != FORMAT_DESCRIPTION_EVENT {
			return errors.Errorf("Binlog:expect a FORMAT_DESCRIPTION_EVENT event.")
		}
	}

	//skip server-id and event-size
	if err := header.Skip(8); err != nil {
		return errors.Trace(err)
	}

	//read position
	if pos, err := header.ReadInt4(); err != nil {
		return errors.Trace(err)
	} else {
		r.pos = pos
	}

	//skip server-id and event-size
	if err := header.Skip(2); err != nil {
		return errors.Trace(err)
	}

	//check binlogversion
	if binlogVersion, err := header.ReadInt2(); err != nil {
		return errors.Annotate(err, "Binlog:fail to read binlog version,the error is ")
	} else {
		if !testBinlogVersionV4(binlogVersion) {
			return errInvalidVersion
		}
	}

	return nil
}
