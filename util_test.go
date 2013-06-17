package libgochewing

import (
	"launchpad.net/gocheck"
)

type TestData struct {
	bopomofo string
	phone    uint16
}

var DATA = []TestData{
	{
		bopomofo: "ㄘㄜˋ",
		phone:    10268,
	},
	{
		bopomofo: "ㄕˋ",
		phone:    8708,
	},
}

func (this *DefaultSuite) TestConvertBopomofoAndPhone(c *gocheck.C) {
	for _, data := range DATA {
		phone, err := convertBopomofoToPhone(data.bopomofo)
		c.Check(phone, gocheck.Equals, data.phone)
		c.Check(err, gocheck.IsNil)

		bopomofo, err := convertPhoneToBopomofo(data.phone)
		c.Check(bopomofo, gocheck.Equals, data.bopomofo)
		c.Check(err, gocheck.IsNil)
	}
}

func (this *DefaultSuite) TestCalculateHammingDistance(c *gocheck.C) {
	base := newFakePhoneSeq([]uint16{10268, 8708})
	dist1 := newFakePhoneSeq([]uint16{10264, 8708})
	dist2 := newFakePhoneSeq([]uint16{8220, 10756})
	distPanic := newFakePhoneSeq([]uint16{10268, 8708, 10268, 8708})

	c.Check(calculateHammingDistance(base, base), gocheck.Equals, 0)
	c.Check(calculateHammingDistance(base, dist1), gocheck.Equals, 1)
	c.Check(calculateHammingDistance(base, dist2), gocheck.Equals, 2)
	c.Check(func() { calculateHammingDistance(base, distPanic) }, gocheck.PanicMatches, `.*`)
}

func (this *DefaultSuite) TestComparePhone(c *gocheck.C) {
	c.Check(comparePhoneSeq([]uint16{8708}, []uint16{10268}, 0) < 0, gocheck.Equals, true)
	c.Check(comparePhoneSeq([]uint16{10262, 10262}, []uint16{10262, 10262}, 0) == 0, gocheck.Equals, true)
	c.Check(comparePhoneSeq([]uint16{10264}, []uint16{10268}, PHONE_FUZZY_TONELESS) == 0, gocheck.Equals, true)
}
