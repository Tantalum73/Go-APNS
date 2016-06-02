package goAPNS

type Header struct {
	APNSID int
}

func NewHeader() *Header {
	h := &Header{1}
	return h
}
