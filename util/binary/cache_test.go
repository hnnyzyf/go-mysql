package binary

import (
	"bytes"
	"io"

	. "gopkg.in/check.v1"
)

type MyCacheSuite struct{}

var _ = Suite(&MyCacheSuite{})

func (s *MyCacheSuite) newReader() *bytes.Buffer {
	b := make([]byte, 4096)
	for i := range b {
		b[i] = uint8(i % 255)
	}
	return bytes.NewBuffer(b)
}

func (s *MyCacheSuite) TestRead(c *C) {
	r := s.newReader()
	cache := NewCacheWithSize(r, 256)
	b := r.Bytes()
	data := []struct {
		size int
		res  []byte
		err  error
	}{
		{10, b[0:10], nil},
		{100, b[10:110], nil},
		{2000, b[110:2110], nil},
		{890, b[2110:3000], nil},
		{1095, b[3000:4095], nil},
		{1, b[4095:], nil},
		{2, make([]byte, 2), io.EOF},
	}

	for i := range data {
		c.Log(data[i].size)
		buffer := make([]byte, data[i].size)
		err := cache.Read(buffer)
		c.Assert(buffer, DeepEquals, data[i].res)
		if data[i].err == nil {
			c.Assert(err, IsNil)
		} else {
			c.Assert(err.Error(), Equals, data[i].err.Error())
		}
	}
}

func (s *MyCacheSuite) TestNext(c *C) {
	r := s.newReader()
	cache := NewCacheWithSize(r, 256)
	b := r.Bytes()
	data := []struct {
		size int
		res  []byte
		err  error
	}{
		{10, b[0:10], nil},
		{100, b[10:110], nil},
		{2000, b[110:2110], nil},
		{890, b[2110:3000], nil},
		{1095, b[3000:4095], nil},
		{1, b[4095:], nil},
		{2, nil, io.EOF},
	}

	for i := range data {
		buffer, err := cache.Next(data[i].size)
		c.Assert(buffer, DeepEquals, data[i].res)
		if data[i].err == nil {
			c.Assert(err, IsNil)
		} else {
			c.Assert(err.Error(), Equals, data[i].err.Error())
		}
	}
}

func (s *MyCacheSuite) TestNextEOF(c *C) {
	r := s.newReader()
	cache := NewCacheWithSize(r, 256)
	data := []struct {
		size int
		res  []byte
		err  string
	}{
		{4097, nil, `.+enough.+`},
		{1000, nil, `EOF`},
	}

	for i := range data {
		buffer, err := cache.Next(data[i].size)
		c.Assert(buffer, DeepEquals, data[i].res)
		c.Assert(err, ErrorMatches, data[i].err)
	}
}
