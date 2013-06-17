package libgochewing

import (
	"testing"
)

func TestPhraseTreeInsertAndQuery(t *testing.T) {
	item1 := PhraseArrayItem{phoneSeq: []uint16{10268, 8708}}
	item2 := PhraseArrayItem{phoneSeq: []uint16{10264, 8708}}
	var ret []*PhraseArrayItem

	tree := newPhraseTree()

	tree.insert(&item1)
	tree.insert(&item2)

	ret = tree.query([]uint16{10268, 8708}, 0)
	if len(ret) != 1 {
		t.Errorf("query shall return 1 PhraseArrayItem. Got %d", len(ret))
	}

	ret = tree.query([]uint16{10265, 8708}, 0)
	if len(ret) != 0 {
		t.Errorf("query shall return 0 PhraseArrayItem. Got %d", len(ret))
	}

	ret = tree.query([]uint16{10268, 8708}, PHONE_FUZZY_TONELESS)
	if len(ret) != 2 {
		t.Errorf("query shall return 2 PhraseArrayItem. Got %d", len(ret))
	}

	ret = tree.query([]uint16{10265, 8708}, PHONE_FUZZY_TONELESS)
	if len(ret) != 2 {
		t.Errorf("query shall return 2 PhraseArrayItem. Got %d", len(ret))
	}
}
