package goapns

type Alert struct {
	Title        string   `json:"title,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`

	Body    string   `json:"body,omitempty"`
	LocKey  string   `json:"loc-key,omitempty"`
	LocArgs []string `json:"loc-args,omitempty"`

	ActionLocKey string `json:"action-loc-key,omitempty"`

	LaunchImage string `json:"launch-image,omitempty"`
}

func NewAlert() *Alert {
	a := &Alert{}
	return a
}
