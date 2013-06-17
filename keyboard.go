package libgochewing

import (
	"strings"
)

const (
	KEYBOARD_MIN, KEYBOARD_DEFAULT, KEYBOARD_MAX = iota, iota, iota
)

type KeyboardTable struct {
	key []string
}

var KEYBOARD_TABLE = [...]KeyboardTable{
	KEYBOARD_DEFAULT: {
		key: []string{
			BOPOMOFO_INITIAL: "1qaz2wsxedcrfv5tgbyhn",
			BOPOMOFO_MIDDLE:  "ujm",
			BOPOMOFO_FINAL:   "8ik,9ol.0p;/-",
			BOPOMOFO_TONE:    "7634",
		},
	},
}

func getPhoneFromKey(key byte, keyboard int) (phone uint16) {
	phone = 0
	for category, item := range KEYBOARD_TABLE[keyboard].key {
		index := strings.IndexRune(item, rune(key))
		if index != -1 {
			phone = uint16((index + 1) << BOPOMOFO_TABLE[category].shift)
			break
		}
	}
	return phone
}
