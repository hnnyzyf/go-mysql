package replication

import (
	"fmt"
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
		Host:   "192.168.195.149:3306",
		User:   "replication",
		Passwd: "123456",

		ServerId: 100,

		FileName: "mysql-bin.000001",
		Position: 4,
	}

	err := s.d.ChangeMaster(cfg)
	fmt.Println(err)
	c.Assert(err, IsNil)

	output, err := s.d.Start()
	fmt.Println(err)
	c.Assert(err, IsNil)

	for {
		if handler, ok := <-output; !ok {
			//c.Assert(s.d.err, IsNil)
			//s.d.Stop()
			break
		} else {
			fmt.Println(handler)
			_ = handler.Accept()
		}
	}

}
