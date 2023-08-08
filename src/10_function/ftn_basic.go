package main

import (
	"com.playground/structures"
	"fmt"
)

func main() {
	// multiple return values
	person, adult := getPerson("Han", "Lee", 30)
	fmt.Println(person, adult)

	// variadic parameters
	checkVariadicParams(1, 2, 3, 4, 5, 6)
	checkVariadicParams[int]() // len 0 cap 0

	ints := []int{0, 1, 2, 3, 4}
	checkNewOne(ints...)
	fmt.Println(ints)
}

func getPerson(first, last string, age int) (structures.Person, bool) {
	return structures.Person{
		First: first,
		Last:  last,
		Age:   age,
	}, age >= 20
}

func checkVariadicParams[T any](elems ...T) {
	fmt.Printf("type = %T\n", elems)
	fmt.Printf("value = %v\n", elems)
	fmt.Println("len =", len(elems))
	fmt.Println("cap =", cap(elems))
}

func checkNewOne(slice ...int) {
	fmt.Println(append(slice, 5))
}
