package main

import (
	"fmt"

	"github.com/tantalum73/Go-APNS"
)

func main() {
	m := goapns.NewMessage().Badge(42).Title("Title").Body("body")
	c, _ := goapns.NewConnection("pathname", "key")
	fmt.Println(m)
	fmt.Println(c)
}
