package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.ParseRequestURI("httkoverflow.com/")
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(u)
}
