package main

import (
	"fmt"

	"github.com/tantalum73/Go-APNS"
)

func main() {
	m := goapns.NewMessage().Badge(42).Title("Title").Body("body")
	c, err := goapns.NewConnection("example/certificate-valid-encrypted.p12", "password")
	if err != nil {

	} else {
		c.Development()
	}
	fmt.Println(m)
	tokens := []string{"1234", "5678"}
	ch := make(chan string, len(tokens))
	c.Push(m, tokens, ch)

	for response := range ch {
		fmt.Println("Received response " + response)
	}
}
