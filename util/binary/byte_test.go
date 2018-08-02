package binary

import (
	//"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyPayloadSuite struct{}

var _ = Suite(&MyPayloadSuite{})

func (s *MyPayloadSuite) TestWriteInteger(c *C) {
	b := make([]byte, 24)
	p := NewWriter(b)
	p.WriteInt1(0x01)
	p.WriteInt2(0x0203)
	p.WriteInt3(0x040506)
	p.WriteInt4(0x0708090A)
	p.WriteInt6(0x123456789ABC)
	p.WriteInt8(0x1F2F3F4F5F6F7F8F)

	res := []byte{0x01, 0x03, 0x02, 0x06, 0x05, 0x04, 0x0A, 0x09, 0x08, 0x07, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}

	for i := range res {
		c.Assert(res[i], Equals, b[i])
	}
}

func (s *MyPayloadSuite) TestReadInteger(c *C) {
	b := make([]byte, 24)
	p := NewWriter(b)
	p.WriteInt1(0x01)
	p.WriteInt2(0x0203)
	p.WriteInt3(0x040506)
	p.WriteInt4(0x0708090A)
	p.WriteInt6(0x123456789ABC)
	p.WriteInt8(0x1F2F3F4F5F6F7F8F)

	p = NewReader(b)

	c.Assert(p.ReadInt1(), Equals, uint8(0x01))
	c.Assert(p.ReadInt2(), Equals, uint16(0x0203))
	c.Assert(p.ReadInt3(), Equals, uint32(0x040506))
	c.Assert(p.ReadInt4(), Equals, uint32(0x0708090A))
	c.Assert(p.ReadInt6(), Equals, uint64(0x123456789ABC))
	c.Assert(p.ReadInt8(), Equals, uint64(0x1F2F3F4F5F6F7F8F))
}

func (s *MyPayloadSuite) TestWriteStringWithNull(c *C) {
	st := []byte("我们做个测试")
	b := make([]byte, len(st)+1)
	p := NewWriter(b)

	p.WriteStringWithNull(st)

	res := []byte{230, 136, 145, 228, 187, 172, 229, 129, 154, 228, 184, 170, 230, 181, 139, 232, 175, 149, 0}

	for i := range res {
		c.Assert(res[i], Equals, b[i])
	}
}

func (s *MyPayloadSuite) TestReadStringWithNull(c *C) {
	st := []byte("我们做个测试")
	b := make([]byte, len(st)+1)
	p := NewWriter(b)
	p.WriteStringWithNull(st)

	res := "我们做个测试"
	p = NewReader(b)
	c.Assert(string(p.ReadStringWithNull()), Equals, res)
}
