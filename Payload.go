package goapns

type Payload struct {
	Badge            int
	Sound            string
	ContentAvailable int
}

func NewPayload() *Payload {
	p := &Payload{1, "Sound", 0}
	return p
}
