package goapns

import (
	"crypto/tls"
	"net/http"

	"encoding/json"
	"fmt"

	"golang.org/x/net/http2"
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

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	//no check if Certificate is present to fail hard if this requirement is not met
	tlsConfig.BuildNameToCertificate()

	transport := &http2.Transport{TLSClientConfig: tlsConfig}

	c.HTTPClient = http.Client{Transport: transport}
	//Default Host is Development Host.
	c.Host = HostDevelopment
	return c, nil
}

func (c *Connection) Development() *Connection {
	c.Host = HostDevelopment
	return c
}
func (c *Connection) Production() *Connection {
	c.Host = HostProduction
	return c
}

func (c *Connection) Push(message *Message, tokens []string, responseChannel chan string) {
	fmt.Printf("Will push to tokens %v \n", tokens)

	jsonMessage, err := json.Marshal(&message)
	if err != nil {
		fmt.Printf("Error while building JSON: %v \n", err)
	} else {
		fmt.Println(jsonMessage)
	}

	messageFromJSON := &Message{}
	err2 := json.Unmarshal(jsonMessage, messageFromJSON)
	if err != nil {
		fmt.Printf("Error while building Message from JSON: %v \n", err2)
	} else {
		fmt.Println(messageFromJSON)
	}
	for _, token := range tokens {
		responseChannel <- token
	}
	close(responseChannel)
}
