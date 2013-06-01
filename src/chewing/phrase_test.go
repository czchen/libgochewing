package chewing

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

func TestInsertAndQuery(t *testing.T) {
    var inputPhrase = [5]*Phrase{}
    var err error

    inputPhrase[0], err = newPhrase("測試", 5)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    inputPhrase[1], err = newPhrase("側室", 4)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    inputPhrase[2], err = newPhrase("側視", 3)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    inputPhrase[3], err = newPhrase("策士", 2)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    inputPhrase[4], err = newPhrase("策試", 1)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    dict := newPhraseDictionary()
    phoneSeq := []uint16{ 10268, 8708 }

    dict.insertPhrase(inputPhrase[4], phoneSeq)
    dict.insertPhrase(inputPhrase[3], phoneSeq)
    dict.insertPhrase(inputPhrase[2], phoneSeq)
    dict.insertPhrase(inputPhrase[1], phoneSeq)
    dict.insertPhrase(inputPhrase[0], phoneSeq)

    queryPhrase := dict.queryPhrase([]uint16{ 10268, 8708 })
    if len(queryPhrase) != 5 {
        t.Errorf("len of queryPhrase shall be 5")
    }

    for i := 0; i < 5; i++ {
        if queryPhrase[i] != inputPhrase[i] {
            t.Errorf("Phrase index %d mismatch: %s != %s", i, queryPhrase[i], inputPhrase[i])
        }
    }
}
