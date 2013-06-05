package chewing

type PhraseTree struct {
    root *PhraseTreeNode
}

type PhraseTreeNode struct {
    children map[uint16] *PhraseTreeNode
    phraseArrayItem []*PhraseArrayItem
}

func newPhraseTree() (phraseTree *PhraseTree) {
    phraseTree = new(PhraseTree)
    phraseTree.root = newPhraseTreeNode()
    return phraseTree
}

func newPhraseTreeNode() (phraseTreeNode *PhraseTreeNode) {
    phraseTreeNode = new(PhraseTreeNode)
    phraseTreeNode.children = make(map[uint16] *PhraseTreeNode)
    return phraseTreeNode
}

func (this *PhraseTree) insert(phraseArrayItem *PhraseArrayItem) {
    current := this.root
    for _, phone := range phraseArrayItem.phoneSeq {
        phone = getFuzzyPhone(phone)
        if current.children[phone] == nil {
            current.children[phone] = newPhraseTreeNode()
        }
        current = current.children[phone]
    }
    current.insert(phraseArrayItem)
}

func (this *PhraseTreeNode) insert(phraseArrayItem *PhraseArrayItem) {
    if this.phraseArrayItem == nil {
        this.phraseArrayItem = make([]*PhraseArrayItem, 0, 1)
    }

    this.phraseArrayItem = append(this.phraseArrayItem, phraseArrayItem)
}

func (this *PhraseTree) query(phoneSeq []uint16, flag uint32) []*PhraseArrayItem {
    current := this.root
    for _, phone := range phoneSeq {
        phone = getFuzzyPhone(phone)
        if current.children[phone] == nil {
            return make([]*PhraseArrayItem, 0)
        }
        current = current.children[phone]
    }
    return current.query(phoneSeq, flag)
}

func (this *PhraseTreeNode) query(phoneSeq []uint16, flag uint32) (phraseArrayItem []*PhraseArrayItem) {
    phraseArrayItem = make([]*PhraseArrayItem, 0, len(this.phraseArrayItem))

    for _, item := range this.phraseArrayItem {
        if comparePhoneSeq(phoneSeq, item.phoneSeq, flag) == 0 {
            phraseArrayItem = append(phraseArrayItem, item)
        }
    }

    return phraseArrayItem
}
