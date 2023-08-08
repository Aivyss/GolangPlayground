package main

import "fmt"

func main() {
	defer foo()
	defer sum(1, 2)
	bar()
}

func foo() {
	fmt.Println("foo invoked")
}

func bar() {
	fmt.Println("bar invoked")
}

func sum(a, b int) int {
	fmt.Println("sum invoked")
	return a + b
}
