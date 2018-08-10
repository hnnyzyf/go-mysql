package binlog

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyBinlogSuite struct {
	r *reader
}

func newSuite() *MyBinlogSuite {
	r, _ := NewReader("./bin/mysql-bin.000001")
	return &MyBinlogSuite{
		r: r,
	}
}

var _ = Suite(newSuite())

func (s *MyBinlogSuite) TestParse(c *C) {
	err := s.r.ParseFile()
	c.Log(err)
	c.Assert(err, IsNil)
}
