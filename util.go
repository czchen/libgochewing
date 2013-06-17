package libgochewing

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
)

const (
	BOPOMOFO_INITIAL = iota
	BOPOMOFO_MIDDLE
	BOPOMOFO_FINAL
	BOPOMOFO_TONE
)

const (
	PHONE_FUZZY_TONELESS = 1 << iota // ˙ˊˇˋ
	PHONE_FUZZY_EN_ENG               // ㄣㄥ
	PHONE_FUZZY_ALL      = (1 << iota) - 1
)

type BopomofoTable struct {
	bopomofoToPhone map[string]uint16
	phoneToBopomofo map[uint16]string
	mask            uint16
}

var BOPOMOFO_TABLE = [...]BopomofoTable{
	BOPOMOFO_INITIAL: {
		bopomofoToPhone: map[string]uint16{
			"ㄅ": 0x0200,
			"ㄆ": 0x0400,
			"ㄇ": 0x0600,
			"ㄈ": 0x0800,
			"ㄉ": 0x0a00,
			"ㄊ": 0x0c00,
			"ㄋ": 0x0e00,
			"ㄌ": 0x1000,
			"ㄍ": 0x1200,
			"ㄎ": 0x1400,
			"ㄏ": 0x1600,
			"ㄐ": 0x1800,
			"ㄑ": 0x1a00,
			"ㄒ": 0x1c00,
			"ㄓ": 0x1e00,
			"ㄔ": 0x2000,
			"ㄕ": 0x2200,
			"ㄖ": 0x2400,
			"ㄗ": 0x2600,
			"ㄘ": 0x2800,
			"ㄙ": 0x2a00,
		},
		phoneToBopomofo: map[uint16]string{
			0x0200: "ㄅ",
			0x0400: "ㄆ",
			0x0600: "ㄇ",
			0x0800: "ㄈ",
			0x0a00: "ㄉ",
			0x0c00: "ㄊ",
			0x0e00: "ㄋ",
			0x1000: "ㄌ",
			0x1200: "ㄍ",
			0x1400: "ㄎ",
			0x1600: "ㄏ",
			0x1800: "ㄐ",
			0x1a00: "ㄑ",
			0x1c00: "ㄒ",
			0x1e00: "ㄓ",
			0x2000: "ㄔ",
			0x2200: "ㄕ",
			0x2400: "ㄖ",
			0x2600: "ㄗ",
			0x2800: "ㄘ",
			0x2a00: "ㄙ",
		},
		mask: 0x3e00,
	},
	BOPOMOFO_MIDDLE: {
		bopomofoToPhone: map[string]uint16{
			"ㄧ": 0x080,
			"ㄨ": 0x100,
			"ㄩ": 0x180,
		},
		phoneToBopomofo: map[uint16]string{
			0x080: "ㄧ",
			0x100: "ㄨ",
			0x180: "ㄩ",
		},
		mask: 0x180,
	},
	BOPOMOFO_FINAL: {
		bopomofoToPhone: map[string]uint16{
			"ㄚ": 0x08,
			"ㄛ": 0x10,
			"ㄜ": 0x18,
			"ㄝ": 0x20,
			"ㄞ": 0x28,
			"ㄟ": 0x30,
			"ㄠ": 0x38,
			"ㄡ": 0x40,
			"ㄢ": 0x48,
			"ㄣ": 0x50,
			"ㄤ": 0x58,
			"ㄥ": 0x60,
			"ㄦ": 0x68,
		},
		phoneToBopomofo: map[uint16]string{
			0x08: "ㄚ",
			0x10: "ㄛ",
			0x18: "ㄜ",
			0x20: "ㄝ",
			0x28: "ㄞ",
			0x30: "ㄟ",
			0x38: "ㄠ",
			0x40: "ㄡ",
			0x48: "ㄢ",
			0x50: "ㄣ",
			0x58: "ㄤ",
			0x60: "ㄥ",
			0x68: "ㄦ",
		},
		mask: 0xf8,
	},
	BOPOMOFO_TONE: {
		bopomofoToPhone: map[string]uint16{
			"˙": 0x1,
			"ˊ": 0x2,
			"ˇ": 0x3,
			"ˋ": 0x4,
		},
		phoneToBopomofo: map[uint16]string{
			0x1: "˙",
			0x2: "ˊ",
			0x3: "ˇ",
			0x4: "ˋ",
		},
		mask: 0x07,
	},
}

var BOPOMOFO_RE = regexp.MustCompile(
	"^" +
		"([ㄅㄆㄇㄈㄉㄊㄋㄌㄍㄎㄏㄐㄑㄒㄓㄔㄕㄖㄗㄘㄙ]?)" +
		"([ㄧㄨㄩ]?)" +
		"([ㄚㄛㄜㄝㄞㄟㄠㄡㄢㄣㄤㄥㄦ]?)" +
		"([˙ˊˇˋ]?)" +
		"$")

func convertBopomofoToPhone(bopomofo string) (phone uint16, err error) {
	match := BOPOMOFO_RE.FindStringSubmatch(bopomofo)
	if match == nil {
		return 0, errors.New(fmt.Sprintf("`%s' is not a valid bopomofo", bopomofo))
	}

	phone = 0
	for index, item := range BOPOMOFO_TABLE {
		current := match[index+1]
		if current == "" {
			continue
		}

		phone |= item.bopomofoToPhone[current]
	}

	return phone, nil
}

func convertPhoneToBopomofo(phone uint16) (bopomofo string, err error) {
	var buf bytes.Buffer

	for _, item := range BOPOMOFO_TABLE {
		buf.WriteString(item.phoneToBopomofo[phone&item.mask])
	}

	return buf.String(), nil
}

func calculateHammingDistance(x PhoneSeq, y PhoneSeq) (distance int) {
	if x.getLength() != y.getLength() {
		panic(fmt.Sprintf("Cannot calculate hamming distance between %s and %s. Different length.", x, y))
	}

	for i := 0; i < x.getLength(); i++ {
		for _, item := range BOPOMOFO_TABLE {
			xx := x.getPhoneAtPos(i) & item.mask
			yy := y.getPhoneAtPos(i) & item.mask
			if xx != yy {
				distance++
			}
		}
	}

	return distance
}

func comparePhoneSeq(x []uint16, y []uint16, flag uint32) int {
	var min int
	lenX := len(x)
	lenY := len(y)

	if lenX > lenY {
		min = lenY
	} else {
		min = lenX
	}

	for i := 0; i < min; i++ {
		compare := comparePhone(x[i], y[i], flag)
		if compare != 0 {
			return compare
		}
	}

	return lenX - lenY
}

func comparePhone(x uint16, y uint16, flag uint32) int {
	if flag&PHONE_FUZZY_TONELESS == PHONE_FUZZY_TONELESS {
		x &^= BOPOMOFO_TABLE[BOPOMOFO_TONE].mask
		y &^= BOPOMOFO_TABLE[BOPOMOFO_TONE].mask
	}

	if flag&PHONE_FUZZY_EN_ENG == PHONE_FUZZY_EN_ENG {
	}

	return int(x) - int(y)
}

func getFuzzyPhone(phone uint16) uint16 {
	// PHONE_FUZZY_TONELESS
	phone &^= BOPOMOFO_TABLE[BOPOMOFO_TONE].mask

	return phone
}
