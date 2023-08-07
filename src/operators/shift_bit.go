package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Printf("%v\t\t\t\t%d\t\t%b\n", x, x, x)
	y := x << 1
	fmt.Printf("%v\t\t\t\t%d\t\t%b\n", y, y, y)

	fmt.Printf("%v\t\t\t\t%d\t\t%b\n", kb, kb, kb)
	fmt.Printf("%v\t\t\t\t%d\t\t%b\n", mb, mb, mb)
	fmt.Printf("%v\t\t\t\t%d\t\t%b\n", gb, gb, gb)

}
