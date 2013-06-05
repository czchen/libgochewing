package chewing

import (
    "bytes"
    "errors"
    "fmt"
    "regexp"
    "strings"
)

const (
    BOPOMOFO_INITIAL = 0
    BOPOMOFO_MIDDLE  = 1
    BOPOMOFO_FINAL   = 2
    BOPOMOFO_TONE    = 3
)

const (
    PHONE_FUZZY_TONELESS = 0x00000001
    PHONE_FUZZY_ALL      = 0xffffffff
)

type BopomofoTable struct {
    name string
    literal string
    shift uint16
    mask uint16
    length uint16 // UTF-8 length
}

var BOPOMOFO_TABLE = [...]BopomofoTable{
    {
        literal: "ㄅㄆㄇㄈㄉㄊㄋㄌㄍㄎㄏㄐㄑㄒㄓㄔㄕㄖㄗㄘㄙ",
        shift: 9,
        mask: 0x1f,
        length: 3,
    },
    {
        literal: "ㄧㄨㄩ",
        shift: 7,
        mask: 0x3,
        length: 3,
    },
    {
        literal: "ㄚㄛㄜㄝㄞㄟㄠㄡㄢㄣㄤㄥㄦ",
        shift: 3,
        mask: 0x1f,
        length: 3,
    },
    {
        literal: "˙ˊˇˋ",
        shift: 0,
        mask: 0x7,
        length: 2,
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
        current := match[index + 1]

        if current == "" {
            continue
        }

        index := strings.Index(item.literal, current)
        if index == -1 {
            panic(fmt.Sprintf("`%s' not in `%s'!", current, item.literal))
        }

        // index is byte index, not UTF-8 character index.
        phone += (uint16(index) / item.length + 1) << item.shift
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

        buf.WriteString(item.literal[index - item.length: index])
    }

    return buf.String(), nil
}

func calculateHammingDistance(x []uint16, y[]uint16) (distance int, err error) {
    if len(x) != len(y) {
        return distance, errors.New(fmt.Sprintf("Cannot calculate hamming distance between %s and %s. Different length.", x, y))
    }

    for i := 0; i < len(x); i++ {
        for _, item := range BOPOMOFO_TABLE {
            xx := (x[i] >> item.shift) & item.mask
            yy := (y[i] >> item.shift) & item.mask
            if xx != yy {
                distance++
            }
        }
    }

    return distance, nil
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
    if flag & PHONE_FUZZY_TONELESS == PHONE_FUZZY_TONELESS {
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
