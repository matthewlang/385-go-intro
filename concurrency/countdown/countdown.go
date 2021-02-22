package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan bool)
	lunch := make(chan bool)

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- true
	}()

	go func() {
		tick := time.Tick(1 * time.Second)
		for countdown := 5; countdown > 0; countdown-- {
			fmt.Printf("%d...", countdown)
			<-tick
		}
		lunch <- true
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")

	select {
	case <-abort:
		fmt.Printf("Launch aborted!\n")
	case <-lunch:
		fmt.Printf("Launch!\n")
	}
}
