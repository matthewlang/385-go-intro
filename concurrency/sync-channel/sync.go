package main

import (
	"fmt"
)

var channel = make(chan string)

func main() {

	go func() {
		channel <- "hello, world\n"
	}()

	a := <-channel
	fmt.Printf("%s\n", a)
}
