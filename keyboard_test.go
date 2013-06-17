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

func (this *MySuite) TestGetPhoneFromKey(c *gocheck.C) {
	c.Check(getPhoneFromKey('1', KEYBOARD_DEFAULT), gocheck.Equals, uint16(512))
	c.Check(getPhoneFromKey('u', KEYBOARD_DEFAULT), gocheck.Equals, uint16(128))
	c.Check(getPhoneFromKey('8', KEYBOARD_DEFAULT), gocheck.Equals, uint16(8))
	c.Check(getPhoneFromKey('7', KEYBOARD_DEFAULT), gocheck.Equals, uint16(1))
	c.Check(getPhoneFromKey('+', KEYBOARD_DEFAULT), gocheck.Equals, uint16(0))
	c.Check(func() { getPhoneFromKey('1', -1) }, gocheck.PanicMatches, `.*`)
}
