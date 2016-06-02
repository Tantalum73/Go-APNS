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
	fmt.Println("Will load cert")
	c := &Connection{}

	cert, err := CertificateFromP12(pathname, key)
	if err != nil {
		return nil, err
	}
	c.Certificate = cert
	//Default to Development Host.
	c.Host = HostDevelopment
	return c, nil
}

// func Development() *Connection {
// 	c := &Connection{}
// 	c.Host = HostDevelopment
// 	return c
// }
// func Production() *Connection {
// 	c := &Connection{}
// 	c.Host = HostProduction
// 	return c
// }
