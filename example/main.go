package main

import (
	"fmt"

	"github.com/tantalum73/Go-APNS"
)

func main() {
	m := goapns.NewMessage().Badge(42).Title("Title").Body("A Test notification :)").Sound("Default")
	m.Custom("key", "val")
	c, err := goapns.NewConnection("../../../../Push Test Push Cert.p12", "PasswortdesZertifikates")
	if err != nil {
		fmt.Println("Error loading cert :(")
	} else {
		c.Development()
	}
	//fmt.Println(m)
	//iPad: 791660155ff167aa766730228fd33f4b0f22d83087448d106ef0a717ef5b2407
	//iPhone: 428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0
	tokens := []string{"791660155ff167aa766730228fd33f4b0f22d83087448d106ef0a717ef5b2407",
		"428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0"}
	ch := make(chan goapns.Response, len(tokens))
	c.Push(m, tokens, ch)

	for response := range ch {
		//fmt.Printf("\nReceived response: %v\n", response)
		if !response.Sent() {
			fmt.Printf("\nThere was an error sending to device %v : %v\n Timestamp: %v, == Zero %v\n\n", response.Token, response.Error, response.Timestamp(), response.Timestamp().IsZero())
		} else {
			fmt.Printf("\nPush successful for token: %v\n", response.Token)
		}

	}
}
