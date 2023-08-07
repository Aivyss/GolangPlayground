package main

import (
	"fmt"
	"strconv"
)

var z = 70

func main() {
	// identical
	var x int
	fmt.Println(x)

	var y = 0
	fmt.Println(y)

	// global variables
	fmt.Println(z)
	foo()
}

func foo() {
	fmt.Println("again: " + strconv.Itoa(z))
	fmt.Println("again: " + fmt.Sprintln(z))
}
