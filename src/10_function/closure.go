package main

import "fmt"

var packageScope = 1

func main() {
	functionScope := 2
	shadow := "before"

	// block scope & shadow
	{
		shadow := "after"
		fmt.Println("block scope")
		fmt.Println("functionScope =", functionScope)
		blockScope := 3
		functionScope += 1
		fmt.Println("functionScope =", functionScope)
		fmt.Println("blockScope =", blockScope)
		fmt.Println("shadow =", shadow)
	}
	fmt.Println("shadow =", shadow)
	// fmt.Println("blockScope =", blockScope) --> you can't

	fmt.Println("packageScope =", packageScope)
	foo2()
	fmt.Println("packageScope =", packageScope)
	fmt.Println("functionScope =", functionScope)

	// closure
	inc := incrementor()
	fmt.Println("inc() = ", inc())
	fmt.Println("inc() = ", inc())
}

func foo2() {
	packageScope += 1
}

func incrementor() func() int {
	x := 0

	return func() int {
		x += 1

		return x
	}
}
