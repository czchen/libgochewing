package libgochewing

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Phrase struct {
	frequency uint32
	phrase    []rune
}

func newPhrase(str string, frequency uint32) (phrase *Phrase, err error) {
	strLen := utf8.RuneCountInString(str)

	phrase = new(Phrase)
	phrase.phrase = make([]rune, strLen)

	for i := 0; i < strLen; i++ {
		r, size := utf8.DecodeRuneInString(str)
		if r == utf8.RuneError {
			return nil, errors.New(fmt.Sprintf("`%s' contains invalid UTF8 character", str))
		}

		phrase.phrase[i] = r

		str = str[size:]
	}

	phrase.frequency = frequency
	return phrase, nil
}

func isTheSamePhrase(x *Phrase, y *Phrase) bool {
	if len(x.phrase) != len(y.phrase) {
		return false
	}

	for i := 0; i < len(x.phrase); i++ {
		if x.phrase[i] != y.phrase[i] {
			return false
		}
	}

	return true
}
