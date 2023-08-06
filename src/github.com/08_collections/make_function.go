package main

import "fmt"

func main() {
	fmt.Println("=====Initialization=====")
	slice := make([]int, 10, 12)
	fmt.Println("slice =", slice)
	fmt.Println("length =", len(slice))
	fmt.Println("capacity =", cap(slice)) // 12

	// test append
	fmt.Println("=====append=====")
	slice = append(slice, 11)
	slice = append(slice, 12)
	slice = append(slice, 13)
	fmt.Println("slice =", slice)
	fmt.Println("length =", len(slice))
	fmt.Println("capacity =", cap(slice)) // 24 = 12 * 2

	/*
		length: slice length
		capacity: length unit of an array underlying slice (slice has an underlying slice with n * capacity size)
	*/
}
