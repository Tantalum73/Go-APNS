package main

import (
	"fmt"
	"log"

	"github.com/tantalum73/Go-APNS"
)

func main() {
	//creating the Connection by passing the path to a valid certificate and its passphrase
	conn, err := goapns.NewConnection("../../../../Push Test Push Cert.p12", "PasswortdesZertifikates")
	if err != nil {
		log.Fatal(err)
	} else {
		conn.Development()
	}

	//composing the Message
	message := goapns.NewMessage().Title("Title").Body("A Test notification :)").Sound("Default").Badge(42)
	message.Custom("key", "val")

	//iPad: a26f0000c05286ee6f31756b1b9d05b4a37ad512fabbe266dd21357b376f0e0e
	//iPhone: 428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0

	//Tokens from a database or, in my case, statically typed
	tokens := []string{"a26f0000c05286ee6f31756b1b9d05b4a37ad512fabbe266dd21357b376f0e0e",
		"428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0"}

	//create a channel that gets the Response object passed in,
	//it expects as many responses as there are token to push to
	channel := make(chan goapns.Response, len(tokens))

	//Print the JSON as it is sent to Apples servers
	fmt.Println(message.JSONstring())

	//Perform the push asynchronosly
	conn.Push(message, tokens, channel)

	//iterate through the responses
	for response := range channel {
		//fmt.Printf("\nReceived response: %v\n", response)
		if !response.Sent() {
			fmt.Printf("\nThere was an error sending to device %v : %v\n Timestamp: %v, == Zero %v\n\n", response.Token, response.Error, response.Timestamp(), response.Timestamp().IsZero())
		} else {
			fmt.Printf("\nPush successful for token: %v\n", response.Token)
		}

	}
}
