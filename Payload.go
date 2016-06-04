package goapns

//Payload defines properties that are described as in Apples documentation.
//Badge, Sound, ContentAvailable and Category are those.
//Think of Payload as a meta-object to your notification as it specify the behaviour
//but not the actual alert.
type Payload struct {
	//Badge is the number to display as the badge of the app icon.
	//If this property is absent, the badge is not changed.
	//To remove the badge, set the value of this property to 0.
	Badge int

	//Sound specified tha name of a sound file in the app bundle or in the Library/Sounds folder of the app’s data container.
	//The sound in this file is played as an alert.
	//If the sound file doesn’t exist or default is specified as the value, the default alert sound is played.
	//The audio must be in one of the audio data formats that are compatible with system sounds.
	Sound string

	//ContentAvailable: if this key is provided with a value of 1 to indicate that new content is available.
	//Including this key and value means that when your app is launched in the background or resumed,
	//application:didReceiveRemoteNotification:fetchCompletionHandler: is called.
	ContentAvailable int

	// Category: provide this key with a string value that represents the identifier property of the UIMutableUserNotificationCategory object you created to define custom actions.
	// To learn more about using custom actions, see Registering Your Actionable Notification Types.
	Category string
}

//NewPayload provides a initializer of Payload with empty values and no badge.
func NewPayload() Payload {
	p := Payload{-1, "", 0, ""}
	return p
}

//MapInto is passed in a map on which the Payload content is appended to.
//It return a new map with every property and key set, ready to build a JSON from it.
func (p *Payload) MapInto(mapped map[string]interface{}) map[string]interface{} {
	if p.Badge >= 0 {
		//Only set badge if the user specified so (by setting a >= 0 value).
		//If not, Badge is ommitted in JSON
		//and therefore the badge on the app is unchanged
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
