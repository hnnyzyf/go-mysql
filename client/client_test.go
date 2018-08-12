package client

import (
	"crypto/tls"
	"testing"

	"github.com/hnnyzyf/go-mysql/util/ssl"
	"github.com/juju/errors"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MyClientSuite struct {
	tls     *tls.Config
	testTls bool
}

func (s *MyClientSuite) initTls() error {
	tls, err := ssl.NewConfig("./tls/ca.pem", "./tls/client-cert.pem", "./tls/client-key.pem")
	if err != nil {
		return errors.Trace(err)
	} else {
		s.tls = tls
	}

	return nil
}

var _ = Suite(&MyClientSuite{testTls: false})

func (s *MyClientSuite) TestConnect(c *C) {
	data := []struct {
		h string
		u string
		p string
		d string
	}{
		{"127.0.0.1:3306", "test", "123456", "test"},
		{"127.0.0.1:3306", "test", "123456", ""},
	}

	for i := range data {
		l := data[i]
		conn, err := Connect(l.h, l.u, l.p, l.d)
		c.Assert(err, IsNil)
		conn.Close()
	}
}

func (s *MyClientSuite) TestConnectWithSSL(c *C) {
	data := []struct {
		h string
		u string
		p string
		d string
	}{
		{"127.0.0.1:3306", "testssl", "123456", "test"},
		{"127.0.0.1:3306", "testssl", "123456", ""},
	}
	if s.testTls {
		cfg := NewConfig()
		if err := s.initTls(); err != nil {
			c.Error(err)
		}
		cfg.SetSSL(true, s.tls)
		for i := range data {
			l := data[i]
			conn, err := ConnectWithConfig(l.h, l.u, l.p, l.d, cfg)
			if err != nil {
				c.Error(err)
			}
			conn.Close()
		}
	}
}

func (s *MyClientSuite) TestConnectWithFailure(c *C) {
	data := []struct {
		h string
		u string
		p string
		d string
		e string
	}{
		{"127.0.0.1:3306", "test", "1234567", "test", `.+Acces.+`},
		{"127.0.0.1:3306", "test", "1234567", "", `.+Acces.+`},
		{"127.0.0.1:3306", "test2", "123456", "test", `.+Acces.+`},
		{"127.0.0.1:1234", "test", "123456", "", `.+connection.+`},
		{"127.0.0.1:3306", "test", "123456", "test2", `.+database.+`},
	}

	for i := range data {
		l := data[i]
		_, err := Connect(l.h, l.u, l.p, l.d)
		c.Assert(err, ErrorMatches, data[i].e)
	}
}
