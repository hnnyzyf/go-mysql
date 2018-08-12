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
