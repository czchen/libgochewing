package chewing

import (
    "testing"
)

type TestData struct {
    bopomofo string
    phone uint16
}

var DATA = []TestData{
    {
        bopomofo: "ㄘㄜˋ",
        phone: 10268,
    },
    {
        bopomofo: "ㄕˋ",
        phone: 8708,
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
