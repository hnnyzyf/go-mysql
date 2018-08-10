package protocol

import (
	"bytes"
	"io"
	"net"
	"testing"
	"time"

	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/math"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyConnSuite struct{}

var _ = Suite(&MyConnSuite{})

func (s *MyConnSuite) create(b []byte) *Conn {
	c := newConn(b)
	return NewConn(c)
}

func (s *MyConnSuite) getByte(c *Conn) []byte {
	return c.conn.(*testconn).b.Bytes()
}

type testconn struct {
	b *bytes.Buffer
}

func newConn(b []byte) *testconn {
	return &testconn{
		b: bytes.NewBuffer(b),
	}
}

func (t *testconn) Read(b []byte) (int, error) {
	return t.b.Read(b)
}

func (t *testconn) Write(b []byte) (int, error) {
	return t.b.Write(b)
}

func (t *testconn) Close() error {
	return nil
}

func (t *testconn) LocalAddr() net.Addr {
	return nil
}

func (t *testconn) RemoteAddr() net.Addr {
	return nil
}

func (t *testconn) SetDeadline(stamp time.Time) error {
	return nil
}

func (t *testconn) SetReadDeadline(stamp time.Time) error {
	return nil
}

func (t *testconn) SetWriteDeadline(stamp time.Time) error {
	return nil
}

func (s *MyConnSuite) TestRead(c *C) {
	data := [][]byte{
		[]byte{0x0f, 0x00, 0x00, 0x00, 0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x0E, 0x0D, 0x0C, 0x0B, 0x0A, 0x09, 0x0F},
	}

	res := []struct {
		size int
		id   uint8
		p    []byte
	}{
		{15, 0, []byte{0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x0E, 0x0D, 0x0C, 0x0B, 0x0A, 0x09, 0x0F}},
	}

	for i := range res {
		r := s.create(data[i])
		p, err := r.ReadPacket()
		c.Assert(err, Equals, nil)
		c.Assert(len(p), Equals, res[i].size)
		c.Assert(p, DeepEquals, res[i].p)
	}

}

func (s *MyConnSuite) TestWrite(c *C) {
	data := []struct {
		size int
		id   uint8
		p    io.Reader
	}{
		{15, 5, bytes.NewBuffer([]byte{0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x0E, 0x0D, 0x0C, 0x0B, 0x0A, 0x09, 0x0F})},
	}

	res := []struct {
		id      uint8
		payload []byte
	}{
		{6, []byte{0x0f, 0x00, 0x00, 0x05, 0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x0E, 0x0D, 0x0C, 0x0B, 0x0A, 0x09, 0x0F}},
	}

	for i := range res {
		conn := s.create([]byte{})
		conn.sid = data[i].id
		err := conn.WritePacket(data[i].size, data[i].p)
		c.Assert(err, IsNil)
		c.Assert(s.getByte(conn), DeepEquals, res[i].payload)
		c.Assert(conn.sid, Equals, res[i].id)
	}

}

func bigPacket(size int) io.Reader {
	return bytes.NewBuffer(make([]byte, size))
}

func bigBuffer(size int, sid uint8) []byte {
	buffer := bytes.NewBuffer([]byte{})
	for size > 0 {
		length := math.MinInt(size, int(PayLoadMaxLen))
		buffer.Write(binary.StoreInt3(uint32(length)))
		sid = sid + 1
		buffer.Write(binary.StoreInt1(sid))

		buffer.Write(make([]byte, length))
		size -= length
	}
	return buffer.Bytes()
}

func (s *MyConnSuite) TestWriteBigPacket(c *C) {
	data := []struct {
		size int
		sid  uint8
		res  uint8
	}{
		{2*int(PayLoadMaxLen) + 1, 255, 2},
		{2*int(PayLoadMaxLen) + 1, 0, 3},
		{8*int(PayLoadMaxLen) + 76543, 253, 6},
	}

	for i := range data {
		conn := s.create([]byte{})
		conn.sid = data[i].sid
		packet := bigPacket(data[i].size)
		b := bigBuffer(data[i].size, data[i].sid)
		err := conn.WritePacket(data[i].size, packet)
		c.Assert(err, IsNil)
		c.Assert(len(s.getByte(conn)), DeepEquals, len(b))
		c.Assert(conn.sid, DeepEquals, data[i].res)
	}
}
