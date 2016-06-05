package goapns

import (
	"bytes"
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
		fmt.Printf("Erroror creating Connection: %v", err)
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

func (c *Connection) Push(message *Message, tokens []string, responseChannel chan Response) {
	fmt.Printf("Will push to tokens %v \n", tokens)
	dataToSend, err := json.Marshal(message)

	if err != nil {
		fmt.Printf("Erroror JSONING the response: %v\naborting\n", err)
		close(responseChannel)
		return
	}

	for index, token := range tokens {

		url := fmt.Sprintf("%v/3/device/%v", c.Host, token)
		request, err := http.NewRequest("POST", url, bytes.NewBuffer(dataToSend))
		if err != nil {
			fmt.Printf("Erroror creating request: %v\naborting\n", err)
			response := Response{}
			response.Erroror = err
			response.Message = message
			responseChannel <- response
			continue
		}

		configureHeader(request, message)
		push := func(token string, responseChannel chan Response) {

			httpResponse, err := c.HTTPClient.Do(request)
			if err != nil {
				fmt.Printf("Erroror during response: %v\naborting\n", err)

				response := Response{}
				response.Erroror = err
				response.Message = message
				responseChannel <- response
			} else {
				defer httpResponse.Body.Close()

				//Response object that will be populated and passed into the responseChannel
				var response Response

				if httpResponse.StatusCode != http.StatusOK {
					//Something went wrong, creating new Response object from the JSON response
					errParsingJSON := json.NewDecoder(httpResponse.Body).Decode(&response)
					if errParsingJSON != nil {
						//We could not parse the response into JSON, we need to pass the received error into the responseChannel
						response.Erroror = err

					} else {
						//We have parsed the error and populated a new Response object with it.

						//Converting the JSON body (string) into an error object
						knownErroror, found := errorReason[response.Reason]

						if !found {
							//We could not find the error in our map so we try to use the HTTP status code to produce some meaningful error object
							knownErroror, found = errorStatus[httpResponse.StatusCode]

							if !found {
								//Could not find the error anywhere :(
								knownErroror = ErrorUnknown
							}
						}
						response.Erroror = knownErroror
					}
				}

				response.Message = message
				response.Token = token
				response.StatusCode = httpResponse.StatusCode
				responseChannel <- response
			}
		}
		go push(token, responseChannel)
		if index == len(tokens) {
			close(responseChannel)
		}
	}

}

func configureHeader(request *http.Request, message *Message) {
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	if message.Header.APNSID != "" {
		request.Header.Set("apns-id", message.Header.APNSID)
	}

	if !message.Header.Expiration.IsZero() {
		request.Header.Set("apns-expiration", fmt.Sprintf("%v", message.Header.Expiration.Unix()))
	}
	//Only set the priority if it is low because high is the default
	//value that is assumed when no priority is specified.
	//We want to omit everything we can to save bandwith.
	if message.Header.Priority == PriorityLow {
		request.Header.Set("apns-priority", fmt.Sprintf("%v", message.Header.Priority))
	}
	if message.Header.Topic != "" {
		request.Header.Set("apns-topic", message.Header.Topic)
	}
}
