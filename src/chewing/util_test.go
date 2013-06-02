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

func TestCalculateHammingDistance(t *testing.T) {
    var distance uint8
    var err error

    base := []uint16{ 10268, 8708 }
    dist1 := []uint16{ 10264, 8708 }
    dist2 := []uint16{ 8220, 10756 }
    distErr := []uint16{ 10268, 8708, 10268, 8708 }

    distance, err = calculateHammingDistance(base, base)
    if distance != 0 {
        t.Errorf("Hamming distance between %s and %s shall be 0. Got %d", base, base, distance)
    }
    if err != nil {
        t.Errorf("calculateHammingDistance shall not return error %s", err.Error())
    }

    distance, err = calculateHammingDistance(base, dist1)
    if distance != 1 {
        t.Errorf("Hamming distance between %s and %s shall be 1. Got %d", base, dist1, distance)
    }
    if err != nil {
        t.Errorf("calculateHammingDistance shall not return error %s", err.Error())
    }

    distance, err = calculateHammingDistance(base, dist2)
    if distance != 2 {
        t.Errorf("Hamming distance between %s and %s shall be 2. Got %d", base, dist2, distance)
    }
    if err != nil {
        t.Errorf("calculateHammingDistance shall not return error %s", err.Error())
    }

    distance, err = calculateHammingDistance(base, distErr)
    if err == nil {
        t.Error("calculateHammingDistance shall return error")
    }
}
