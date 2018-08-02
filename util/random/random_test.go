package random

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) BenchmarkRandomString(c *C) {
	for i := 0; i < c.N; i++ {
		String(20)
	}
}
