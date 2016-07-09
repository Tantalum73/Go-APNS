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

	collapseID := "com.example.euroApp.scroreChanged"
	m.CollapseID(collapseID)
	assert.Equal(t, collapseID, m.Header.CollapseID)

	m.Subtitle("subtitle")
	assert.Equal(t, "subtitle", m.Alert.Subtitle)
}

func TestMessageJSON(t *testing.T) {
	//Filling every entry of the Alert and Payload
	//Then comparing the created JSON that will be sent to Apples servers
	//with a expected JSON.
	array := []string{"1", "2"}

	m := goapns.NewMessage()

	//Filling Alert
	m.Title("Title").Subtitle("Subtitle")
	m.TitleLocKey("Tkey").TitleLocArgs(array)
	m.Body("body")
	m.LocKey("Lkey").LocArgs(array)
	m.ActionLocKey("Akey").LaunchImage("imageName")

	//Filling Payload
	m.Badge(42).Sound("sound")
	m.ContentAvailable()
	m.MutableContent()

	//Hard coded expected json, every field is set
	expected := []byte(`{"aps":{"alert":{"title":"Title","subtitle":"Subtitle","title-loc-key":"Tkey","title-loc-args":["1","2"],"body":"body","loc-key":"Lkey","loc-args":["1","2"],"action-loc-key":"Akey","launch-image":"imageName"},"badge":42,"content-available":1,"mutable-content":1,"sound":"sound"}}`)
	json, _ := m.MarshalJSON()

	assert.Equal(t, expected, json)
}
