package goapns

type Alert struct {
	title string
	body  string
}

func NewAlert() *Alert {
	a := &Alert{"title", "body"}
	return a
}
