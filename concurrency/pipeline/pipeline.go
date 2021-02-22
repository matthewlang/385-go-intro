package main

import "fmt"

func generator(out chan<- int, n int) {
	for x := 0; x < n; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go generator(naturals, 20)
	go squarer(squares, naturals)

	for v := range squares {
		fmt.Printf("%d\n", v)
	}
}
