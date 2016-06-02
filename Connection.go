package goapns

import (
	"crypto/tls"
	"net/http"
	//"golang.org/x/net/http2"

	"fmt"
)

type Connection struct {
	HTTPClient  http.Client
	Certificate tls.Certificate
	Host        string
}

// Apple HTTP/2 Development & Production urls
const (
	HostDevelopment = "https://api.development.push.apple.com"
	HostProduction  = "https://api.push.apple.com"
)

func NewConnection(pathname string, key string) (*Connection, error) {
	c := &Connection{}

	cert, err := CertificateFromP12(pathname, key)
	if err != nil {
		fmt.Printf("Error creating Connection: %v", err)
		return nil, err
	}
	c.Certificate = cert
	//Default to Development Host.
	c.Host = HostDevelopment
	return c, nil
}

func (c *Connection) Development() *Connection {
	c.Host = HostDevelopment
	return c
}
func Production(c *Connection) *Connection {
	c.Host = HostProduction
	return c
}
