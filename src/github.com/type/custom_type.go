package main

import "fmt"

type hotdog int

var a = 48
var b hotdog = 24

func main() {
	// check type
	fmt.Printf("type: %T, value: %v\n", a, a)
	fmt.Printf("type: %T, value: %v\n", b, b)

	// b = a CAN NOT cast int to hotdog
	// a = b CAN NOT cast hotdog to int

	// type conversion
	c := int(b)
	fmt.Printf("type: %T, value: %v\n", c, c)

	d := hotdog(a)
	fmt.Printf("type: %T, value: %v\n", d, d)
}
