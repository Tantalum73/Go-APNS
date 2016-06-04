package goapns

import "fmt"

type Payload struct {
	Badge            int
	Sound            string
	ContentAvailable int
	Category         string
}

func NewPayload() Payload {
	p := Payload{-1, "", 0, ""}
	fmt.Printf("INITING, Badge: %v\n", p.Badge)
	return p
}

func (p *Payload) MapInto(mapped map[string]interface{}) map[string]interface{} {
	fmt.Printf("JSONING, Badge: %v\n", p.Badge)
	if p.Badge >= 0 {
		//Only set badge if the user specified so (by setting a >= 0 value).
		//If not, the badge is unchanged in the app.
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
