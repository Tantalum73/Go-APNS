package goapns

import (
	"errors"
	"net/http"
	"time"
)

// Service error responses.
var (
	ErrorPayloadEmpty              = errors.New("The message payload was empty, expected HTTP/2 status code is 400.")
	ErrorPayloadTooLarge           = errors.New("The message payload was too large. The maximum payload size is 4096 bytes.")
	ErrorBadTopic                  = errors.New("The apns-topic was invalid.")
	ErrorTopicDisallowed           = errors.New("Pushing to this topic is not allowed.")
	ErrorBadMessageID              = errors.New("The APNSID value is bad.")
	ErrorBadExpirationDate         = errors.New("The expiration value is bad.")
	ErrorBadPriority               = errors.New("The apns-priority value is bad.")
	ErrorMissingDeviceToken        = errors.New("The device token is not specified in the request. Verify that the message is sent to a device token.")
	ErrorBadDeviceToken            = errors.New("The specified device token was bad. Verify that the request contains a valid token and that the token matches the environment.")
	ErrorDeviceTokenNotForTopic    = errors.New("The device token does not match the specified topic.")
	ErrorUnregistered              = errors.New("The device token is inactive for the specified topic. Expected HTTP/2 status code is 410.")
	ErrorDuplicateHeaders          = errors.New("One or more headers were repeated.")
	ErrorBadCertificateEnvironment = errors.New("The client certificate was for the wrong environment.")
	ErrorBadCertificate            = errors.New("The certificate was bad.")
	ErrorForbidden                 = errors.New("The specified action is not allowed.")
	ErrorBadPath                   = errors.New("The request contained a bad :path value.")
	ErrorMethodNotAllowed          = errors.New("The specified :method was not POST.")
	ErrorTooManyRequests           = errors.New("Too many requests were made consecutively to the same device token.")
	ErrorIdleTimeout               = errors.New("Idle time out.")
	ErrorShutdown                  = errors.New("The server is shutting down.")
	ErrorInternalServerError       = errors.New("An internal server error occurred.")
	ErrorServiceUnavailable        = errors.New("The service is unavailable.")
	ErrorMissingTopic              = errors.New("The apns-topic header of the request was not specified and was required. The apns-topic header is mandatory when the client is connected using a certificate that supports multiple topics.")

	// HTTP Status errors.

	ErrorBadRequest = errors.New("Bad request.")
	ErrorUnknown    = errors.New("Unknown error.")
)

// The possible Reason error codes returned from APNs.
// From table 6-6 in the Apple Local and Remote Notification Programming Guide.
var errorReason = map[string]error{
	"PayloadEmpty":              ErrorPayloadEmpty,
	"PayloadTooLarge":           ErrorPayloadTooLarge,
	"BadTopic":                  ErrorBadTopic,
	"TopicDisallowed":           ErrorTopicDisallowed,
	"BadMessageId":              ErrorBadMessageID,
	"BadExpirationDate":         ErrorBadExpirationDate,
	"BadPriority":               ErrorBadPriority,
	"MissingDeviceToken":        ErrorMissingDeviceToken,
	"BadDeviceToken":            ErrorBadDeviceToken,
	"DeviceTokenNotForTopic":    ErrorDeviceTokenNotForTopic,
	"Unregistered":              ErrorUnregistered,
	"DuplicateHeaders":          ErrorDuplicateHeaders,
	"BadCertificateEnvironment": ErrorBadCertificateEnvironment,
	"BadCertificate":            ErrorBadCertificate,
	"Forbidden":                 ErrorForbidden,
	"BadPath":                   ErrorBadPath,
	"MethodNotAllowed":          ErrorMethodNotAllowed,
	"TooManyRequests":           ErrorTooManyRequests,
	"IdleTimeout":               ErrorIdleTimeout,
	"Shutdown":                  ErrorShutdown,
	"InternalServerErroror":     ErrorInternalServerError,
	"ServiceUnavailable":        ErrorServiceUnavailable,
	"MissingTopic":              ErrorMissingTopic,
}

var errorStatus = map[int]error{
	http.StatusBadRequest:            ErrorBadRequest,
	http.StatusForbidden:             ErrorForbidden,
	http.StatusMethodNotAllowed:      ErrorMethodNotAllowed,
	http.StatusGone:                  ErrorUnregistered,
	http.StatusRequestEntityTooLarge: ErrorPayloadTooLarge,
	http.StatusTooManyRequests:       ErrorTooManyRequests,
	http.StatusInternalServerError:   ErrorInternalServerError,
	http.StatusServiceUnavailable:    ErrorServiceUnavailable,
}

//Response defines properties that are useful to inform the calling script what happened with the
//request to Apples Servers. It defines a StatusCode, a Reason (if on is provided by Apple),
//a Timestamt that can be used to identify since when a device became unavailable
//(of type int64, if you need a time.Time object, please use Timestamt() method),
//an Error (nil if everything went fine), the Token of the device the notification was sent to
//and the Message object itself (can be used to figure out, what went wrong).
type Response struct {
	StatusCode      int
	Reason          string `json:"reason"`
	TimestempNumber int64  `json:"timestamp"`
	Error           error
	Token           string
	Message         *Message
}

//Sent return true if the notification was sent successfully (http status code == 200).
//Additionally, the Error of the Response will be nil.
func (r *Response) Sent() bool {
	return r.StatusCode == http.StatusOK
}

//Timestamp converts the int64 type from a Response into a time.Time object.
func (r *Response) Timestamp() time.Time {
	// if r.TimestempNumber != 0 {
	return time.Unix(r.TimestempNumber/1000, 0)
	// }

}
