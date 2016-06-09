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

//Badge is the number to display as the badge of the app icon.
//This method sets the value to its underlaying Payload object.
//If this property is absent, the badge is not changed.
//To remove the badge, set the value of this property to 0.
func (m *Message) Badge(number int) *Message {
	m.Payload.Badge = number
	return m
}

//NoBadgeChange is a method that resets the badge.
//This method sets the value to its underlaying Payload object.
//If the Badge value is ommitted, it stays unchanged on the app.
//Use this method if you set the Badge (by accident) and want to unset it to let it unchanged.
func (m *Message) NoBadgeChange() *Message {
	m.Payload.Badge = -1
	return m
}

//Sound specified tha name of a sound file in the app bundle or in the Library/Sounds folder of the app’s data container.
//This method sets the value to its underlaying Payload object.
//The sound in this file is played as an alert.
//If the sound file doesn’t exist or default is specified as the value, the default alert sound is played.
//The audio must be in one of the audio data formats that are compatible with system sounds.
func (m *Message) Sound(sound string) *Message {
	m.Payload.Sound = sound
	return m
}

// Category: provide this key with a string value that represents the identifier property of the UIMutableUserNotificationCategory object you created to define custom actions.
//This method sets the value to its underlaying Payload object.
// To learn more about using custom actions, see Registering Your Actionable Notification Types.
func (m *Message) Category(category string) *Message {
	m.Payload.Category = category
	return m
}

//ContentAvailable: if this key is provided with a value of 1 to indicate that new content is available.
//This method sets the value to its underlaying Payload object.
//Including this key and value means that when your app is launched in the background or resumed,
//application:didReceiveRemoteNotification:fetchCompletionHandler: is called.
//This method sets ContentAvailable to 1 and the priority to Low according to Apples documentation.
func (m *Message) ContentAvailable() *Message {
	m.Payload.ContentAvailable = 1
	m.Header.Priority = PriorityLow
	return m
}

//ContentUnavailable lets you set the ContentAvailable flag to 0 and the priority to High.
//This method sets the value to its underlaying Payload object.
//Use this method if you set ContentAvailable() by accident.
func (m *Message) ContentUnavailable() *Message {
	m.Payload.ContentAvailable = 0
	m.Header.Priority = PriorityHigh
	return m
}

/******************************
Configuring Header: APNSID, Expiration, Priority, Topic
******************************/

//APNSID is a canonical UUID that identifies the notification.
//This method sets the value to its underlaying Header object.
//If there is an error sending the notification, APNs uses this value to identify the notification to your server.
//The canonical form is 32 lowercase hexadecimal digits, displayed in five groups separated by hyphens in the form 8-4-4-4-12.
//An example UUID is as follows:
//123e4567-e89b-12d3-a456-42665544000
//If you omit this header, a new UUID is created by APNs and returned in the response.
func (m *Message) APNSID(id string) *Message {
	m.Header.APNSID = id
	return m
}

//Expiration is a UNIX epoch date expressed in seconds (UTC).
//This method sets the value to its underlaying Header object.
//This header identifies the date when the notification is no longer valid and can be discarded.
//If this value is nonzero, APNs stores the notification and tries to deliver it at least once,
//repeating the attempt as needed if it is unable to deliver the notification the first time.
//If the value is 0, APNs treats the notification as if it expires immediately and does not store the notification or attempt to redeliver it.
func (m *Message) Expiration(time time.Time) *Message {
	m.Header.Expiration = time
	return m
}

//PriorityHigh sepcifies a high priority for the notification.
//This method sets the value to its underlaying Header object.
//Default priotity is PriorityHigh so if you don't specify any priority,
//PriorityHigh is assumed.
//
//It sends the push message immediately.
//Notifications with this priority must trigger an alert, sound, or badge on the target device.
//It is an error to use this priority for a push notification that contains only the content-available key
func (m *Message) PriorityHigh() *Message {
	m.Header.Priority = PriorityHigh
	return m
}

//PriorityLow sepcifies a low priority for the notification.
//This method sets the value to its underlaying Header object.
//Default priotity is PriorityHigh so if you don't specify any priority,
//PriorityHigh is assumed.
//
//Send the push message at a time that takes into account power considerations for the device.
//Notifications with this priority might be grouped and delivered in bursts.
//They are throttled, and in some cases are not delivered.
func (m *Message) PriorityLow() *Message {
	m.Header.Priority = PriorityLow
	return m
}

//Topic of the remote notification, which is typically the bundle ID for your app.
//This method sets the value to its underlaying Header object.
//The certificate you create in Member Center must include the capability for this topic.
//If your certificate includes multiple topics, you must specify a value for this header.
//If you omit this header and your APNs certificate does not specify multiple topics,
//the APNs server uses the certificate’s Subject as the default topic.
func (m *Message) Topic(topic string) *Message {
	m.Header.Topic = topic
	return m
}

/******************************
Custom parameter
******************************/
//Custom lets you set a custom key and value. It will be appended to the Notification.
//In your AppDelegate, you can extract those custom values.
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
