package binlog

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyBinlogSuite struct {
	r *Reader
}

func newSuite() *MyBinlogSuite {
	r, _ := NewReader("./bin/mysql-bin.000001")
	return &MyBinlogSuite{
		r: r,
	}
}

var _ = Suite(newSuite())

func (s *MyBinlogSuite) TestParse(c *C) {
	output, stop, err := s.r.ParseFile()
	c.Assert(err, IsNil)

Done:
	for {
		select {
		case handler := <-output:
			_ = handler.Parse()
		case <-stop:
			c.Log(s.r.Error())
			break Done
		}
	}

}
