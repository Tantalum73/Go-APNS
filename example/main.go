package main

import (
	"fmt"

	"github.com/tantalum73/Go-APNS"
)

func main() {
	m := goapns.NewMessage().Badge(42).Title("Title").Body("body")
	c, err := goapns.NewConnection("../../../../Push Test Push Cert.p12", "PasswortdesZertifikates")
	if err != nil {

	} else {
		c.Development()
	}
	fmt.Println(m)
	//iPad: e72c54b8ca3fa12ac078846430660b5c031fade782a7ffafc4c3743e18a59bb2
	//iPhone: 428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0
	tokens := []string{"e72c54b8ca3fa12ac078846430660b5c031fade782a7ffafc4c3743e18a59bb2",
		"428dc1d681e576f69f337cd0061b1cdd8da9b76daab39203fa649c26187722c0"}
	ch := make(chan string, len(tokens))
	c.Push(m, tokens, ch)

	for response := range ch {
		fmt.Println("Received response " + response)
	}
}
