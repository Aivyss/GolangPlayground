package main

import "fmt"

func main() {
	x := 64
	fmt.Println("main value = ", x)
	foo(x)
	fmt.Println("main value = ", x)
	bar(&x)
	fmt.Println("main value = ", x)
}

func foo(value int) {
	value = -1
	fmt.Println("foo value = ", value)
}

func bar(value *int) {
	*value = -2
	fmt.Println("bar *value = ", *value)
}
