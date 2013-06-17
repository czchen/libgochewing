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

func TestConvertBopomofoToPhone(t *testing.T) {
	for _, data := range DATA {
		phone, err := convertBopomofoToPhone(data.bopomofo)
		if err != nil {
			t.Errorf("convertBopomofoToPhone shall not return error %s", err.Error())
		}

		if phone != data.phone {
			t.Errorf("%s -> %d (shall be %d)", data.bopomofo, phone, data.phone)
		}
	}
}

func TestConvertPhoneToBopomofo(t *testing.T) {
	for _, data := range DATA {
		bopomofo, err := convertPhoneToBopomofo(data.phone)
		if err != nil {
			t.Errorf("convertPhoneToBopomofo shall not return error %s", err.Error())
		}

		if bopomofo != data.bopomofo {
			t.Errorf("%d -> %s (shall be %s)", data.phone, bopomofo, data.bopomofo)
		}
	}
}

func (this *MySuite) TestCalculateHammingDistance(c *gocheck.C) {
	base := newFakePhoneSeq([]uint16{10268, 8708})
	dist1 := newFakePhoneSeq([]uint16{10264, 8708})
	dist2 := newFakePhoneSeq([]uint16{8220, 10756})
	distPanic := newFakePhoneSeq([]uint16{10268, 8708, 10268, 8708})

	c.Check(calculateHammingDistance(base, base), gocheck.Equals, 0)
	c.Check(calculateHammingDistance(base, dist1), gocheck.Equals, 1)
	c.Check(calculateHammingDistance(base, dist2), gocheck.Equals, 2)
	c.Check(func() { calculateHammingDistance(base, distPanic) }, gocheck.PanicMatches, `.*`)
}

func TestComparePhone(t *testing.T) {
	var ret int
	var x []uint16
	var y []uint16

	x = []uint16{8708}
	y = []uint16{10268}
	ret = comparePhoneSeq(x, y, 0)
	if ret >= 0 {
		t.Errorf("comparePhoneSeq(%s, %s, 0) shall < 0, but got %d", x, y, ret)
	}

	x = []uint16{10268}
	y = []uint16{8708, 8708}
	ret = comparePhoneSeq(x, y, 0)
	if ret <= 0 {
		t.Errorf("comparePhoneSeq(%s, %s, 0) shall >= 0, but got %d", x, y, ret)
	}

	x = []uint16{10262, 10262}
	y = []uint16{10262, 10262}
	ret = comparePhoneSeq(x, y, 0)
	if ret != 0 {
		t.Errorf("comparePhoneSeq(%s, %s, 0) shall = 0, but got %d", x, y, ret)
	}

	x = []uint16{10264}
	y = []uint16{10268}
	ret = comparePhoneSeq(x, y, PHONE_FUZZY_TONELESS)
	if ret != 0 {
		t.Errorf("comparePhoneSeq(%s, %s, PHONE_FUZZY_TONELESS) shall = 0, but got %d", x, y, ret)
	}
}
