package main

import "fmt"

func main() {
	arrayBasics()
}

func arrayBasics() {
	var arr [5]int

	// set
	arr[3] = 22

	// get
	fmt.Println(arr[3])

	// size
	fmt.Println(len(arr))

	// all
	fmt.Println(arr)
}
