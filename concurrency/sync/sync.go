package main

import (
	"fmt"
	"sync"
)

var a string

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		a = "hello, world\n"
		sum := 0
		for i := 0; i < 100_000; i++ {
			sum += i
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("%s\n", a)
}
