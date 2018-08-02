package client

import (
	"testing"

	. "gopkg.in/check.v1"
)

type MyClientSuite struct{}

var _ = Suite(&MyClientSuite{})

func (s *MyClientSuite) TestClient(c *C) {}
