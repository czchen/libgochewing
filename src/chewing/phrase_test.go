package chewing

import (
    "testing"
)

func TestGetNewPhrase(t *testing.T) {
    phrase, err := newPhrase("測試", "ㄘㄜˋ ㄕˋ", 10000)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    if phrase.frequency != 10000 {
        t.Errorf("frequency shall be %d, but it is %d", 10000, phrase.frequency)
    }

    var word Word

    word.word = 0x6e2c // 測 = u+6e2c
    word.phone = 10268
    if phrase.phrase[0] != word {
        t.Errorf("word in Phrase is not expected value. %s", phrase.phrase[0])
    }

    word.word = 0x8a66 // 試 = u+8a66
    word.phone = 8708
    if phrase.phrase[1] != word {
        t.Errorf("word in Phrase is not expected value. %s", phrase.phrase[0])
    }
}

func TestInsertAndQuery(t *testing.T) {
    var inputPhrase = [5]*Phrase{}
    var err error

    dict := newPhraseDictionary()

    inputPhrase[0], err = newPhrase("測試", "ㄘㄜˋ ㄕˋ", 5)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    inputPhrase[1], err = newPhrase("側室", "ㄘㄜˋ ㄕˋ", 4)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    inputPhrase[2], err = newPhrase("側視", "ㄘㄜˋ ㄕˋ", 3)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    inputPhrase[3], err = newPhrase("策士", "ㄘㄜˋ ㄕˋ", 2)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    inputPhrase[4], err = newPhrase("側視", "ㄘㄜˋ ㄕˋ", 1)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    dict.insertPhrase(inputPhrase[4])
    dict.insertPhrase(inputPhrase[3])
    dict.insertPhrase(inputPhrase[2])
    dict.insertPhrase(inputPhrase[1])
    dict.insertPhrase(inputPhrase[0])

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
