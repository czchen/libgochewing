package chewing

import (
    "errors"
    "fmt"
    "strings"
    "unicode/utf8"
)

type Word struct {
    word rune
    phone uint16
}

type Phrase struct {
    frequency uint32
    phrase []Word
}

type PhraseTreeNode struct {
    children map[uint16] *PhraseTreeNode
    allPhrase []Phrase
}

type PhraseDictionary struct {
    root *PhraseTreeNode
}

func newPhrase(str string, bopomofo string, frequency uint32) (phrase *Phrase, err error) {
    strLen := utf8.RuneCountInString(str)
    bopomofoArray := strings.Split(bopomofo, " ")
    bopomofoLen := len(bopomofoArray)

    if strLen != bopomofoLen {
        return nil, errors.New(fmt.Sprintf("len(%s) = %d, len(%s) = %d", str, strLen, bopomofo, bopomofoLen))
    }

    phrase = new(Phrase)
    phrase.phrase = make([]Word, strLen)

    for i := 0; i < strLen; i++ {
        r, size := utf8.DecodeRuneInString(str)
        if r == utf8.RuneError {
            return nil, errors.New(fmt.Sprintf("`%s' contains invalid UTF8 character", str))
        }

        phone, err := convertBopomofoToPhone(bopomofoArray[i])
        if err != nil {
            return nil, err
        }

        phrase.phrase[i].word = r
        phrase.phrase[i].phone = phone

        str = str[size:]
    }

    phrase.frequency = frequency
    return phrase, nil
}

func newPhraseTreeNode() (node *PhraseTreeNode) {
    node = new(PhraseTreeNode)
    node.children = make(map[uint16] *PhraseTreeNode)

    return node
}

func newPhraseDictionary() (dict *PhraseDictionary) {
    dict = new(PhraseDictionary)
    dict.root = newPhraseTreeNode()

    return dict
}
