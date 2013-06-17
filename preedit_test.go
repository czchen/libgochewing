package libgochewing

type FakePhoneSeq struct {
	phoneSeq []uint16
}

func newFakePhoneSeq(phoneSeq []uint16) (fake *FakePhoneSeq) {
	fake = new(FakePhoneSeq)
	fake.phoneSeq = phoneSeq
	return fake
}

func (this *FakePhoneSeq) getLength() int {
	return len(this.phoneSeq)
}

func (this *FakePhoneSeq) getPhoneAtPos(pos int) uint16 {
	return this.phoneSeq[pos]
}
