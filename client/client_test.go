package client

import (
	"fmt"
	"net"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyClientSuite struct {
	conn *session
}

func (s *MyClientSuite) init() error {
	host := "127.0.0.1:3306"
	raddr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return err
	}
	c, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return err
	}
	s.conn = &session{
		c:        c,
		username: "test",
		password: "123456",
		database: "test",
	}
	return nil
}

var _ = Suite(&MyClientSuite{})

func (s *MyClientSuite) TestReadHandshakeV10(c *C) {
	err := s.init()
	if err != nil {
		c.Error(err)
	}
	defer s.conn.Close()
	err = s.conn.ReadHandshakeV10()
	if err != nil {
		c.Error(err)
	}

	fmt.Println(s.conn.capabilities)
}

//func (s *MyClientSuite) TestWriteHandshakeResponse41(c *C) {
//	err := s.init()
//	if err != nil {
//		c.Error(err)
//	}
//	defer s.conn.Close()
//	err = s.conn.ReadHandshakeV10()
//	if err != nil {
//		c.Error(err)
//	}
//
//	err = s.conn.WriteHandshakeResponse41()
//	if err != nil {
//		c.Error(err)
//	}
//}
