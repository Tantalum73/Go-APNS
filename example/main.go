package main

import (
	"fmt"

	"github.com/tantalum73/Go-APNS"
)

func main() {
	m := goapns.NewMessage().Badge(42).Title("Title").Body("body")
	c, err := goapns.NewConnection("pathname", "key")
	if err != nil {

	} else {
		c.Development()
	}
	fmt.Println(m)
	fmt.Println(c)
}
