package client

import (
	. "gopkg.in/check.v1"
	"io"
)

type MyResponseSuite struct{}

var _ = Suite(&MyResponseSuite{})

func (s *MyResponseSuite) TestGenericPacket(c *C) {
	data := [][]byte{
		[]byte{},
		[]byte{0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00},
		[]byte{0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[]byte{0xff, 0x00},
		[]byte{0x00, 0xff},
		[]byte{0xfe, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00},
	}

	res := []struct {
		t  func([]byte) bool
		ok bool
	}{
		{isOkPacket, false},
		{isOkPacket, true},
		{isOkPacket, false},
		{isErrPacket, true},
		{isErrPacket, false},
		{isEofPacket, false},
	}

	for i := range data {
		c.Assert(res[i].t(data[i]), Equals, res[i].ok)
	}
}

func next(res *response) ([]byte, error) {
	if res, ok := <-res.output; ok {
		return res, nil
	}

	if e, ok := <-res.err; ok {
		return nil, e
	}

	if res.eof {
		return nil, io.EOF
	}

	return nil, nil
}

func (s *MyResponseSuite) TestQueryResponsePacket(c *C) {
	query := "select * from test"
	conn, err := Connect("127.0.0.1:3306", "test", "123456", "test")
	c.Assert(err, IsNil)
	defer conn.Close()

	err = conn.RequestComQuery(query)
	c.Assert(err, IsNil)

	err = conn.ReadComQueryResponse()
	c.Assert(err, IsNil)

	res := [][]byte{
		[]byte{0x01, 0x31, 0x01, 0x32},
		[]byte{0x01, 0x33, 0x02, 0x31, 0x30},
	}

	i := 0
	for {
		if buffer, err := next(conn.res); err != nil {
			if err != io.EOF {
				c.Error(err)
			}
			break
		} else {
			c.Assert(buffer, DeepEquals, res[i])
		}
		i++
	}

}
