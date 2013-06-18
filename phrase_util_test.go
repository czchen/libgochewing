package libgochewing

import (
	"launchpad.net/gocheck"
)

func (this *DefaultSuite) TestGetNewPhrase(c *gocheck.C) {
	phrase, err := newPhrase("測試", 10000)
	c.Check(err, gocheck.IsNil)
	c.Check(phrase.frequency, gocheck.Equals, uint32(10000))
	c.Check(phrase.phrase[0], gocheck.Equals, '測')
	c.Check(phrase.phrase[1], gocheck.Equals, '試')
}
