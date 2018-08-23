package client

import (
	"github.com/hnnyzyf/go-mysql/protocol"
	. "gopkg.in/check.v1"
)

type MyConfigSuite struct{}

var _ = Suite(&MyConfigSuite{})

func (s *MyConfigSuite) TestAttr(c *C) {
	attr := Attr{
		"_os":             "1",
		"_client_name":    "2",
		"_pid":            "3",
		"_client_version": "4",
		"_platform":       "5",
		"program_name":    "6",
	}

	size := len("_os"+"_client_name"+"_pid"+"_client_version"+"_platform"+"program_name") + 6 + 6 + 6
	c.Assert(attr.size(), Equals, size)

	_, err := attr.marshal()
	c.Assert(err, IsNil)

}
