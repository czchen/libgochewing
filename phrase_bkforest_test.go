package libgochewing

import (
	"testing"
)

func TestBKForestInsertAndQuery(t *testing.T) {
	item1 := PhraseArrayItem{phoneSeq: []uint16{10268, 8708}}
	item2 := PhraseArrayItem{phoneSeq: []uint16{10264, 8708}}
	var ret []*PhraseArrayItem

	forest := newPhraseBKForest()

	forest.insert(&item1)
	forest.insert(&item2)

	ret = forest.query(newFakePhoneSeq([]uint16{10268, 8708}), 0)
	if len(ret) != 1 {
		t.Errorf("query shall return 1 PhraseArrayItem. Got %d", len(ret))
	}

	ret = forest.query(newFakePhoneSeq([]uint16{10268, 8708}), 1)
	if len(ret) != 2 {
		t.Errorf("query shall return 2 PhraseArrayItem. Got %d", len(ret))
	}
}
