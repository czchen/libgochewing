package chewing

import(
    "errors"
    "fmt"
)

type PhraseArray struct {
    array []*PhraseArrayItem
}

type PhraseArrayItem struct {
    phoneSeq []uint16
    phrase []*Phrase
}

func newPhraseArray() (phraseArray *PhraseArray) {
    phraseArray = new(PhraseArray)
    phraseArray.array = make([]*PhraseArrayItem, 0, 1024)
    return phraseArray
}

func newPhraseArrayItem(phoneSeq []uint16) (phraseArrayItem *PhraseArrayItem) {
    phraseArrayItem = new(PhraseArrayItem)
    phraseArrayItem.phoneSeq = phoneSeq
    return phraseArrayItem
}

func (this *PhraseArray) insert(phrase *Phrase, phoneSeq []uint16) (err error) {
    begin := 0
    end := len(this.array)

    for begin < end {
        pos := (begin + end) / 2
        compare := comparePhoneSeq(this.array[pos].phoneSeq, phoneSeq, 0)

        if compare == 0 {
            return this.array[pos].insert(phrase)
        } else if compare > 0 {
            begin = pos + 1
        } else {
            end = pos
        }
    }

    this.array = append(this.array, nil)
    copy(this.array[begin + 1:], this.array[begin:])

    newPhraseArrayItem := newPhraseArrayItem(phoneSeq)
    newPhraseArrayItem.insert(phrase)
    this.array[begin] = newPhraseArrayItem

    return nil
}

func (this *PhraseArrayItem) insert(phrase *Phrase) (err error) {
    pos := 0
    for i, item := range this.phrase {
        if isTheSamePhrase(item, phrase) {
            return errors.New(fmt.Sprintf("Phrase %s already in phrase tree", phrase.phrase))
        }

        if phrase.frequency < item.frequency {
            pos = i + 1
        }
    }

    this.phrase = append(this.phrase, nil)
    copy(this.phrase[pos + 1:], this.phrase[pos:])
    this.phrase[pos] = phrase

    return nil
}
