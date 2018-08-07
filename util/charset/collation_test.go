package charset

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyCollationSuite struct{}

var _ = Suite(&MyCollationSuite{})

func (s *MyCollationSuite) TestCollation(c *C) {
	data := []uint8{
		0x21,
	}

	res := []Collation{
		{"utf8_general_ci", 0x21, "utf8", true},
	}

	for i := range data {
		x := collations[data[i]]
		c.Assert(x.Name, Equals, res[i].Name)
		c.Assert(x.ID, Equals, res[i].ID)
		c.Assert(x.Charset, Equals, res[i].Charset)
		c.Assert(x.IsDefault, Equals, res[i].IsDefault)
	}

}
