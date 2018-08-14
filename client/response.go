package client

import (
	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/hack"
	"github.com/juju/errors"
)

//IsOkPacket tests whether current packet is ok packet
func IsOkPacket(buffer []byte) bool {
	if header, err := binary.ReadInt1(buffer); err != nil {
		return false
	} else {
		return testOkPacket(header)
	}
}

//IsErrpacket tests whether current packet is err packet
func IsErrPacket(buffer []byte) bool {
	if header, err := binary.ReadInt1(buffer); err != nil {
		return false
	} else {
		return testErrPacket(header)
	}
}

//IsEofPacket tests whether current packet is eof packet
func IsEofPacket(buffer []byte) bool {
	if header, err := binary.ReadInt1(buffer); err != nil {
		return false
	} else {
		return testEofPacket(header)
	}
}

//IsResultPacket tests whether current packet is a ProtocolText::Resultset packet
func IsResultPacket(buffer []byte) bool {
	if _, header, err := binary.ReadLengthEncodedInteger(buffer); err != nil && header > 0 {
		return false
	} else {
		return testResultPacket(header)
	}
}

//readErrPacket parse the payload of errPacket
func ReadErrPacket(buffer []byte) error {
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


//ReadResultPacket read the payload of resultpacket
func ReadResultPacket(buffer []byte) error {
	
}