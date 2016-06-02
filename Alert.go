package goapns

type Alert struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

func NewAlert() *Alert {
	a := &Alert{"title", "body"}
	return a
}
