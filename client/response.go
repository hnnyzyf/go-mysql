package client

import (
	"fmt"
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

//isOkPacket tests whether current packet is ok packet
func isOkPacket(buffer []byte) bool {
	if header, err := binary.ReadInt1(buffer); err != nil {
		return false
	} else {
		return testOkPacket(header) && len(buffer) > 7-4
	}
}

//isErrPacket tests whether current packet is err packet
func isErrPacket(buffer []byte) bool {
	if header, err := binary.ReadInt1(buffer); err != nil {
		return false
	} else {
		return testErrPacket(header)
	}
}

//IsEofPacket tests whether current packet is eof packet
func isEofPacket(buffer []byte) bool {
	if header, err := binary.ReadInt1(buffer); err != nil {
		return false
	} else {
		return testEofPacket(header) && len(buffer) < 9-4
	}
}

//readErrPacket parse the payload of errPacket
func readErrPacket(buffer []byte) error {
	payload := binary.NewBuffer(buffer)

	//code
	code, err := payload.ReadInt2()
	if err != nil {
		return errERRPacket
	}

	//marker
	if _, err = payload.ReadStringWithFixLen(1); err != nil {
		return errERRPacket
	}

	//skip
	if err := payload.Skip(1); err != nil {
		return errERRPacket
	}
	//SQL State
	if _, err := payload.ReadStringWithFixLen(5); err != nil {
		return errERRPacket
	}

	//human readable error message
	msg, err := payload.ReadStringWithEOF()
	if err != nil {
		return errERRPacket
	}

	return errors.Errorf("Client:error(%d):%s", code, hack.String(msg))
}

//IsOkPacket is a wrapper for IsOkPacket
func (s *Session) IsOkPacket(buffer []byte) bool {
	return isOkPacket(buffer)
}

//IsErrPacket is a wrapper for IsErrPacket
func (s *Session) IsErrPacket(buffer []byte) bool {
	return isErrPacket(buffer)
}

//IsEofPacket is a wrapper,it uses different methods according to the  CLIENT_DEPRECATE_EOF flag
func (s *Session) IsEofPacket(buffer []byte) bool {
	if testDeprecateEof(s.capabilities) {
		return isOkPacket(buffer)
	} else {
		return isEofPacket(buffer)
	}
}

//readErrPacket parse the payload of errPacket
func (s *Session) ReadErrPacket(buffer []byte) error {
	return readErrPacket(buffer)
}

//isInFilePacket packet
func isInFilePacket(buffer []byte) bool {
	return true
}

//IsResultPacket tests whether current packet is a ProtocolText::Resultset packet
func isResultPacket(buffer []byte) bool {
	if val, header, err := binary.ReadLengthEncodedInteger(buffer); err != nil && header > 0 {
		return false
	} else {
		return testResultPacket(val)
	}
}

//reeaMetapacket read the real count from buffer
func readMetaPacket(buffer []byte) (uint64, error) {
	switch {
	case isOkPacket(buffer):
		return 0, nil
	case isResultPacket(buffer):
		val, _, err := binary.ReadLengthEncodedInteger(buffer)
		return val, err
	case isErrPacket(buffer):
		return 0, readErrPacket(buffer)
	case isInFilePacket(buffer):
		return 0, errors.Errorf("Client:not support local infile query")
	default:
		return 0, errors.Errorf("Client:not a valid query response meta packet")
	}
}

type response struct {
	//For result to use
	output chan []byte
	err    chan error

	//for column definition
	def []*definition

	//no recoard to read
	eof bool
}

//defination repsent the basice inforamtion for a column in result set
type definition struct {
	name []byte
	// is the column character set and is defined in Protocol::CharacterSet.
	charset uint16
	//maximum length of the field
	length uint32
	//type of the column as defined in Column Type
	column uint8
	//max shown decimal digits
	//		0x00 for integers and static strings
	//		0x1f for dynamic strings, double, float
	//		0x00 to 0x51 for decimals
	decimals uint8
}

const CataLog string = "def"

//readColumnPacket41 read definition from buffer
func readColumnPacket41(buffer []byte) (*definition, error) {
	//init
	d := new(definition)
	payload := binary.NewBuffer(buffer)

	//read catalog
	if catalog, err := payload.ReadLengthEncodedString(); err != nil {
		return nil, errors.Trace(err)
	} else {
		if hack.String(catalog) != CataLog {
			return nil, errors.Errorf("client:expect catalog:def")
		}
	}

	//read schema
	if _, err := payload.ReadLengthEncodedString(); err != nil {
		return nil, errors.Trace(err)
	}

	//read table
	if _, err := payload.ReadLengthEncodedString(); err != nil {
		return nil, errors.Trace(err)
	}

	//read org_table
	if _, err := payload.ReadLengthEncodedString(); err != nil {
		return nil, errors.Trace(err)
	}

	//read name
	if name, err := payload.ReadLengthEncodedString(); err != nil {
		return nil, errors.Trace(err)
	} else {
		d.name = name
	}

	//read org_name
	if _, err := payload.ReadLengthEncodedString(); err != nil {
		return nil, errors.Trace(err)
	}

	//read length of fixed-length field
	if val, _, err := payload.ReadLengthEncodedInteger(); err != nil {
		return nil, errors.Trace(err)
	} else {
		if val != 0x0c {
			return nil, errors.Errorf("client:expect fix length fileds:0x0c")
		}
	}

	//read character set
	if charset, err := payload.ReadInt2(); err != nil {
		return nil, errors.Trace(err)
	} else {
		d.charset = charset
	}

	//read column length
	if _, err := payload.ReadInt4(); err != nil {
		return nil, errors.Trace(err)
	}

	//read type
	if column, err := payload.ReadInt1(); err != nil {
		return nil, errors.Trace(err)
	} else {
		d.column = column
	}

	//read decimal
	if decimals, err := payload.ReadInt1(); err != nil {
		return nil, errors.Trace(err)
	} else {
		d.decimals = decimals
	}

	//skip  filler
	if err := payload.Skip(2); err != nil {
		return nil, errors.Trace(err)
	}

	return d, nil
}

//readColumnPacket41 read definition from buffer
func readColumnPacket320(buffer []byte) (*definition, error) {
	return nil, errors.Errorf("Client:not support ColumnDefinition320 packet")
}

//ReadComQueryResponse reads the result from connection
func (s *Session) ReadComQueryResponse() error {
	//read the meta packet
	if count, err := s.ReadMetaPacket(); err != nil {
		return errors.Trace(err)
	} else if count == 0 {
		return nil
	} else {
		//read the column definition
		if err := s.ReadColumnPacket(count); err != nil {
			return errors.Trace(err)
		}
	}

	//init the result channel
	s.res.output = make(chan []byte)
	s.res.err = make(chan error)

	//try to read result set
	go s.ReadResultSet()

	return nil
}

//ReadMetaPacket read meta packet
func (s *Session) ReadMetaPacket() (uint64, error) {
	if buffer, err := s.ReadPacket(); err != nil {
		return 0, errors.Trace(err)
	} else {
		//A packet containing a Protocol::LengthEncodedInteger column_count
		return readMetaPacket(buffer)
	}
}

//ReadMetaPacket read meta packet
func (s *Session) ReadColumnPacket(count uint64) error {
	s.res.def = make([]*definition, count)

	readColumnPacket := readColumnPacket41
	if !testProtocol41(s.capabilities) {
		readColumnPacket = readColumnPacket320
	}

	//column_count * Protocol::ColumnDefinition packets
	var i uint64
	for i = 0; i < count; i++ {
		if buffer, err := s.ReadPacket(); err != nil {
			return errors.Trace(err)
		} else {
			if definition, err := readColumnPacket(buffer); err != nil {
				return errors.Trace(err)
			} else {
				s.res.def[i] = definition
			}
		}
	}

	//If the CLIENT_DEPRECATE_EOF client capability flag is not set, EOF_Packet
	if !testDeprecateEof(s.capabilities) {
		if buffer, err := s.ReadPacket(); err != nil {
			return errors.Trace(err)
		} else {
			if !isEofPacket(buffer) {
				return errors.Errorf("client:not accept an EOF packet after recieving column definition packets")
			}
		}
	}

	return nil
}

//ReadResultSet read the result packet one by one
func (s *Session) ReadResultSet() {
	i := 0
	for {
		if buffer, err := s.ReadPacket(); err != nil {
			fmt.Println("fail", i)
			s.res.err <- errors.Trace(err)
			return
		} else {
			b := buffer
			switch {
			//if err occurs,return error
			case s.IsErrPacket(b):
				s.res.err <- s.ReadErrPacket(b)
				return
			//According to https://dev.mysql.com/doc/internals/en/com-query-response.html
			//ERR_Packet in case of error. Otherwise:
			//If the CLIENT_DEPRECATE_EOF client capability flag is set,
			//OK_Packet; else EOF_Packet.
			case s.IsEofPacket(b):
				s.res.eof = true
				close(s.res.err)
				close(s.res.output)
				return
				//read result set
			default:
				//do a copy instead
				s.res.output <- b
			}
		}
	}
}
