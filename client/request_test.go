package client

import (
	. "gopkg.in/check.v1"
)

type MyRequestSuite struct {
	host     string
	user1    string
	user2    string
	password string
	db       string
	ssl      bool
}

var _ = Suite(&MyRequestSuite{
	host:     "127.0.0.1:3306",
	user1:    "test",
	user2:    "testssl",
	password: "123456",
	db:       "test",
	ssl:      false,
})

func (s *MyRequestSuite) connect() (*Session, error) {
	return Connect(s.host, s.user1, s.password, s.db)
}

func (s *MyRequestSuite) checkResponse(conn *Session, testPacket func([]byte) bool) (bool, error) {
	buffer, err := conn.ReadPacket()
	if err != nil {
		return false, err
	}

	if conn.IsErrPacket(buffer) {
		return testPacket(buffer), conn.ReadErrPacket(buffer)
	} else {
		return testPacket(buffer), nil
	}
}

func (s *MyRequestSuite) TestBasicRequest(c *C) {
	conn, err := s.connect()
	c.Assert(err, IsNil)
	defer conn.Close()

	data := []struct {
		fu    func() error
		check func([]byte) bool
	}{
		{conn.RequestComSleep, conn.IsErrPacket},
		{conn.RequestComConnect, conn.IsErrPacket},
		{conn.RequestComPing, conn.IsOkPacket},
		{conn.RequestComTime, conn.IsErrPacket},
		{conn.RequestComDelayedInsert, conn.IsErrPacket},
		{conn.RequestComResetConnection, conn.IsOkPacket},
		{conn.RequestComDeamon, conn.IsErrPacket},
	}

	for i := range data {
		err := data[i].fu()
		c.Assert(err, IsNil)
		ok, _ := s.checkResponse(conn, data[i].check)
		c.Assert(ok, Equals, true)
	}
}

func (s *MyRequestSuite) TestOneParaRequest(c *C) {
	conn, err := s.connect()
	c.Assert(err, IsNil)
	defer conn.Close()

	data := []struct {
		fu    func(string) error
		d     string
		check func([]byte) bool
		err   string
	}{
		{conn.RequestComInitDB, "mysql", conn.IsOkPacket, ""},
		{conn.RequestComInitDB, "women", conn.IsOkPacket, ""},
		{conn.RequestComCreateDB, "women", conn.IsErrPacket, ".+known.+"},
		{conn.RequestComDropDB, "women", conn.IsErrPacket, ".+known.+"},
	}

	for i := range data {
		err := data[i].fu(data[i].d)
		c.Assert(err, IsNil)
		ok, err := s.checkResponse(conn, data[i].check)
		if data[i].err == "" {
			c.Assert(err, IsNil)
		} else {
			c.Assert(err, ErrorMatches, data[i].err)
		}
		c.Assert(ok, Equals, true)
	}
}

func (s *MyRequestSuite) TestQueryRequest(c *C) {
	conn, err := s.connect()
	c.Assert(err, IsNil)
	defer conn.Close()

	query := "select * from test"

	err = conn.RequestComQuery(query)
	c.Assert(err, IsNil)
}
