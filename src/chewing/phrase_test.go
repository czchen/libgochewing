package chewing

import (
    "testing"
)

func TestPhraseInsert(t *testing.T) {
    dict := newPhraseDictionary()
    dict.insert("測試", []uint16{ 10268, 8708 }, 0)
}
