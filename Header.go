package goapns

import (
	"time"
)

//PriorityHigh sepcifies a high priority for the notification.
//Default priotity is PriorityHigh so if you don't specify any priority,
//PriorityHigh is assumed.
//
//It sends the push message immediately.
//Notifications with this priority must trigger an alert, sound, or badge on the target device.
//It is an error to use this priority for a push notification that contains only the content-available key
const PriorityHigh = 10

//PriorityLow sepcifies a low priority for the notification.
//Default priotity is PriorityHigh so if you don't specify any priority,
//PriorityHigh is assumed.
//
//Send the push message at a time that takes into account power considerations for the device.
//Notifications with this priority might be grouped and delivered in bursts.
//They are throttled, and in some cases are not delivered.
const PriorityLow = 5

//Header collects the header fields for the notification.
type Header struct {

	//APNSID is a canonical UUID that identifies the notification.
	//If there is an error sending the notification, APNs uses this value to identify the notification to your server.
	//The canonical form is 32 lowercase hexadecimal digits, displayed in five groups separated by hyphens in the form 8-4-4-4-12.
	//An example UUID is as follows:
	//123e4567-e89b-12d3-a456-42665544000
	//If you omit this header, a new UUID is created by APNs and returned in the response.
	APNSID string

	//Expiration is a UNIX epoch date expressed in seconds (UTC).
	//This header identifies the date when the notification is no longer valid and can be discarded.
	//If this value is nonzero, APNs stores the notification and tries to deliver it at least once,
	//repeating the attempt as needed if it is unable to deliver the notification the first time.
	//If the value is 0, APNs treats the notification as if it expires immediately and does not store the notification or attempt to redeliver it.
	Expiration time.Time

	//Priority is the priority of the notification.
	//Set it to PriorityHigh or PriorityLow but do not
	//use custom numbers in any case.
	Priority int

	//Topic of the remote notification, which is typically the bundle ID for your app.
	//The certificate you create in Member Center must include the capability for this topic.
	//If your certificate includes multiple topics, you must specify a value for this header.
	//If you omit this header and your APNs certificate does not specify multiple topics,
	//the APNs server uses the certificate’s Subject as the default topic.
	Topic string

	//CollapseID specifies a string that is used to replace a former sent notification by the latest one.
	//Notifications with the same CollapseID string will be collapsed so that only the newest notification is displayed.
	//Usefull if you want to present a  scrore of a fotball math or something that gets frequently updated.
	CollapseID string
}

//NewHeader creates a new Header with high priority.
func NewHeader() Header {
	h := Header{"", time.Time{}, PriorityHigh, "", ""}
	return h
}
