package goapns_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tantalum73/Go-APNS"
)

func mockMessage() *goapns.Message {
	m := goapns.NewMessage().Badge(42).Title("title").Body("body")
	return m
}
func mockConnection(t *testing.T) *goapns.Connection {
	pathToCert := "example/certificate-valid-encrypted.p12"
	conn, err := goapns.NewConnection(pathToCert, "password")
	assert.Nil(t, err)
	assert.NotNil(t, conn)
	client := http.Client{}
	//Mocking the client to, because otherwise we would have to use HTTP
	//for testing, which is harder to do. We also do not test security here.
	conn.HTTPClient = client
	return conn
}
func TestConnectionCertificateWrongPath(t *testing.T) {
	pathToCert := "example/nowhere"
	conn, err := goapns.NewConnection(pathToCert, "wrongPassword")
	assert.Error(t, err)
	assert.Nil(t, conn)
}
func TestConnectionCertificateWrongPassphrase(t *testing.T) {
	pathToCert := "example/certificate-valid-encrypted.p12"
	conn, err := goapns.NewConnection(pathToCert, "wrongPassword")
	assert.Error(t, err)
	assert.Nil(t, conn)
}
func TestConnectionCertificateCorrectPassphrase(t *testing.T) {
	pathToCert := "example/certificate-valid-encrypted.p12"
	conn, err := goapns.NewConnection(pathToCert, "password")
	assert.Nil(t, err)
	assert.NotNil(t, conn)
}
func TestConnectionToken(t *testing.T) {
	conn := mockConnection(t)

	token := []string{"1234567890"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, fmt.Sprintf("/3/device/%s", token[0]), r.URL.String())
	}))
	defer server.Close()

	channel := make(chan goapns.Response, 1)

	conn.Host = server.URL
	conn.Push(mockMessage(), token, channel)
	// for response := range channel {
	// 	fmt.Println("Received response")
	// }
}

func TestConnectionHeaderDefaults(t *testing.T) {
	conn := mockConnection(t)

	token := []string{"1234567890"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json; charset=utf-8", r.Header.Get("Content-Type"))
		assert.Equal(t, "", r.Header.Get("apns-id"))
		assert.Equal(t, "", r.Header.Get("apns-priority"))
		assert.Equal(t, "", r.Header.Get("apns-topic"))
		assert.Equal(t, "", r.Header.Get("apns-expiration"))
	}))
	defer server.Close()

	channel := make(chan goapns.Response, 1)

	conn.Host = server.URL
	conn.Push(mockMessage(), token, channel)

}
func TestConnectionHeader(t *testing.T) {
	conn := mockConnection(t)

	token := []string{"1234567890"}
	message := mockMessage().APNSID("102").PriorityHigh().Topic("topic").Expiration(time.Now())

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, message.Header.APNSID, r.Header.Get("apns-id"))
		assert.Equal(t, "10", message.Header.Priority, r.Header.Get("apns-priority"))
		assert.Equal(t, message.Header.Topic, r.Header.Get("apns-topic"))
		assert.Equal(t, fmt.Sprintf("%v", message.Header.Expiration.Unix()), r.Header.Get("apns-expiration"))
	}))
	defer server.Close()

	channel := make(chan goapns.Response, 1)

	conn.Host = server.URL
	conn.Push(message, token, channel)

}
func TestConnectionTokenExpired(t *testing.T) {
	conn := mockConnection(t)

	token := []string{"12345678912"}
	message := mockMessage().APNSID("102").PriorityHigh().Topic("topic").Expiration(time.Now())
	expired := time.Now().UTC().Unix()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("apns-id", message.Header.APNSID)
		w.WriteHeader(http.StatusGone)
		responseJSON := fmt.Sprintf("{\"reason\":\"Unregistered\" ,\"timestamp\":%v}", expired)
		w.Write([]byte(responseJSON))
	}))
	defer server.Close()

	channel := make(chan goapns.Response, 1)

	conn.Host = server.URL
	conn.Push(message, token, channel)
	for response := range channel {
		//Timestamp set correctly
		assert.Equal(t, response.TimestempNumber, expired)
		//Reason set correctly
		assert.IsType(t, goapns.ErrorUnregistered, response.Error)
		assert.False(t, response.Sent())
	}
}

func TestConnectionBadPriority(t *testing.T) {
	conn := mockConnection(t)

	token := []string{"12345678912"}
	message := mockMessage().APNSID("102")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(`{"reason": "BadPriority"}`))
	}))
	defer server.Close()

	channel := make(chan goapns.Response, 1)

	conn.Host = server.URL
	conn.Push(message, token, channel)
	for response := range channel {
		assert.False(t, response.Sent())
		//Reason set correctly
		assert.IsType(t, goapns.ErrorBadPriority, response.Error)
	}
}
