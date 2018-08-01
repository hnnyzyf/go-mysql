package math

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestMaxMin(c *C) {
	var a uint8 = 122
	var b uint8 = 133

	c.Assert(MaxUint8(a, b), Equals, uint8(133))
	c.Assert(MinUint8(a, b), Equals, uint8(122))
}
