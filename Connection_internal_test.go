package goapns

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	pathToCert := "example/certificate-valid-encrypted.p12"
	conn, err := NewConnection(pathToCert, "password")
	assert.Nil(t, err)
	assert.NotNil(t, conn)

	id := "APNSID"
	collapseID := "com.example.euroApp.scroreChanged"
	now := time.Now()
	topic := "topic"

	m := NewMessage()
	m.APNSID(id)
	m.Expiration(now)
	//Low will be set, high will be omitted
	m.PriorityLow()
	m.Topic(topic)
	m.CollapseID(collapseID)

	request, _ := http.NewRequest("POST", "", nil)
	assert.NotNil(t, request)
	//configuring the header using constructed Message
	configureHeader(request, m)

	//Testing if header values are set correctly
	testRequestHeader(t, request, "apns-id", id)
	testRequestHeader(t, request, "apns-expiration", fmt.Sprintf("%v", now.Unix()))
	//Low will be set, high will be omitted
	testRequestHeader(t, request, "apns-priority", fmt.Sprintf("%v", 5))
	testRequestHeader(t, request, "apns-topic", topic)
	testRequestHeader(t, request, "apns-collapse-id", collapseID)

	m.PriorityHigh()
	newRequest, _ := http.NewRequest("POST", "", nil)
	configureHeader(newRequest, m)
	testRequestHeader(t, newRequest, "apns-priority", "")

}

func testRequestHeader(t *testing.T, request *http.Request, key string, expected string) {
	actual := request.Header.Get(key)
	assert.Equal(t, expected, actual)
}
