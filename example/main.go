package main

import (
	"fmt"

	"github.com/tantalum73/Go-APNS"
)

func main() {
	m := goapns.NewMessage().Badge(42).Title("Title").Body("body")
	c, err := goapns.NewConnection("../../../../Push Test Push Cert.p12", "PasswortdesZertifikates")
	if err != nil {
		fmt.Println("Error loading cert :(")
	} else {
		c.Development()
	}
	fmt.Println(m)
	//iPad: 791660155ff167aa766730228fd33f4b0f22d83087448d106ef0a717ef5b2407
	//iPhone: 428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0
	tokens := []string{"791660155ff167aa766730228fd33f4b0f22d83087448d106ef0a717ef5b2407",
		"428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0"}
	ch := make(chan string, len(tokens))
	c.Push(m, tokens, ch)

	for response := range ch {
		fmt.Println("Received response " + response)
	}
}
