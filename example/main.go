package main

import (
	"fmt"

	"github.com/tantalum73/Go-APNS"
)

func main() {
	m := goapns.NewMessage().Title("Title").Body("A Test notification :)").Sound("Default").Badge(42)
	m.Custom("key", "val")
	c, err := goapns.NewConnection("../../../../Push Test Push Cert.p12", "PasswortdesZertifikates")
	if err != nil {
		fmt.Println("Error loading cert :(")
	} else {
		c.Development()
	}

	//fmt.Println(m)
	//iPad: a26f0000c05286ee6f31756b1b9d05b4a37ad512fabbe266dd21357b376f0e0e
	//iPhone: 428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0
	tokens := []string{"a26f0000c05286ee6f31756b1b9d05b4a37ad512fabbe266dd21357b376f0e0e",
		"428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0"}
	ch := make(chan goapns.Response, len(tokens))
	c.Push(m, tokens, ch)
	fmt.Println(m.JSONstring())

	for response := range ch {
		//fmt.Printf("\nReceived response: %v\n", response)
		if !response.Sent() {
			fmt.Printf("\nThere was an error sending to device %v : %v\n Timestamp: %v, == Zero %v\n\n", response.Token, response.Error, response.Timestamp(), response.Timestamp().IsZero())
		} else {
			fmt.Printf("\nPush successful for token: %v\n", response.Token)
		}

	}
}
