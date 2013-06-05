package chewing

import (
    "errors"
)

type PhraseBKForest struct {
    tree map[int]*PhraseBKTreeNode
}

type PhraseBKTreeNode struct {
    children map[int]*PhraseBKTreeNode
    phraseArrayItem *PhraseArrayItem
}

func NewPhraseBKForest() (phraseBKForest *PhraseBKForest) {
    phraseBKForest = new(PhraseBKForest)
    phraseBKForest.tree = make(map[int]*PhraseBKTreeNode)
    return phraseBKForest
}

func NewPhraseBKTreeNode(phraseArrayItem *PhraseArrayItem) (phraseBKTreeNode *PhraseBKTreeNode) {
    phraseBKTreeNode = new(PhraseBKTreeNode)
    phraseBKTreeNode.children = make(map[int]*PhraseBKTreeNode)
    phraseBKTreeNode.phraseArrayItem = phraseArrayItem
    return phraseBKTreeNode
}

func (this *PhraseBKForest) insert(phraseArrayItem *PhraseArrayItem) (err error) {
    length := len(phraseArrayItem.phoneSeq)
    if this.tree[length] == nil {
        this.tree[length] = NewPhraseBKTreeNode(phraseArrayItem)
        return nil
    } else {
        return this.tree[length].insert(phraseArrayItem)
    }
}

func (this *PhraseBKTreeNode) insert(phraseArrayItem *PhraseArrayItem) (err error) {
    distance, err := calculateHammingDistance(this.phraseArrayItem.phoneSeq, phraseArrayItem.phoneSeq)
    if err != nil {
        return err
    }
    if distance == 0 {
        return errors.New("Duplicate phoneSeq insert")
    }

    if this.children[distance] == nil {
        this.children[distance] = NewPhraseBKTreeNode(phraseArrayItem)
        return nil
    } else {
        return this.children[distance].insert(phraseArrayItem)
    }
}

func (this *PhraseBKForest) query(phoneSeq []uint16, threshold int) (phraseArrayItem []*PhraseArrayItem) {
    length := len(phoneSeq)
    if this.tree[length] == nil {
        return make([]*PhraseArrayItem, 0)
    }

    phraseArrayItem = make([]*PhraseArrayItem, 0, 2)
    result := make(chan *PhraseArrayItem)
    count := make(chan int)

    counter := 1
    go this.tree[length].query(phoneSeq, threshold, count, result)

    for counter > 0 {
        select {
        case res := <- result:
            if len(phraseArrayItem) == cap(phraseArrayItem) {
                origin := phraseArrayItem
                phraseArrayItem = make([]*PhraseArrayItem, 0, len(origin))
                copy(phraseArrayItem, origin)
            }
            phraseArrayItem = phraseArrayItem[:len(phraseArrayItem) + 1]
            phraseArrayItem[len(phraseArrayItem) - 1] = res
        case c := <- count:
            counter += c
        }
    }

    return phraseArrayItem
}

func (this *PhraseBKTreeNode) query(phoneSeq []uint16, threshold int, count chan<- int, result chan<- *PhraseArrayItem) {
    diff, err := calculateHammingDistance(phoneSeq, this.phraseArrayItem.phoneSeq)
    if err != nil {
        panic("calculateHammingDistance fails in PhraseBKTreeNode.query")
    }
    if diff <= threshold {
        result <- this.phraseArrayItem
    }

    for i := diff - threshold; i <= diff + threshold; i++ {
        if this.children[i] != nil {
            count <- 1
            go this.children[i].query(phoneSeq, threshold, count, result)
        }
    }

    count <- -1
}
