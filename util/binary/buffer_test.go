package binary

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyBufferSuite struct{}

var _ = Suite(&MyBufferSuite{})

func (s *MyBufferSuite) TestWriteInteger(c *C) {
	b := make([]byte, 24)
	p := NewBuffer(b)

	err := p.WriteInt1(0x01)
	c.Assert(err, Equals, nil)

	err = p.WriteInt2(0x0203)
	c.Assert(err, Equals, nil)

	err = p.WriteInt3(0x040506)
	c.Assert(err, Equals, nil)

	err = p.WriteInt4(0x0708090A)
	c.Assert(err, Equals, nil)

	err = p.WriteInt6(0x123456789ABC)
	c.Assert(err, Equals, nil)

	err = p.WriteInt8(0x1F2F3F4F5F6F7F8F)
	c.Assert(err, Equals, nil)

	res := []byte{0x01, 0x03, 0x02, 0x06, 0x05, 0x04, 0x0A, 0x09, 0x08, 0x07, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}

	for i := range res {
		c.Assert(res[i], Equals, b[i])
	}

}

func (s *MyBufferSuite) TestReadInteger(c *C) {
	b := []byte{0x01, 0x03, 0x02, 0x06, 0x05, 0x04, 0x0A, 0x09, 0x08, 0x07, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12,
		0x8F, 0x7F, 0x6F, 0x5F, 0x4F, 0x3F, 0x2F, 0x1F}

	p := NewBuffer(b)

	res1, err := p.ReadInt1()
	c.Assert(err, Equals, nil)
	c.Assert(res1, Equals, uint8(0x01))

	res2, err := p.ReadInt2()
	c.Assert(err, Equals, nil)

	c.Assert(res2, Equals, uint16(0x0203))

	res3, err := p.ReadInt3()
	c.Assert(err, Equals, nil)

	c.Assert(res3, Equals, uint32(0x040506))

	res4, err := p.ReadInt4()
	c.Assert(err, Equals, nil)

	c.Assert(res4, Equals, uint32(0x0708090A))

	res5, err := p.ReadInt6()
	c.Assert(err, Equals, nil)

	c.Assert(res5, Equals, uint64(0x123456789ABC))

	res6, err := p.ReadInt8()
	c.Assert(err, Equals, nil)

	c.Assert(res6, Equals, uint64(0x1F2F3F4F5F6F7F8F))

}

func (s *MyBufferSuite) TestWriteStringWithNull(c *C) {
	st := []byte("我们做个测试")
	b := make([]byte, len(st)+1)
	p := NewBuffer(b)

	err := p.WriteStringWithNull(st)
	c.Assert(err, Equals, nil)

	res := []byte{230, 136, 145, 228, 187, 172, 229, 129, 154, 228, 184, 170, 230, 181, 139, 232, 175, 149, 0}

	for i := range res {
		c.Assert(res[i], Equals, b[i])
	}

}
func (s *MyBufferSuite) TestReadStringWithNull(c *C) {
	b := []byte{230, 136, 145, 228, 187, 172, 229, 129, 154, 228, 184, 170, 230, 181, 139, 232, 175, 149, 0}
	p := NewBuffer(b)

	res := "我们做个测试"
	st, err := p.ReadStringWithNull()
	c.Assert(err, Equals, nil)

	c.Assert(string(st), Equals, res)

}

func (s *MyBufferSuite) TestLengthEncodeInteger(c *C) {
	b := make([]byte, 1+3+4+9)
	data := []struct {
		val uint64
		n   int
	}{
		{0x01, 1},
		{0x0203, 3},
		{0x040506, 4},
		{0x0708091011121314, 9},
	}

	for i := range data {
		val := data[i].val
		n := data[i].n
		c.Assert(LengthOfInteger(val), Equals, n)
	}

	res := []byte{0x01, 0xfc, 0x03, 0x02, 0xfd, 0x06, 0x05, 0x04, 0xfe, 0x14, 0x13, 0x12, 0x11, 0x10, 0x09, 0x08, 0x07}
	p := NewBuffer(b)
	for i := range data {
		val := data[i].val
		n := data[i].n
		l, err := p.WriteLengthEncodedInteger(val)
		c.Assert(err, Equals, nil)
		c.Assert(l, Equals, n)
	}

	c.Assert(b, DeepEquals, res)
	p.off = 0

	for i := range data {
		val, n, err := p.ReadLengthEncodedInteger()
		c.Assert(err, Equals, nil)
		c.Assert(n, Equals, data[i].n)
		c.Assert(val, Equals, data[i].val)
	}

}
