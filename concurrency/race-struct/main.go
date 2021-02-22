package main

import (
	"flag"
	"fmt"
	"sync"
)

type Tallier struct {
	tally  int64
	clicks int64
	mu     sync.Mutex // protects tally and clicks
}

func (t *Tallier) Add() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.tally += 1
	t.clicks += 1
}

func (t *Tallier) Sub() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.tally -= 1
	t.clicks += 1
}

func (t *Tallier) String() string {
	t.mu.Lock()
	defer t.mu.Unlock()

	return fmt.Sprintf("tally: %d, clicks: %d", t.tally, t.clicks)
}

func Clicker(t *Tallier, n int, done chan<- bool) {
	for i := 0; i < n; i++ {
		t.Add()
	}

	for i := 0; i < n; i++ {
		t.Sub()
	}
	done <- true
}

func main() {
	var n int
	var threads int
	flag.IntVar(&n, "n", 10, "how many iterations per thread")
	flag.IntVar(&threads, "t", 2, "how many threads to run")
	flag.Parse()

	tallier := &Tallier{}

	done := make(chan bool, threads)

	for i := 0; i < threads; i++ {
		go Clicker(tallier, n, done)
	}

	for i := 0; i < threads; i++ {
		<-done
	}

	fmt.Printf("State after %v iters by %v threads (%v total): %v\n", n*2, threads, 2*n*threads, tallier)
}
