package goapns_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tantalum73/Go-APNS"
)

func TestMessageSetter(t *testing.T) {
	m := goapns.NewMessage()

	m.Badge(42)
	assert.Equal(t, 42, m.Payload.Badge, "Badge not set correctly")

	//Testing default priority
	assert.Equal(t, 10, m.Header.Priority)

	m.NoBadgeChange()
	assert.Equal(t, -1, m.Payload.Badge, "Badge not re-set correctly")

	m.ContentAvailable()
	assert.Equal(t, 1, m.Payload.ContentAvailable, "ContentAvailable not set correctly")
	assert.Equal(t, 5, m.Header.Priority, "ContentAvailable not set correctly")

	m.ContentUnavailable()
	assert.Equal(t, 0, m.Payload.ContentAvailable, "ContentUnavailable not set correctly")
	assert.Equal(t, 10, m.Header.Priority, "ContentUnavailable not set correctly")

	m.Body("body")
	assert.Equal(t, "body", m.Alert.Body)

	m.MutableContent()
	assert.Equal(t, 1, m.Payload.MutableContent)

}
