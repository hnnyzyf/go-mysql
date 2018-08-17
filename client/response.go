package client

import (
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

//isOkPacket tests whether current packet is ok packet
func isOkPacket(buffer []byte) bool {
	if header, err := binary.ReadInt1(buffer); err != nil {
		return false
	} else {
		return testOkPacket(header)
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
		return testEofPacket(header)
	}
}

//IsResultPacket tests whether current packet is a ProtocolText::Resultset packet
func isResultPacket(buffer []byte) bool {
	if _, header, err := binary.ReadLengthEncodedInteger(buffer); err != nil && header > 0 {
		return false
	} else {
		return testResultPacket(header)
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
		return 0, errors.Errorf("Client:still not support local infile query")
	default:
		return 0, errors.Errorf("Client:not a valid query response meta packet")
	}
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
	d:=new(definition)
	payload:=binary.NewBuffer(buffer)

	//read catalog
	if _,err:=payload.ReadLengthEncodedString();err!=nil{
		return errors.Trace(err)
	}

	//read 

}

//readColumnPacket41 read definition from buffer
func readColumnPacket320(buffer []byte) (*definition, error) {
	return nil,errors.Trace("Client:not support ColumnDefinition320 packet")
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

//IsEofPacket is a wrapper,it uses different methods according to the  CLIENT_DEPRECATE_EOF flag
func (s *Session) IsResultPacket(buffer []byte) bool {
	return isResultPacket(buffer)
}

//readErrPacket parse the payload of errPacket
func (s *Session) ReadErrPacket(buffer []byte) error {
	return readErrPacket(buffer)
}

//ReadComQueryResponse reads the result from connection
func (s *Session) ReadComQueryResponse() error {
	//read the meta packet
	if count, err := s.ReadMetaPacket(); err != nil {
		return errors.Trace(ok)
	} else if count == 0 {
		return nil
	} else {
		//read the column definition
		if err := s.ReadColumnPacket(count); err != nil {
			return errors.Trace(err)
		}
	}

	//try to read result set
	go s.ReadResultSet()

	return nil
}

//ReadMetaPacket read meta packet
func (s *Session) ReadMetaPacket() (uint64, error) {
	if buffer, err := s.ReadPacket(); err != nil {
		return 0, errors.Trace(err)
	} else {
		return readMetaPacket(buffer)
	}
}

//ReadMetaPacket read meta packet
func (s *Session) ReadColumnPacket(count uint64) error {
	s.def = make([]*definition, count)
	var i uint64
	readColumnPacket := readColumnPacket41
	if !testProtocol41(s.capabilities) {
		readColumnPacket = readColumnPacket320
	}

	for i = 0; i < count; i++ {
		if buffer, err := s.ReadPacket(); err != nil {
			return errors.Trace(err)
		} else {
			if definition, err := readColumnPacket(buffer); err != nil {
				return errors.Trace(err)
			} else {
				s.def[i] = definition
			}
		}
	}
	return nil
}
