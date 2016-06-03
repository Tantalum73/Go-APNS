package goapns

import (
	"errors"
	"net/http"
)

// Service error responses.
var (
	// These could be checked prior to sending the request to Apple.

	ErrPayloadEmpty    = errors.New("the message payload was empty")
	ErrPayloadTooLarge = errors.New("the message payload was too large")

	// Device token errors.

	ErrMissingDeviceToken = errors.New("device token was not specified")
	ErrBadDeviceToken     = errors.New("bad device token")
	ErrTooManyRequests    = errors.New("too many requests were made consecutively to the same device token")

	// Header errors.

	ErrBadMessageID      = errors.New("the ID header value is bad")
	ErrBadExpirationDate = errors.New("the Expiration header value is bad")
	ErrBadPriority       = errors.New("the apns-priority value is bad")
	ErrBadTopic          = errors.New("the Topic header was invalid")

	// Certificate and topic errors.

	ErrBadCertificate            = errors.New("the certificate was bad")
	ErrBadCertificateEnvironment = errors.New("certificate was for the wrong environment")
	ErrForbidden                 = errors.New("there was an error with the certificate")

	ErrMissingTopic           = errors.New("the Topic header of the request was not specified and was required")
	ErrTopicDisallowed        = errors.New("pushing to this topic is not allowed")
	ErrUnregistered           = errors.New("device token is inactive for the specified topic")
	ErrDeviceTokenNotForTopic = errors.New("device token does not match the specified topic")

	// These errors should never happen when using Push.

	ErrDuplicateHeaders = errors.New("one or more headers were repeated")
	ErrBadPath          = errors.New("the request contained a bad :path")
	ErrMethodNotAllowed = errors.New("the specified :method was not POST")

	// Fatal server errors.

	ErrIdleTimeout         = errors.New("idle time out")
	ErrShutdown            = errors.New("the server is shutting down")
	ErrInternalServerError = errors.New("an internal server error occurred")
	ErrServiceUnavailable  = errors.New("the service is unavailable")

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
	StatusCode int
	Reason     string `json:"reason"`
	Timestamp  int64  `json:"timestamp"`
	Message    *Message
	Error      error
	Token      string
}

func (r *Response) Sent() bool {
	return r.StatusCode == Success
}
