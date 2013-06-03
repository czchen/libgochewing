package chewing

import(
    "testing"
)

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

    duplicatedPhrase, err := newPhrase("測試", 5)
    if err != nil {
        t.Errorf("newPhrase shall not fail. It fails with %s", err.Error())
    }

    dict := newPhraseDictionary()
    phoneSeq := []uint16{ 10268, 8708 }

    err = dict.insertPhrase(inputPhrase[4], phoneSeq)
    if err != nil {
        t.Errorf("insertPhrase shall success")
    }

    err = dict.insertPhrase(inputPhrase[3], phoneSeq)
    if err != nil {
        t.Errorf("insertPhrase shall success")
    }

    err = dict.insertPhrase(inputPhrase[2], phoneSeq)
    if err != nil {
        t.Errorf("insertPhrase shall success")
    }

    err = dict.insertPhrase(inputPhrase[1], phoneSeq)
    if err != nil {
        t.Errorf("insertPhrase shall success")
    }

    err = dict.insertPhrase(inputPhrase[0], phoneSeq)
    if err != nil {
        t.Errorf("insertPhrase shall success")
    }

    err = dict.insertPhrase(duplicatedPhrase, phoneSeq)
    if err == nil {
        t.Errorf("insertPhrase shall reject duplicated phrase")
    }

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
