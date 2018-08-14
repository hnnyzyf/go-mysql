package replication

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyDumperSuite struct {
	d *Dumper
}

var _ = Suite(&MyDumperSuite{
	d: NewDumper("dumper-test"),
})

func (s *MyDumperSuite) TestDumper(c *C) {
	cfg := &Config{
		Host:   "127.0.0.1:3306",
		User:   "replication",
		Passwd: "123456",

		ServerId: 100,

		FileName: "mysql-bin.000001",
		Position: 4,
	}

	err := s.d.ChangeMaster(cfg)
	c.Assert(err, IsNil)

	err = s.d.Start()
	c.Assert(err, IsNil)

	for {
		handler, err := s.d.Next()
		if err != nil {
			c.Error(err)
			c.Assert(err, IsNil)
		} else {
			_ = handler.Accept()
		}
	}

}
