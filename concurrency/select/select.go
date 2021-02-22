package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doWork(id string, howHard time.Duration, result chan<- int) {
	fmt.Printf("%v starting...\n", id)

	time.Sleep(time.Duration(float64(howHard) * (0.5 + rand.Float64())))
	val := rand.Int()

	fmt.Printf("%v done...\n", id)

	result <- val
}

func main() {
	fmt.Println("Starting some work...")

	a := make(chan int)
	b := make(chan int)

	go doWork("a", 1*time.Second, a)
	go doWork("b", 3*time.Second, b)

	ct := 0
	for {
		select {
		case ra := <-a:
			fmt.Printf("a done: %v\n", ra)
			ct += 1
		case rb := <-b:
			fmt.Printf("b done: %v\n", rb)
			ct += 1
		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Println("waiting...")
			if ct >= 2 {
				return
			}
		}
	}

}
