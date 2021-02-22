package main

import (
	"flag"
	"fmt"
	"sync"
)

var a int

func foo() {
	fmt.Println("I'm done!")
}

func printNum(num int) {
	fmt.Println(num)
}

func adder(n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	defer printNum(1)
	defer printNum(2)
	defer printNum(3)

	for i := 0; i < n; i++ {
		a += 1

		if a == 500 {
			return
		}
	}
}

func main() {
	var n int
	var threads int
	flag.IntVar(&n, "n", 10, "how many iterations per thread")
	flag.IntVar(&threads, "t", 2, "how many threads to run")
	flag.Parse()

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go adder(n, &wg, &mu)
	}

	wg.Wait()

	fmt.Printf("a = %d (%d iterations per %d threads)\n", a, n, threads)
}
