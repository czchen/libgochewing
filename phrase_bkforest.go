package libgochewing

import (
	"errors"
)

type PhraseBKForest struct {
	tree map[int]*PhraseBKTreeNode
}

type PhraseBKTreeNode struct {
	children        map[int]*PhraseBKTreeNode
	phraseArrayItem *PhraseArrayItem
}

func newPhraseBKForest() (phraseBKForest *PhraseBKForest) {
	phraseBKForest = new(PhraseBKForest)
	phraseBKForest.tree = make(map[int]*PhraseBKTreeNode)
	return phraseBKForest
}

func newPhraseBKTreeNode(phraseArrayItem *PhraseArrayItem) (phraseBKTreeNode *PhraseBKTreeNode) {
	phraseBKTreeNode = new(PhraseBKTreeNode)
	phraseBKTreeNode.children = make(map[int]*PhraseBKTreeNode)
	phraseBKTreeNode.phraseArrayItem = phraseArrayItem
	return phraseBKTreeNode
}

func (this *PhraseBKForest) insert(phraseArrayItem *PhraseArrayItem) (err error) {
	length := len(phraseArrayItem.phoneSeq)
	if this.tree[length] == nil {
		this.tree[length] = newPhraseBKTreeNode(phraseArrayItem)
		return nil
	} else {
		return this.tree[length].insert(phraseArrayItem)
	}
}

func (this *PhraseBKTreeNode) insert(phraseArrayItem *PhraseArrayItem) (err error) {
	distance, err := calculateHammingDistance(this.phraseArrayItem, phraseArrayItem)
	if err != nil {
		return err
	}
	if distance == 0 {
		return errors.New("Duplicate phoneSeq insert")
	}

	if this.children[distance] == nil {
		this.children[distance] = newPhraseBKTreeNode(phraseArrayItem)
		return nil
	} else {
		return this.children[distance].insert(phraseArrayItem)
	}
}

func (this *PhraseBKForest) query(phoneSeq PhoneSeq, threshold int) (phraseArrayItem []*PhraseArrayItem) {
	length := phoneSeq.getLength()
	if this.tree[length] == nil {
		return make([]*PhraseArrayItem, 0)
	}

	phraseArrayItem = make([]*PhraseArrayItem, 0, 1)
	result := make(chan *PhraseArrayItem, 1000)
	count := make(chan int, 1000)

	counter := 1
	go this.tree[length].query(phoneSeq, threshold, count, result)

	for counter > 0 || len(result) > 0 {
		select {
		case res := <-result:
			phraseArrayItem = append(phraseArrayItem, res)
		case c := <-count:
			counter += c
		}
	}

	return phraseArrayItem
}

func (this *PhraseBKTreeNode) query(phoneSeq PhoneSeq, threshold int, count chan<- int, result chan<- *PhraseArrayItem) {
	diff, err := calculateHammingDistance(phoneSeq, this.phraseArrayItem)
	if err != nil {
		panic("calculateHammingDistance fails in PhraseBKTreeNode.query")
	}
	if diff <= threshold {
		result <- this.phraseArrayItem
	}

	for i := diff - threshold; i <= diff+threshold; i++ {
		if this.children[i] != nil {
			count <- 1
			go this.children[i].query(phoneSeq, threshold, count, result)
		}
	}

	count <- -1
}
