package goapns

import (
	"time"
)

const PriorityHigh = 10
const PriorityLow = 5

type Header struct {
	APNSID     string
	Expiration time.Time
	Priority   int
	Topic      string
}

func NewHeader() *Header {
	h := &Header{"", time.Time{}, PriorityHigh, ""}
	return h
}
