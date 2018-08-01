package hack

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestString(c *C) {
	b := []byte{48, 49, 50, 51}
	ps := String(b)

	c.Assert(ps, Equals, "0123")
}

func (s *MySuite) TestSlice(c *C) {
	ps := "0123"
	b := Slice(ps)
	res := []byte{48, 49, 50, 51}

	for i := range b {
		c.Assert(b[i], Equals, res[i])
	}
}
