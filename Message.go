package goapns

import (
	"bytes"
	"encoding/json"
	"time"
)

//Message collects Header, Payload and Alert and also provides methods to configure them.
type Message struct {

	//Payload defines properties as they are described in Apples documentation.
	Payload Payload

	//Alert stores properties that belong to the Alert as specified by Apple.
	Alert Alert

	//Header collects the header fields for the notification.
	Header Header

	//custom stores custom keys and values the user set. It will be passed into your app as a dictionary as the user launches it.
	custom map[string]interface{}
}

//NewMessage creates a new Message with default Alert, Payload and Header objects.
func NewMessage() *Message {
	m := &Message{}
	m.Alert = NewAlert()
	m.Header = NewHeader()
	m.Payload = NewPayload()

	return m
}

/******************************
Configuring Alert:  Body, Title, TitleLocKey, TitleLocArgs, ActionLocKey, LocKey, LocArgs, LaunchImage
******************************/

//Title sets the title of the alert in its underlaying Alert object. It is a short string describing the purpose of the notification.
func (m *Message) Title(title string) *Message {
	m.Alert.Title = title
	return m
}

//Body sets the text of the alert message in the underlaying Alert object.
func (m *Message) Body(body string) *Message {
	m.Alert.Body = body
	return m
}

// TitleLocKey is the key to a title string in the Localizable.strings file for the current localization.
// This method sets the value to its underlaying Alert object.
// The key string can be formatted with %@ and %n$@ specifiers to take the variables specified in the title-loc-args array.
// See Localized Formatted Strings for more information.
// This key was added in iOS 8.2.
func (m *Message) TitleLocKey(key string) *Message {
	m.Alert.TitleLocKey = key
	return m
}

//TitleLocArgs are variable string values to appear in place of the format specifiers in title-loc-key.
//This method sets the value to its underlaying Alert object.
//See Localized Formatted Strings for more information.
//This key was added in iOS 8.2
func (m *Message) TitleLocArgs(args []string) *Message {
	m.Alert.TitleLocArgs = args
	return m
}

//ActionLocKey: If a string is specified, the system displays an alert that includes the Close and View buttons.
//This method sets the value to its underlaying Alert object.
//The string is used as a key to get a localized string in the current localization to use for the right button’s title instead of “View”.
//See Localized Formatted Strings for more information.
func (m *Message) ActionLocKey(key string) *Message {
	m.Alert.ActionLocKey = key
	return m
}

//LocKey is a key to an alert-message string in a Localizable.strings file for the current localization (which is set by the user’s language preference).
//This method sets the value to its underlaying Alert object.
//The key string can be formatted with %@ and %n$@ specifiers to take the variables specified in the loc-args array.
//See Localized Formatted Strings for more information.
func (m *Message) LocKey(key string) *Message {
	m.Alert.LocKey = key
	return m
}

//LocArgs are variable string values to appear in place of the format specifiers in loc-key.
//This method sets the value to its underlaying Alert object.
//See Localized Formatted Strings for more information.
func (m *Message) LocArgs(args []string) *Message {
	m.Alert.LocArgs = args
	return m
}

//LaunchImage is the filename of an image file in the app bundle, with or without the filename extension.
//This method sets the value to its underlaying Alert object.
//The image is used as the launch image when users tap the action button or move the action slider.
//If this property is not specified, the system either uses the previous snapshot,uses the image identified by the UILaunchImageFile key in the app’s Info.plist file, or falls back to Default.png.
//This property was added in iOS 4.0.
func (m *Message) LaunchImage(imageName string) *Message {
	m.Alert.LaunchImage = imageName
	return m
}

/******************************
Configuring Payload: Badge, Sound, ContentAvailable, Category
******************************/
func (m *Message) Badge(number int) *Message {
	m.Payload.Badge = number
	return m
}
func (m *Message) NoBadgeChange() *Message {
	m.Payload.Badge = -1
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
func (m *Message) ContentUnavailable() *Message {
	m.Payload.ContentAvailable = 0
	m.Header.Priority = PriorityHigh
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
//MarshalJSON builds a []byte that stores the Message object in JSON.
func (m *Message) MarshalJSON() ([]byte, error) {
	payload := make(map[string]interface{}, 4)
	payload["alert"] = m.Alert
	payload = m.Payload.MapInto(payload)

	jsonMappedWithAPSKey := map[string]interface{}{"aps": payload}

	for key, object := range m.custom {
		jsonMappedWithAPSKey[key] = object
	}

	return json.Marshal(jsonMappedWithAPSKey)
}

//JSONstring returns the entire Message object as JSON exactly as it will
//be sent to Apples servers.
//You can use this method to debug your code.
func (m *Message) JSONstring() string {
	b, _ := json.Marshal(m)
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, b, "", "\t")
	if error != nil {
	}
	return string(prettyJSON.Bytes())
}
