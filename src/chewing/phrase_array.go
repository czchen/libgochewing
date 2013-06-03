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

    newPhraseArrayItem := newPhraseArrayItem(phoneSeq)
    newPhraseArrayItem.insert(phrase)

    length := len(this.array)
    if length == cap(this.array) {
        original := this.array
        this.array = make([]*PhraseArrayItem, length + 1, length * 2)
        copy(this.array[:begin], original[:begin])
        copy(this.array[begin + 1:length + 1], this.array[begin:length])
        this.array[begin] = newPhraseArrayItem
    } else {
        this.array = this.array[:length + 1]
        copy(this.array[begin + 1:length + 1], this.array[begin:length])
        this.array[begin] = newPhraseArrayItem
    }

    return nil
}

func (this *PhraseArrayItem) insert(phrase *Phrase) (err error) {
    length := len(this.phrase)
    if length == cap(this.phrase) {
        original := this.phrase
        this.phrase = make([]*Phrase, length, length + 1)
        copy(this.phrase, original)
    }

    pos := 0
    for i := 0; i < length; i++ {
        if isTheSamePhrase(this.phrase[i], phrase) {
            return errors.New(fmt.Sprintf("Phrase %s already in phrase tree", phrase.phrase))
        }

        if phrase.frequency < this.phrase[i].frequency {
            pos = i + 1
        }
    }

    this.phrase = this.phrase[:length + 1]
    copy(this.phrase[pos + 1: length + 1], this.phrase[pos: length])
    this.phrase[pos] = phrase

    return nil
}
