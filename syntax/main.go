package main

import (
	"errors"
	"fmt"
)

func fail() error {
	return errors.New("i always fail😿")
}

func myFunc() (x int, y int) {
	x = 1
	y = 2
	return
}

func myFunc2() (int, int) {
	var x int
	var y int

	x = 1
	y = 2

	return x, y
}

func myFunc3() (int, int) {
	return 1, 2
}

func fact(n int) (val int, err error) {

	if n < 0 {
		err = errors.New("factorial is not defined for negative numbers")
		return
	}

	val = 1
	for i := 1; i <= n; i++ {
		val = val * i
	}

	return
}

func main() {
	n := 4
	result, err := fact(n)
	if err != nil {
		fmt.Printf("oh no oh noooo oh no 😿")
	} else {
		fmt.Printf("fact(%v) = %v (error: %v)\n", n, result, err)
	}

	n = 10
	result, err = fact(n)
	if err != nil {
		fmt.Printf("oh no oh noooo oh no 😿")
	} else {
		fmt.Printf("fact(%v) = %v (error: %v)\n", n, result, err)
	}

	n = -5
	result, err = fact(n)
	if err != nil {
		fmt.Printf("oh no oh noooo oh no 😿")
	} else {
		fmt.Printf("fact(%v) = %v (error: %v)\n", n, result, err)
	}
}
