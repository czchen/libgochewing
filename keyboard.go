package libgochewing

const (
	KEYBOARD_MIN, KEYBOARD_DEFAULT, KEYBOARD_MAX = iota, iota, iota
)

type KeyboardTable struct {
	mapping map[byte]uint16
}

var KEYBOARD_TABLE = [...]KeyboardTable{
	KEYBOARD_DEFAULT: {
		mapping: map[byte]uint16{
			// BOPOMOFO_INITIAL
			'1': 0x0200,
			'q': 0x0400,
			'a': 0x0600,
			'z': 0x0800,
			'2': 0x0a00,
			'w': 0x0c00,
			's': 0x0e00,
			'x': 0x1000,
			'e': 0x1200,
			'd': 0x1400,
			'c': 0x1600,
			'r': 0x1800,
			'f': 0x1a00,
			'v': 0x1c00,
			'5': 0x1e00,
			't': 0x2000,
			'g': 0x2200,
			'b': 0x2400,
			'y': 0x2600,
			'h': 0x2800,
			'n': 0x2a00,
			// BOPOMOFO_MIDDLE
			'u': 0x080,
			'j': 0x100,
			'm': 0x180,
			// BOPOMOFO_FINAL
			'8': 0x08,
			'i': 0x10,
			'k': 0x18,
			',': 0x20,
			'9': 0x28,
			'o': 0x30,
			'l': 0x38,
			'.': 0x40,
			'0': 0x48,
			'p': 0x50,
			';': 0x58,
			'/': 0x60,
			'-': 0x68,
			// BOPOMOFO_TONE
			'7': 0x1,
			'6': 0x2,
			'3': 0x3,
			'4': 0x4,
		},
	},
}

func getPhoneFromKey(key byte, keyboard int) uint16 {
	return KEYBOARD_TABLE[keyboard].mapping[key]
}
