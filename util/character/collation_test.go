package character

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyCollationSuite struct{}

var _ = Suite(&MyCollationSuite{})

func (s *MyCollationSuite) TestCollation(c *C) {
	var i uint8 = 1
	for k, v := range collation {
		c.Assert(k, Equals, i)
		i++
	}
}
