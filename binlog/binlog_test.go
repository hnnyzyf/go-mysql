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
	r := NewReader()
	return &MyBinlogSuite{
		r: r,
	}
}

var _ = Suite(newSuite())

func (s *MyBinlogSuite) TestParse(c *C) {
	output, err := s.r.Open("./bin/mysql-bin.000001")
	defer s.r.Close()
	c.Assert(err, IsNil)
	for {
		handler, ok := <-output
		if !ok {
			c.Log(s.r.Error())
			break
		}
		//parese the hadler
		_ = handler.Accept()
	}

}
