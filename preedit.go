package libgochewing

type PhoneSeq interface {
	getLength() int
	getPhoneAtPos(pos int) uint16
}

type PreeditBufItem struct {
	char  string
	phone uint16
}

type PreeditBuf struct {
	buffer []PreeditBufItem
}

func NewPreeditBuf() (preeditBuf *PreeditBuf) {
	preeditBuf = new(PreeditBuf)
	preeditBuf.buffer = make([]PreeditBufItem, 0)
	return preeditBuf
}

func (this *PreeditBuf) getLength() int {
	return len(this.buffer)
}

func (this *PreeditBuf) getPhoneAtPos(pos int) uint16 {
	return this.buffer[pos].phone
}
