package chewing

type Word struct {
    word string
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

func (this *PhraseDictionary) insert(phrase string, phone []uint16, frequency uint32) {
    currentNode := this.root
    for _, item := range phone {
        nextNode, ok := currentNode.children[item]
        if !ok {
            nextNode = newPhraseTreeNode()
            currentNode.children[item] = nextNode
        }
        currentNode = nextNode
    }
}
