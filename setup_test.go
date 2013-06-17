package libgochewing

import (
	"launchpad.net/gocheck"
	"testing"
)

func TestHook(t *testing.T) {
	gocheck.TestingT(t)
}

type EmptySuite struct{}

var _ = gocheck.Suite(&EmptySuite{})
