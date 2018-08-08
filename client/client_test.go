package client

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyClientSuite struct {
	conn *session
}

var _ = Suite(&MyClientSuite{})

func (s *MyClientSuite) TestConnect(c *C) {
	data := []struct {
		h string
		u string
		p string
		d string
		e error
	}{
		{"127.0.0.1:3306", "test", "123456", "test", nil},
	}

	for i := range data {
		l := data[i]
		conn, err := Connect(l.h, l.u, l.p, l.d)
		fmt.Println(err)
		c.Assert(err, Equals, nil)
		conn.Close()
	}
}
