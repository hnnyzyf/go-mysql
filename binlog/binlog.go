package binlog

import (
	"os"

	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

var errInvalidBinlog error = errors.Errorf("Binlog:not a valid binlog file")
var errInvalidVersion error = errors.Errorf("Binlog:only supprot binlog version:v4")

//we only support
type Reader struct {
	//the buffer stores the all information in file
	cache *binary.Cache

	//pos represent the position of next event
	pos uint32

	//represent the error message when error occures
	err error
}

//Create a Reader for binlog
func NewReader() *Reader {
	return &Reader{}
}

//Parse a complete binlog file
func (r *Reader) Open(path string) (chan EventHandler, error) {
	//open binlog file
	if file, err := os.Open(path); err != nil {
		return nil, errors.Trace(err)
	} else {
		r.cache = binary.NewCache(file)
	}

	//check magic number
	if err := r.readMagicNumber(); err != nil {
		return nil, errors.Trace(err)
	}

	//check binlog version
	if err := r.readFormatDescriptionEvent(); err != nil {
		return nil, errors.Trace(err)
	}

	//init output channel
	output := make(chan EventHandler)

	//parse all event for loop
	go func() {
		//loop
		for {
			//read header
			header, err := r.readHeader()
			if err != nil {
				r.err = errors.Trace(err)
				close(output)
				return
			}

			//read payload
			buffer, err := r.readPayload(header)
			if err != nil {
				r.err = errors.Trace(err)
				close(output)
				return
			}

			//create a event handler
			handler, err := NewEventHandler(header.Kind(), buffer)
			if err != nil {
				r.err = errors.Trace(err)
				close(output)
				return
			}

			//add handler to output channel
			output <- handler
		}
	}()

	return output, nil
}

//readMagicNumber check the binlog file header
func (r *Reader) readMagicNumber() error {
	//read magic number
	if buffer, err := r.cache.Read(4); err != nil {
		return errors.Trace(err)
	} else {
		//checkc the magic hader
		if magic := hack.String(buffer); magic != MagicNumber {
			return errInvalidBinlog
		} else {
			r.pos = 4
			return nil
		}
	}
}

//readHeader parse the header and return the header
func (r *Reader) readHeader() (*Header, error) {
	if buffer, err := r.cache.Read(int(EventHeaderLen)); err != nil {
		return nil, errors.Trace(err)
	} else {
		r.pos += EventHeaderLen
		return NewHeader(buffer), nil
	}
}

//readHeader parse the header and return the header
func (r *Reader) readPayload(h *Header) ([]byte, error) {
	//calculate the next event pos and check binlogversion
	pos := h.Pos()
	if buffer, err := r.cache.Read(int(pos - r.pos)); err != nil {
		return nil, errors.Trace(err)
	} else {
		r.pos = pos
		return buffer, nil
	}
}

//By the time you read the first event from the log you don't know what binlog version the binlog has.
//To determine the version correctly it has to be checked if the first event is:
//
//	- FORMAT_DESCRIPTION_EVENT version = 4
//  - START_EVENT_V3(not support)
//		- if event-size == 13 + 56: version = 1
//		- if event-size == 19 + 56: version = 3
//		- otherwise: invalid
func (r *Reader) readFormatDescriptionEvent() error {
	//check header
	header, err := r.readHeader()
	if err != nil {
		return errors.Trace(err)
	} else {
		if header.Kind() != FORMAT_DESCRIPTION_EVENT {
			return errors.Errorf("Binlog:expect a FORMAT_DESCRIPTION_EVENT event.")
		}
	}

	//calculate the next event pos and check binlogversion
	if buffer, err := r.readPayload(header); err != nil {
		return errors.Trace(err)
	} else {
		if binlogVersion, err := binary.ReadInt2(buffer); err != nil {
			return errors.Annotate(err, "Binlog:fail to read binlog version,the error is ")
		} else {
			if !testBinlogVersionV4(binlogVersion) {
				return errInvalidVersion
			}
		}
	}
	return nil
}

//Close closes cache
func (r *Reader) Close() {
	r.cache.Close()
}

//Error return the error
func (r *Reader) Error() error {
	return r.err
}
