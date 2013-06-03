package chewing

import (
    "errors"
    "fmt"
)

type PhraseTreeNode struct {
    children map[uint16] *PhraseTreeNode
    allPhrase []*Phrase
}

type PhraseDictionary struct {
    root *PhraseTreeNode
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

func (this *PhraseTreeNode) insertPhrase(phrase *Phrase) (err error) {
    if this.allPhrase == nil {
        this.allPhrase = make([]*Phrase, 0, 1)
    }

    length := len(this.allPhrase)
    if length >= cap(this.allPhrase) {
        original := this.allPhrase
        this.allPhrase = make([]*Phrase, length, length + 1)
        copy(this.allPhrase, original)
    }

    pos := 0

    for i := 0; i < length; i++ {
        if isTheSamePhrase(this.allPhrase[i], phrase) {
            return errors.New(fmt.Sprintf("Phrase %s already in phrase tree", phrase.phrase))
        }

        if phrase.frequency < this.allPhrase[i].frequency {
            pos = i + 1
        }
    }

    this.allPhrase = this.allPhrase[:length + 1]
    copy(this.allPhrase[pos + 1: length + 1], this.allPhrase[pos: length])
    this.allPhrase[pos] = phrase

    return nil
}

func (this *PhraseDictionary) insertPhrase(phrase *Phrase, phoneSeq []uint16) (err error) {
    current := this.root
    for _, phone := range phoneSeq {
        if current.children[phone] == nil {
            current.children[phone] = newPhraseTreeNode()
        }
        current = current.children[phone]
    }
    return current.insertPhrase(phrase)
}

func (this *PhraseDictionary) queryPhrase(phoneSeq []uint16) (phrase []*Phrase){
    current := this.root

    for _, phone := range phoneSeq {
        if current.children[phone] == nil {
            return nil
        }
        current = current.children[phone]
    }
    return current.allPhrase
}
