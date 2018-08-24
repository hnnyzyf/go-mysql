package client

import (
	. "gopkg.in/check.v1"
	"time"
)

type MyDriverSuite struct {
}

var _ = Suite(&MyDriverSuite{})

func (s *MyDriverSuite) TestParseValue(c *C) {
	data := []struct {
		val interface{}
		res string
	}{
		{+10, "10"},
		{int8(-10), "-10"},
		{int16(-1000), "-1000"},
		{int32(-3232), "-3232"},
		{int64(-10323232300), "-10323232300"},
		{uint8(0xff), "255"},
		{uint16(0xffff), "65535"},
		{uint32(0x1111ffff), "286392319"},
		{uint64(0x22221111ffff), "37529710624767"},
		{true, "true"},
		{false, "false"},
		{1.23456, "1.23456"},
		{-1.23456, "-1.23456"},
		{-1.23456333333333333333, "-1.2345633333333332"},
		{"dsadasdasdasd", `"dsadasdasdasd"`},
		{[]byte{0x01, 0x02, 0x03}, string([]byte{0x01, 0x02, 0x03})},
		{time.Date(2009, time.November, 10, 23, 0, 2, 2, time.UTC), "2009-11-10 23:00:02.000000002"},
		{time.Date(2009, time.November, 10, 23, 0, 2, 0, time.UTC), "2009-11-10 23:00:02.000000000"},
	}

	for i := range data {
		r, err := parseValue(data[i].val)
		c.Assert(err, IsNil)
		c.Assert(r, Equals, data[i].res)
	}
}
