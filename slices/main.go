package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr: %v, (len = %d, cap = %d)\n", arr, len(arr), cap(arr))

	var arr2 [5]int
	arr2 = arr
	arr2[0] = 42
	fmt.Printf("arr: %v, (len = %d, cap = %d)\n", arr, len(arr), cap(arr))
	fmt.Printf("arr2: %v, (len = %d, cap = %d)\n", arr2, len(arr2), cap(arr2))

	var slice []int
	slice = arr[:]

	slice[0] = 42
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("arr: %v\n", arr)

	a := make([]int, 5, 8)
	b := a
	fmt.Printf("a: %p, %[1]v len = %d, cap = %d\n", a, len(a), cap(a))
	fmt.Printf("b: %p, %[1]v len = %d, cap = %d\n", b, len(b), cap(b))

	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Printf("a: %p, %[1]v len = %d, cap = %d\n", a, len(a), cap(a))
	fmt.Printf("b: %p, %[1]v len = %d, cap = %d\n", b, len(b), cap(b))

	b = b[0:cap(b)]
	fmt.Printf("b: %p, %[1]v len = %d, cap = %d\n", b, len(b), cap(b))

	a = append(a, 4)
	fmt.Printf("a: %p, %[1]v len = %d, cap = %d\n", a, len(a), cap(a))
	fmt.Printf("b: %p, %[1]v len = %d, cap = %d\n", b, len(b), cap(b))

	b = a
	fmt.Printf("a: %p, %[1]v len = %d, cap = %d\n", a, len(a), cap(a))
	fmt.Printf("b: %p, %[1]v len = %d, cap = %d\n", b, len(b), cap(b))

}
