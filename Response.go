package goapns

import (
	"errors"
	"net/http"
	"time"
)

// Service error responses.
var (
	// These could be checked prior to sending the request to Apple.

	ErrPayloadEmpty              = errors.New("The message payload was empty, expected HTTP/2 status code is 400.")
	ErrPayloadTooLarge           = errors.New("The message payload was too large. The maximum payload size is 4096 bytes.")
	ErrBadTopic                  = errors.New("The apns-topic was invalid.")
	ErrTopicDisallowed           = errors.New("Pushing to this topic is not allowed.")
	ErrBadMessageID              = errors.New("The APNSID value is bad.")
	ErrBadExpirationDate         = errors.New("The expiration value is bad.")
	ErrBadPriority               = errors.New("The apns-priority value is bad.")
	ErrMissingDeviceToken        = errors.New("The device token is not specified in the request. Verify that the message is sent to a device token.")
	ErrBadDeviceToken            = errors.New("The specified device token was bad. Verify that the request contains a valid token and that the token matches the environment.")
	ErrDeviceTokenNotForTopic    = errors.New("The device token does not match the specified topic.")
	ErrUnregistered              = errors.New("The device token is inactive for the specified topic. Expected HTTP/2 status code is 410.")
	ErrDuplicateHeaders          = errors.New("One or more headers were repeated.")
	ErrBadCertificateEnvironment = errors.New("The client certificate was for the wrong environment.")
	ErrBadCertificate            = errors.New("The certificate was bad.")
	ErrForbidden                 = errors.New("The specified action is not allowed.")
	ErrBadPath                   = errors.New("The request contained a bad :path value.")
	ErrMethodNotAllowed          = errors.New("The specified :method was not POST.")
	ErrTooManyRequests           = errors.New("Too many requests were made consecutively to the same device token.")
	ErrIdleTimeout               = errors.New("Idle time out.")
	ErrShutdown                  = errors.New("The server is shutting down.")
	ErrInternalServerError       = errors.New("An internal server error occurred.")
	ErrServiceUnavailable        = errors.New("The service is unavailable.")
	ErrMissingTopic              = errors.New("The apns-topic header of the request was not specified and was required. The apns-topic header is mandatory when the client is connected using a certificate that supports multiple topics.")

	// HTTP Status errors.

	ErrBadRequest = errors.New("bad request")
	ErrUnknown    = errors.New("unknown error")
)

// The possible Reason error codes returned from APNs.
// From table 6-6 in the Apple Local and Remote Notification Programming Guide.
var errorReason = map[string]error{
	"PayloadEmpty":              ErrPayloadEmpty,
	"PayloadTooLarge":           ErrPayloadTooLarge,
	"BadTopic":                  ErrBadTopic,
	"TopicDisallowed":           ErrTopicDisallowed,
	"BadMessageId":              ErrBadMessageID,
	"BadExpirationDate":         ErrBadExpirationDate,
	"BadPriority":               ErrBadPriority,
	"MissingDeviceToken":        ErrMissingDeviceToken,
	"BadDeviceToken":            ErrBadDeviceToken,
	"DeviceTokenNotForTopic":    ErrDeviceTokenNotForTopic,
	"Unregistered":              ErrUnregistered,
	"DuplicateHeaders":          ErrDuplicateHeaders,
	"BadCertificateEnvironment": ErrBadCertificateEnvironment,
	"BadCertificate":            ErrBadCertificate,
	"Forbidden":                 ErrForbidden,
	"BadPath":                   ErrBadPath,
	"MethodNotAllowed":          ErrMethodNotAllowed,
	"TooManyRequests":           ErrTooManyRequests,
	"IdleTimeout":               ErrIdleTimeout,
	"Shutdown":                  ErrShutdown,
	"InternalServerError":       ErrInternalServerError,
	"ServiceUnavailable":        ErrServiceUnavailable,
	"MissingTopic":              ErrMissingTopic,
}

var errorStatus = map[int]error{
	http.StatusBadRequest:            ErrBadRequest,
	http.StatusForbidden:             ErrForbidden,
	http.StatusMethodNotAllowed:      ErrMethodNotAllowed,
	http.StatusGone:                  ErrUnregistered,
	http.StatusRequestEntityTooLarge: ErrPayloadTooLarge,
	http.StatusTooManyRequests:       ErrTooManyRequests,
	http.StatusInternalServerError:   ErrInternalServerError,
	http.StatusServiceUnavailable:    ErrServiceUnavailable,
}

const Success = http.StatusOK

type Response struct {
	StatusCode      int
	Reason          string `json:"reason"`
	TimestempNumber int64  `json:"timestamp"`
	Message         *Message
	Error           error
	Token           string
}

func (r *Response) Sent() bool {
	return r.StatusCode == Success
}
func (r *Response) Timestamp() time.Time {
	// if r.TimestempNumber != 0 {
	return time.Unix(r.TimestempNumber/1000, 0)
	// }

}
