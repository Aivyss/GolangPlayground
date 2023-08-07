package main

import "fmt"

func main() {
	x := 25
	fmt.Printf("%T\n", x)
	fmt.Println(`first line
second line
third line`)

	// zero values
	var a int
	var y bool
	var z string
	var b float32
	fmt.Println(a)
	fmt.Println(y)
	fmt.Println(b)
	fmt.Println("-->" + z + "<--")
}
