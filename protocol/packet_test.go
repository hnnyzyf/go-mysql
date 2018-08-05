package protocol

import (
	"bytes"
	"testing"

	"github.com/hnnyzyf/go-mysql/util/binary"
	"github.com/hnnyzyf/go-mysql/util/math"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyPacketSuite struct{}

var _ = Suite(&MyPacketSuite{})

func (s *MyPacketSuite) TestRead(c *C) {
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
		r := bytes.NewBuffer(data[i])
		p, err := ReadPacket(r, 0)
		c.Assert(err, Equals, nil)
		c.Assert(p.ReadLength(), Equals, res[i].size)
		c.Assert(p.ReadId(), Equals, res[i].id)

		payload, err := p.ReadPayload()
		c.Assert(err, Equals, nil)
		c.Assert(payload.Len(), Equals, p.ReadLength())
		c.Assert(payload.Bytes(), DeepEquals, res[i].p)
	}

}

func (s *MyPacketSuite) TestWrite(c *C) {
	data := []*Packet{
		&Packet{
			length:  15,
			sid:     0,
			payload: bytes.NewBuffer([]byte{0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x0E, 0x0D, 0x0C, 0x0B, 0x0A, 0x09, 0x0F}),
		},
	}

	res := []struct {
		sid     uint8
		payload []byte
	}{
		{1, []byte{0x0f, 0x00, 0x00, 0x01, 0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x0E, 0x0D, 0x0C, 0x0B, 0x0A, 0x09, 0x0F}},
	}

	for i := range res {

		w := bytes.NewBuffer([]byte{})
		sid, err := WritePacket(w, data[i])
		c.Assert(err, Equals, nil)
		c.Assert(w.Bytes(), DeepEquals, res[i].payload)
		c.Assert(sid, Equals, res[i].sid)
	}

}

func bigPacket(size int, sid uint8) *Packet {
	return &Packet{
		length:  size,
		sid:     sid,
		payload: bytes.NewBuffer(make([]byte, size)),
	}
}

func bigBuffer(size int, sid uint8) ([]byte, uint8) {
	buffer := bytes.NewBuffer([]byte{})

	for size > 0 {
		length := math.MinInt(size, int(PayLoadMaxLen))
		buffer.Write(binary.StoreInt3(uint32(length)))
		sid = sid + 1
		buffer.Write(binary.StoreInt1(sid))

		buffer.Write(make([]byte, length))
		size -= length
	}

	return buffer.Bytes(), sid
}

func (s *MyPacketSuite) TestWriteBigPacket(c *C) {
	data := []struct {
		size int
		sid  uint8
	}{
		{2*int(PayLoadMaxLen) + 1, 5},
		{2*int(PayLoadMaxLen) + 15647, 0},
		{15*int(PayLoadMaxLen) + 76543, 255},
	}

	for i := range data {
		w := bytes.NewBuffer([]byte{})
		packet := bigPacket(data[i].size, data[i].sid)
		b, id := bigBuffer(data[i].size, data[i].sid)

		sid, err := WritePacket(w, packet)
		c.Assert(err, Equals, nil)
		c.Assert(len(w.Bytes()), Equals, len(b))
		c.Assert(w.Bytes(), DeepEquals, b)
		c.Assert(sid, Equals, id)

	}

}
