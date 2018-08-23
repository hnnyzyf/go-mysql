package client

import (
	"crypto/tls"
	. "gopkg.in/check.v1"
	"time"
)

type MyDriverSuite struct {
}

var _ = Suite(&MyDriverSuite{})

//username:password@protocol(address)/dbname?param=value
func (s *MyDriverSuite) TestParseDSNAndFormatDSN(c *C) {
	data := []struct {
		dsn string
		cfg *Cfg
	}{
		{"test:123456@TCP(127.0.0.1:3306)/test", &Cfg{
			Username: "test",
			Password: "123456",
			Address:  "127.0.0.1:3306",
			DBname:   "test",
			Uid:      "",
			Config:   defaultConfig,
		}},
		{"test:123456@TCP(127.0.0.1:3306)/test?timeout=1m0s?readtime=2ms?writetime=3h0m0s?charset=utf8mb4?ssl=true?autoCommit=false",
			&Cfg{
				Username: "test",
				Password: "123456",
				Address:  "127.0.0.1:3306",
				DBname:   "test",
				Uid:      "",
				Config: &Config{
					Timeout:             1 * time.Minute,
					ReadTime:            2 * time.Millisecond,
					WriteTime:           3 * time.Hour,
					Capabilities:        defaultConfig.Capabilities,
					Charset:             "utf8mb4",
					MaxAllowedPacket:    defaultConfig.MaxAllowedPacket,
					AllowSSL:            true,
					TlsConfig:           defaultConfig.TlsConfig,
					AllowMutilStatement: defaultConfig.AllowMutilStatement,
					AllowAutoCommit:     false,
					Attrs:               defaultConnectionAttrs,
				},
			},
		},
	}

	for i := range data {
		cfg, err := ParseDSN(data[i].dsn)
		c.Assert(err, IsNil)
		c.Assert(cfg, DeepEquals, data[i].cfg)

		cfg = data[i].cfg
		str, err := cfg.FormatDSN()
		c.Assert(err, IsNil)
		c.Assert(str, Equals, data[i].dsn)
	}
}

func (s *MyConfigSuite) TestParseDsnError(c *C) {
	data := []struct {
		dsn string
		err string
	}{
		{"test", "EOF"},
		{"test:123456dsdTCP(127.0.0.1:3306)/test", "EOF"},
		{"test:123456@TCP(127.0.0.1:3306)/test?", `.*endwith.*`},
		{"test:123456@TCP(127.0.0.1:3306)/test?dasdasdadasd", `.*option.*`},
		{"test:123456@TCP(127.0.0.1:3306)/test?a=", `.*accpted.*`},
		{"test:123456@TCP(127.0.0.1:3306)/test?a=1", `.*supported.*`},
		{"test:123456@TCP(127.0.0.1:3306)/test?timeout=10", `.*missing.*`},
		{"test:123456@TCP(127.0.0.1:3306)/test?ssl=10", `.*SSL.*`},
		{"test:123456@TCP(127.0.0.1:3306)/test?tlsconfig=10", `.*registered.*`},
		{"test:123456@TCP(127.0.0.1:3306)/test?timeout=1m?timeout=2m", `.*twice.*`},
	}

	for i := range data {
		_, err := ParseDSN(data[i].dsn)
		c.Assert(err, ErrorMatches, data[i].err)
	}
}

func (s *MyConfigSuite) TestRegistSSL(c *C) {
	cfg := &tls.Config{}
	uid := RegisterTlsConfig(cfg)
	tmp, err := FindTlsConfig(uid)
	c.Assert(err, IsNil)
	c.Assert(tmp, Equals, cfg)
}
