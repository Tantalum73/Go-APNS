package goapns

//Alert stores properties that belong to the Alert as specified by Apple:
// If this property is included, the system displays a standard alert or a banner, based on the user’s setting.
// You can specify a string or a dictionary as the value of alert.
// If you specify a string, it becomes the message text of an alert with two buttons: Close and View. If the user taps View, the app launches.
// If you specify a dictionary, refer to Table 5-2 for descriptions of the keys of this dictionary.
type Alert struct {
	//Title is a short string describing the purpose of the notification.
	// Apple Watch displays this string as part of the notification interface.
	// This string is displayed only briefly and should be crafted so that it can be understood quickly.
	// This key was added in iOS 8.2.
	Title string `json:"title,omitempty"`

	// TitleLocKey is the key to a title string in the Localizable.strings file for the current localization.
	// The key string can be formatted with %@ and %n$@ specifiers to take the variables specified in the title-loc-args array.
	// See Localized Formatted Strings for more information.
	// This key was added in iOS 8.2.
	TitleLocKey string `json:"title-loc-key,omitempty"`

	//TitleLocArgs are variable string values to appear in place of the format specifiers in title-loc-key.
	//See Localized Formatted Strings for more information.
	//This key was added in iOS 8.2
	TitleLocArgs []string `json:"title-loc-args,omitempty"`

	//Body is the text of the alert message.
	Body string `json:"body,omitempty"`

	//LocKey is a key to an alert-message string in a Localizable.strings file for the current localization (which is set by the user’s language preference).
	// The key string can be formatted with %@ and %n$@ specifiers to take the variables specified in the loc-args array.
	// See Localized Formatted Strings for more information.
	LocKey string `json:"loc-key,omitempty"`

	//LocArgs are variable string values to appear in place of the format specifiers in loc-key.
	//See Localized Formatted Strings for more information.
	LocArgs []string `json:"loc-args,omitempty"`

	//ActionLocKey: If a string is specified, the system displays an alert that includes the Close and View buttons.
	//The string is used as a key to get a localized string in the current localization to use for the right button’s title instead of “View”.
	//See Localized Formatted Strings for more information.
	ActionLocKey string `json:"action-loc-key,omitempty"`

	//LaunchImage is the filename of an image file in the app bundle, with or without the filename extension.
	// The image is used as the launch image when users tap the action button or move the action slider.
	// If this property is not specified, the system either uses the previous snapshot,uses the image identified by the UILaunchImageFile key in the app’s Info.plist file, or falls back to Default.png.
	// This property was added in iOS 4.0.
	LaunchImage string `json:"launch-image,omitempty"`
}

//NewAlert creates a new Alert with empty values.
func NewAlert() Alert {
	a := Alert{}
	return a
}
