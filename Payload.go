package goapns

type Payload struct {
	Badge int
	Sound string
}

func NewPayload() *Payload {
	p := &Payload{1, "Sound"}
	return p
}
