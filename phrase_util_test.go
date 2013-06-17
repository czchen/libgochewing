package libgochewing

import (
	"testing"
)

func TestGetNewPhrase(t *testing.T) {
	phrase, err := newPhrase("測試", 10000)
	if err != nil {
		t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
	}

	if phrase.frequency != 10000 {
		t.Errorf("frequency shall be %d, but it is %d", 10000, phrase.frequency)
	}

	var word Word

	word.word = 0x6e2c // 測 = u+6e2c
	if phrase.phrase[0] != word {
		t.Errorf("word in Phrase is not expected value. %s", phrase.phrase[0])
	}

	word.word = 0x8a66 // 試 = u+8a66
	if phrase.phrase[1] != word {
		t.Errorf("word in Phrase is not expected value. %s", phrase.phrase[0])
	}
}
