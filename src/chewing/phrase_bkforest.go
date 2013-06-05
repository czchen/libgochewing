package chewing

import (
    "errors"
)

type PhraseBKForest struct {
    tree map[int]*PhraseBKTreeNode
}

type PhraseBKTreeNode struct {
    children map[uint8]*PhraseBKTreeNode
    phraseArrayItem *PhraseArrayItem
}

func NewPhraseBKForest() (phraseBKForest *PhraseBKForest) {
    phraseBKForest = new(PhraseBKForest)
    phraseBKForest.tree = make(map[int]*PhraseBKTreeNode)
    return phraseBKForest
}

func NewPhraseBKTreeNode(phraseArrayItem *PhraseArrayItem) (phraseBKTreeNode *PhraseBKTreeNode) {
    phraseBKTreeNode = new(PhraseBKTreeNode)
    phraseBKTreeNode.children = make(map[uint8]*PhraseBKTreeNode)
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
