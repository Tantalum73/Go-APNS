package goapns

type Payload struct {
	Badge            int
	Sound            string
	ContentAvailable int
	Category         string
}

func NewPayload() *Payload {
	p := &Payload{1, "Default", 0, ""}
	return p
}

func (p *Payload) MapInto(mapped map[string]interface{}) map[string]interface{} {
	//mapped := make(map[string]interface{}, 4)
	if p.Badge != 0 {
		mapped["badge"] = p.Badge
	}
	if p.Sound != "" {
		mapped["sound"] = p.Sound
	}
	if p.ContentAvailable != 0 {
		mapped["content-available"] = 1
	}
	if p.Category != "" {
		mapped["category"] = p.Category
	}
	return mapped
}
