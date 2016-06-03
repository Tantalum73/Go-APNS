package goapns

import "encoding/json"

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
func (m *Message) PriorityHigh() *Message {
	m.Header.Priority = PriorityHigh
	return m
}
func (m *Message) PriorityLow() *Message {
	m.Header.Priority = PriorityLow
	return m
}
func (m *Message) ContentAvailable() *Message {
	m.Payload.ContentAvailable = 1
	m.Header.Priority = PriorityLow
	return m
}

func (m Message) MarshalJSON() ([]byte, error) {
	payload := make(map[string]interface{}, 4)
	payload["alert"] = m.Alert

	return json.Marshal(payload)
}
