package pgt

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&TimeSuite{})
	Suite(&ArraySuite{})
	Suite(&UUIDSuite{})
	Suite(&StringSuite{})
}
