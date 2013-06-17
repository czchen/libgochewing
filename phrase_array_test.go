package libgochewing

import (
	"testing"
)

func TestPhraseArrayInsert(t *testing.T) {
	var inputPhrase = [5]*Phrase{}
	var err error

	inputPhrase[0], err = newPhrase("測試", 5)
	if inputPhrase[0] == nil {
		t.Fatal("newPhrase shall success")
	}
	if err != nil {
		t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
	}

	inputPhrase[1], err = newPhrase("側室", 4)
	if inputPhrase[1] == nil {
		t.Fatal("newPhrase shall success")
	}
	if err != nil {
		t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
	}

	inputPhrase[2], err = newPhrase("側視", 3)
	if inputPhrase[2] == nil {
		t.Fatal("newPhrase shall success")
	}
	if err != nil {
		t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
	}

	inputPhrase[3], err = newPhrase("策士", 2)
	if inputPhrase[3] == nil {
		t.Fatal("newPhrase shall success")
	}
	if err != nil {
		t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
	}

	inputPhrase[4], err = newPhrase("策試", 1)
	if inputPhrase[4] == nil {
		t.Fatal("newPhrase shall success")
	}
	if err != nil {
		t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
	}

	duplicatedPhrase, err := newPhrase("測試", 5)
	if duplicatedPhrase == nil {
		t.Fatal("newPhrase shall success")
	}
	if err != nil {
		t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
	}

	phraseArray := newPhraseArray()
	if phraseArray == nil {
		t.Fatal("newPhraseArray shall not fail")
	}
	phoneSeq := []uint16{10268, 8708}

	err = phraseArray.insert(inputPhrase[4], phoneSeq)
	if err != nil {
		t.Errorf("insert shall success")
	}

	err = phraseArray.insert(inputPhrase[3], phoneSeq)
	if err != nil {
		t.Errorf("insert shall success")
	}

	err = phraseArray.insert(inputPhrase[2], phoneSeq)
	if err != nil {
		t.Errorf("insert shall success")
	}

	err = phraseArray.insert(inputPhrase[1], phoneSeq)
	if err != nil {
		t.Errorf("insert shall success")
	}

	err = phraseArray.insert(inputPhrase[0], phoneSeq)
	if err != nil {
		t.Errorf("insert shall success")
	}

	err = phraseArray.insert(duplicatedPhrase, phoneSeq)
	if err == nil {
		t.Errorf("insert shall reject duplicated phrase")
	}

	if comparePhoneSeq(phraseArray.array[0].phoneSeq, phoneSeq, 0) != 0 {
		t.Errorf("phoneSeq %s is not expected value %s", phraseArray.array[0].phoneSeq, phoneSeq)
	}

	for i := 0; i < 5; i++ {
		if phraseArray.array[0].phrase[i] != inputPhrase[i] {
			t.Errorf("Phrase index %d mismatch: %s != %s", i, phraseArray.array[0].phrase[i], inputPhrase[i])
		}
	}
}
