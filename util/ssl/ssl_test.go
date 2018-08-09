package ssl

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySSLSuite struct {
	cakey  string
	capem  string
	clcert string
	clkey  string
	prikey string
	pubkey string
	secert string
	sekey  string
}

var _ = Suite(&MySSLSuite{
	cakey:  "./pem/ca-key.pem",
	capem:  "./pem/ca.pem",
	clcert: "./pem/client-cert.pem",
	clkey:  "./pem/client-key.pem",
	prikey: "./pem/private_key.pem",
	pubkey: "./pem/public_key.pem",
	secert: "./pem/server-cert.pem",
	sekey:  "./pem/server-key.pem",
})

func (s *MySSLSuite) TestConfig(c *C) {
	_, err := NewConfig(s.capem, s.clcert, s.clkey)
	if err != nil {
		c.Error(err)
	}
}
