package utf8mb4

import (
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestReadRune(c *C) {
	b := []byte{0xE5, 0xBE, 0x88, 0xE5, 0xB1, 0x8C, 0xF7, 0xBE, 0xBC, 0xB4}
	e := NewEncoder(b)

	res := []struct {
		r   []byte
		err error
	}{
		{[]byte{0xE5, 0xBE, 0x88}, nil},
		{[]byte{0xE5, 0xB1, 0x8C}, nil},
		{[]byte{0xF7, 0xBE, 0xBC, 0xB4}, nil},
		{nil, io.EOF},
		{nil, io.EOF},
	}

	for i := range res {
		r, err := e.ReadRune()
		c.Assert(r, DeepEquals, res[i].r)
		c.Assert(err, Equals, res[i].err)
	}

}
