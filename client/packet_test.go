package client

import (
	. "gopkg.in/check.v1"
)

type MyPacketSuite struct{}

var _ = Suite(&MyPacketSuite{})

func (s *MyPacketSuite) TestPacket(c *C) {
	data := [][]byte{
		[]byte{},
		[]byte{0x00, 0xff},
		[]byte{0xff, 0x00},
		[]byte{0xff, 0x00},
		[]byte{0x00, 0xff},
	}

	res := []struct {
		t  func([]byte) bool
		ok bool
	}{
		{IsOkPacket, false},
		{IsOkPacket, true},
		{IsOkPacket, false},
		{IsErrPacket, true},
		{IsErrPacket, false},
	}

	for i := range data {
		c.Assert(res[i].t(data[i]), Equals, res[i].ok)
	}
}
