package goapns

import (
	"encoding/json"
	"time"
)

type Message struct {
	Payload Payload
	Alert   Alert
	Header  Header
	custom  map[string]interface{}
}

func NewMessage() *Message {
	return &Message{}
}

/******************************
Configuring Payload: Badge, Sound, ContentAvailable, Category
******************************/
func (m *Message) Badge(number int) *Message {
	m.Payload.Badge = number
	return m
}
func (m *Message) Sound(sound string) *Message {
	m.Payload.Sound = sound
	return m
}
func (m *Message) Category(category string) *Message {
	m.Payload.Category = category
	return m
}
func (m *Message) ContentAvailable() *Message {
	m.Payload.ContentAvailable = 1
	m.Header.Priority = PriorityLow
	return m
}

/******************************
Configuring Header: APNSID, Expiration, Priority, Topic
******************************/
func (m *Message) APNSID(id string) *Message {
	m.Header.APNSID = id
	return m
}
func (m *Message) Expiration(time time.Time) *Message {
	m.Header.Expiration = time
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

/******************************
Configuring Alert:  Body, Title, TitleLocKey, TitleLocArgs, ActionLocKey, LocKey, LocArgs, LaunchImage
******************************/
func (m *Message) Title(title string) *Message {
	m.Alert.Title = title
	return m
}
func (m *Message) Body(body string) *Message {
	m.Alert.Body = body
	return m
}
func (m *Message) TitleLocKey(key string) *Message {
	m.Alert.TitleLocKey = key
	return m
}
func (m *Message) TitleLocArgs(args []string) *Message {
	m.Alert.TitleLocArgs = args
	return m
}
func (m *Message) ActionLocKey(key string) *Message {
	m.Alert.ActionLocKey = key
	return m
}
func (m *Message) LocKey(key string) *Message {
	m.Alert.LocKey = key
	return m
}
func (m *Message) LocArgs(args []string) *Message {
	m.Alert.LocArgs = args
	return m
}
func (m *Message) LaunchImage(imageName string) *Message {
	m.Alert.LaunchImage = imageName
	return m
}

/******************************
Custom parameter
******************************/

func (m *Message) Custom(key string, object interface{}) *Message {
	m.custom = make(map[string]interface{})
	m.custom[key] = object
	return m
}

/******************************
JSON encoding
******************************/

func (m Message) MarshalJSON() ([]byte, error) {
	payload := make(map[string]interface{}, 4)
	payload["alert"] = m.Alert
	payload = m.Payload.MapInto(payload)

	jsonMappedWithAPSKey := map[string]interface{}{"aps": payload}

	for key, object := range m.custom {
		jsonMappedWithAPSKey[key] = object
	}

	return json.Marshal(jsonMappedWithAPSKey)
}
