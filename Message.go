package goAPNS

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
	m.Alert.title = title
	return m
}
func (m *Message) Body(body string) *Message {
	m.Alert.body = body
	return m
}
