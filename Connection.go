package goapns

import (
	"bytes"
	"crypto/tls"
	"net/http"

	"encoding/json"
	"fmt"

	"golang.org/x/net/http2"
)

//Connection stores properties that are necessary to perform a request to Apples servers.
type Connection struct {
	//HTTPClient the HTTP client that is used. It is configured with a tls.Config that uses
	//the specified certificate. It is constructed for you in NewConnection(pathname, key)
	HTTPClient http.Client
	//Certificate is the certificate that you specified during construction of the Connection
	//by using NewConnection(pathname string, key string)
	Certificate tls.Certificate
	//Host is the host to which the request is sent to.
	Host string
}

// Apple HTTP/2 Development & Production urls
const (
	HostDevelopment = "https://api.development.push.apple.com"
	HostProduction  = "https://api.push.apple.com"
)

//NewConnection creates a new Connection object. A Certificate is required to
//send requests to Apples servers. You can specify the path to a .p12 certificate
//and its passphrase.
//The default host is the development host. connection.Production() if you want to
//use the production environment.
//It will return a *Connection or an error. One of this is always nil.
func NewConnection(pathname string, key string) (*Connection, error) {
	c := &Connection{}

	cert, err := CertificateFromP12(pathname, key)
	if err != nil {
		//fmt.Printf("Error creating Connection: %v", err)
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

//Development sets the host to Apples development environment.
//Use this while your app is in development and not published.
//This host is set by default.
func (c *Connection) Development() *Connection {
	c.Host = HostDevelopment
	return c
}

//Production sets the host to Apples development environment.
//Use this while your app is in production.
func (c *Connection) Production() *Connection {
	c.Host = HostProduction
	return c
}

//Push sends a request to Apples servers.
//It takes a Message and and tokens in an array to which the message should be pushed to.
//The result of a request is pushed into the responseChannel as an Response object.
//As the network operation is performed asynchronously (using go keyword)
//the method will return immediately. Use responseChannel to watch the results.
//You will get one Response object for every request that is sent (one request per token).
func (c *Connection) Push(message *Message, tokens []string, responseChannel chan Response) {
	// fmt.Printf("Will push to tokens %v , URL: %v\n", tokens, c.Host)
	dataToSend, err := json.Marshal(message)

	if err != nil {
		fmt.Printf("Error JSONING the request: %v\naborting\n", err)
		close(responseChannel)
		return
	}

	for index, token := range tokens {

		url := fmt.Sprintf("%v/3/device/%v", c.Host, token)
		request, err := http.NewRequest("POST", url, bytes.NewBuffer(dataToSend))
		if err != nil {
			fmt.Printf("Error creating request: %v\naborting\n", err)
			response := Response{}
			response.Error = err
			response.Message = message
			responseChannel <- response
			continue
		}

		configureHeader(request, message)
		push := func(token string, responseChannel chan Response, shouldCloseChannelWhenDone bool) {
			if shouldCloseChannelWhenDone {
				defer close(responseChannel)
			}

			httpResponse, err := c.HTTPClient.Do(request)
			if httpResponse != nil {
				defer httpResponse.Body.Close()
			}

			if err != nil {
				fmt.Printf("Error during response: %v\nAborting.\n", err)

				response := Response{}
				response.Error = err
				response.Message = message
				response.Token = token
				responseChannel <- response
			} else {

				//Response object that will be populated and passed into the responseChannel
				var response Response

				if httpResponse.StatusCode != http.StatusOK {
					//Something went wrong, creating new Response object from the JSON response
					errParsingJSON := json.NewDecoder(httpResponse.Body).Decode(&response)

					if errParsingJSON != nil {
						//We could not parse the response into JSON, we need to pass the received error into the responseChannel
						response.Error = err

					} else {
						//We have parsed the error and populated a new Response object with it.

						//Converting the JSON body (string) into an error object
						knownError, found := errorReason[response.Reason]

						if !found {
							//We could not find the error in our map so we try to use the HTTP status code to produce some meaningful error object
							knownError, found = errorStatus[httpResponse.StatusCode]

							if !found {
								//Could not find the error anywhere :(
								knownError = ErrorUnknown
							}
						}
						response.Error = knownError
					}
				}

				response.Message = message
				response.Token = token
				response.StatusCode = httpResponse.StatusCode
				responseChannel <- response
			}
		}
		shouldCloseChannelWhenDone := index == (len(tokens) - 1)
		go push(token, responseChannel, shouldCloseChannelWhenDone)

	}

}

//configureHader takes a Message and a htto.Request. It sets the header properties
//of it as Apples documentation says. Therefore, it writes values from the
//message.Header object into a http header.
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
	if message.Header.CollapseID != "" {
		request.Header.Set("apns-collapse-id", message.Header.CollapseID)
	}
}
