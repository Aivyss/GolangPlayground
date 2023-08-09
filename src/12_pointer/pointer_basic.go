package main

import (
	"com.playground/structures"
	"fmt"
)

func main() {
	fmt.Println("=====[primitive type]=====")
	x := 64

	fmt.Printf("value = %v\n", x)
	fmt.Printf("pointer value = %v\n", &x)

	fmt.Printf("type = %T\n", x)
	fmt.Printf("pointer type = %T\n", &x)

	pointerX := &x
	fmt.Println("*pointerX = ", *pointerX)
	*pointerX = -1
	fmt.Printf("value = %v\n", x)

	y := 25
	z := 26
	fmt.Println("y, z = ", y, z)
	*&y, *&z = z, y
	fmt.Println("y, z = ", y, z)

	fmt.Println("=====[composite type]=====")
	var p structures.Person

	fmt.Printf("value = %v\n", p)
	fmt.Printf("pointer value = %v\n", &p)

	fmt.Printf("type = %T\n", p)
	fmt.Printf("pointer type = %T\n", &p)

	pointerP := &p
	fmt.Println(*pointerP)
}
