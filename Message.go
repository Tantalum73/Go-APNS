package goapns

type Message struct {
	Payload Payload
	Alert   Alert
	Header  Header
}

func NewMessage() *Message {
	return &Message{}
}
func (m *Message) Badge(number int) *Message {
	m.Payload.Badge = number
	return m
}
func (m *Message) Title(title string) *Message {
	m.Alert.Title = title
	return m
}
func (m *Message) Body(body string) *Message {
	m.Alert.Body = body
	return m
}
