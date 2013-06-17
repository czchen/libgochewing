package libgochewing

import (
	"launchpad.net/gocheck"
	"testing"
)

func TestHook(t *testing.T) {
	gocheck.TestingT(t)
}

type MySuite struct{}

var _ = gocheck.Suite(&MySuite{})
