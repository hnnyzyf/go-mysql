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
type Reader struct {
	//the buffer stores the all information in file
	b []byte

	//cuurent reprensent the payload position of current event
	current uint32

	//next reprensent the position of next event
	next uint32

	//output channel
	output chan EventHandler

	//stop channel
	stop chan struct{}

	//represent the error message when error occures
	err error
}

//Create a Reader for binlog
func NewReader(path string) (*Reader, error) {
	if b, err := ioutil.ReadFile(path); err != nil {
		return nil, errors.Annotate(err, "Binlog:fail to create Reader,the error is ")
	} else {
		return &Reader{
			b:       b,
			current: 0,
			next:    MagicNumberLen,
			output:  make(chan EventHandler),
			stop:    make(chan struct{}),
		}, nil
	}
}

//Parse a complete binlog file
func (r *Reader) ParseFile() (chan EventHandler, chan struct{}, error) {
	//check magic number
	if err := r.readMagicNumber(); err != nil {
		return nil, nil, errors.Trace(err)
	}

	//check binlog version
	if err := r.ReadFormatDescriptionEvent(); err != nil {
		return nil, nil, errors.Trace(err)
	}

	//parse all event for loop
	go func() {
		//loop
		for !r.IsEOF() {
			if event, err := r.ReadHeader(); err != nil {
				r.err = errors.Trace(err)
				break
			} else {
				//create a event handler
				if handler, err := NewEventHandler(event, r.b[r.current:r.next]); err != nil {
					r.err = errors.Trace(err)
					break
				} else {
					//add handler to output channel
					r.output <- handler
				}
			}
		}
		//stop
		r.stop <- struct{}{}
	}()

	return r.output, r.stop, nil
}

//Error return the error
func (r *Reader) Error() error {
	return r.err
}

//readMagicNumber check the binlog file header
func (r *Reader) readMagicNumber() error {
	//checkc the magic hader
	if magic := hack.String(r.b[:MagicNumberLen]); magic != MagicNumber {
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
func (r *Reader) ReadFormatDescriptionEvent() error {
	//check header
	if event, err := r.ReadHeader(); err != nil {
		return errors.Trace(err)
	} else {
		if event != FORMAT_DESCRIPTION_EVENT {
			return errors.Errorf("Binlog:expect a FORMAT_DESCRIPTION_EVENT event.")
		}
	}

	//check binlogversion
	buffer := r.b[MagicNumberLen+EventHeaderLen:]
	if binlogVersion, err := binary.ReadInt2(buffer); err != nil {
		return errors.Annotate(err, "Binlog:fail to read binlog version,the error is ")
	} else {
		if !testBinlogVersionV4(binlogVersion) {
			return errInvalidVersion
		}
	}

	return nil
}

//IsEOF test whether we read the eof
func (r *Reader) IsEOF() bool {
	return r.next == uint32(len(r.b))
}

//ReadHeader parse the header and return the event type
func (r *Reader) ReadHeader() (uint8, error) {
	buffer := r.b[r.next:]

	//read
	header := binary.NewBuffer(buffer[:EventHeaderLen])

	//skil timestamp
	if err := header.Skip(4); err != nil {
		return 0, errors.Trace(err)
	}

	//event type
	event, err := header.ReadInt1()
	if err != nil {
		return 0, errors.Trace(err)
	}

	//skip server-id and event-size
	if err := header.Skip(8); err != nil {
		return 0, errors.Trace(err)
	}

	//read position
	if pos, err := header.ReadInt4(); err != nil {
		return 0, errors.Trace(err)
	} else {
		r.current = r.next + EventHeaderLen
		r.next = pos
	}

	return event, nil
}
