package libgochewing

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	BOPOMOFO_INITIAL = iota
	BOPOMOFO_MIDDLE
	BOPOMOFO_FINAL
	BOPOMOFO_TONE
)

const (
	PHONE_FUZZY_TONELESS = 1 << iota
	PHONE_FUZZY_ALL      = 0xffffffff
)

type BopomofoTable struct {
	name    string
	literal string
	shift   uint16
	mask    uint16
	length  uint16 // UTF-8 length
}

var BOPOMOFO_TABLE = [...]BopomofoTable{
	BOPOMOFO_INITIAL: {
		literal: "ㄅㄆㄇㄈㄉㄊㄋㄌㄍㄎㄏㄐㄑㄒㄓㄔㄕㄖㄗㄘㄙ",
		shift:   9,
		mask:    0x1f,
		length:  3,
	},
	BOPOMOFO_MIDDLE: {
		literal: "ㄧㄨㄩ",
		shift:   7,
		mask:    0x3,
		length:  3,
	},
	BOPOMOFO_FINAL: {
		literal: "ㄚㄛㄜㄝㄞㄟㄠㄡㄢㄣㄤㄥㄦ",
		shift:   3,
		mask:    0x1f,
		length:  3,
	},
	BOPOMOFO_TONE: {
		literal: "˙ˊˇˋ",
		shift:   0,
		mask:    0x7,
		length:  2,
	},
}

var BOPOMOFO_RE = regexp.MustCompile(
	"^" +
		"([" + BOPOMOFO_TABLE[BOPOMOFO_INITIAL].literal + "]?)" +
		"([" + BOPOMOFO_TABLE[BOPOMOFO_MIDDLE].literal + "]?)" +
		"([" + BOPOMOFO_TABLE[BOPOMOFO_FINAL].literal + "]?)" +
		"([" + BOPOMOFO_TABLE[BOPOMOFO_TONE].literal + "]?)" +
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

		index := strings.Index(item.literal, current)
		if index == -1 {
			panic(fmt.Sprintf("`%s' not in `%s'!", current, item.literal))
		}

		// index is byte index, not UTF-8 character index.
		phone += (uint16(index)/item.length + 1) << item.shift
	}

	return phone, nil
}

func convertPhoneToBopomofo(phone uint16) (bopomofo string, err error) {
	var buf bytes.Buffer

	for _, item := range BOPOMOFO_TABLE {
		index := (phone >> item.shift) & item.mask
		if index == 0 {
			continue
		}

		// index is byte index, not UTF-8 character index.
		index *= item.length

		if len(item.literal) < int(index) {
			return "", errors.New(fmt.Sprintf("%d is not a valid phone", phone))
		}

		buf.WriteString(item.literal[index-item.length : index])
	}

	return buf.String(), nil
}

func calculateHammingDistance(x PhoneSeq, y PhoneSeq) (distance int) {
	if x.getLength() != y.getLength() {
		panic(fmt.Sprintf("Cannot calculate hamming distance between %s and %s. Different length.", x, y))
	}

	for i := 0; i < x.getLength(); i++ {
		for _, item := range BOPOMOFO_TABLE {
			xx := (x.getPhoneAtPos(i) >> item.shift) & item.mask
			yy := (y.getPhoneAtPos(i) >> item.shift) & item.mask
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
		x &^= (BOPOMOFO_TABLE[BOPOMOFO_TONE].mask << BOPOMOFO_TABLE[BOPOMOFO_TONE].shift)
		y &^= (BOPOMOFO_TABLE[BOPOMOFO_TONE].mask << BOPOMOFO_TABLE[BOPOMOFO_TONE].shift)
	}

	return int(x) - int(y)
}

func getFuzzyPhone(phone uint16) uint16 {
	// PHONE_FUZZY_TONELESS
	phone &^= (BOPOMOFO_TABLE[BOPOMOFO_TONE].mask << BOPOMOFO_TABLE[BOPOMOFO_TONE].shift)

	return phone
}
