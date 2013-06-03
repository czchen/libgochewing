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

    length := len(this.phraseArrayItem)
    if length == cap(this.phraseArrayItem) {
        original := this.phraseArrayItem
        this.phraseArrayItem = make([]*PhraseArrayItem, length, length + 1)
        copy(this.phraseArrayItem, original)
    }

    this.phraseArrayItem = this.phraseArrayItem[:length + 1]
    this.phraseArrayItem[length] = phraseArrayItem
}
