package main

import (
	"fmt"
)

func main() {
	sliceBasic()
	rangeExpression()
	slicing()
	appendValue()
}

func sliceBasic() {
	fmt.Println("=====Slice Basics=====")
	slice := []int{1, 2, 3, 4}

	// set
	slice[3] = 99

	// get
	fmt.Println(slice[3])

	// size
	fmt.Println(len(slice))

	// all
	fmt.Println(slice)
}

func rangeExpression() {
	fmt.Println("=====Range Expression=====")
	slice := []int{10, 20, 30, 40}

	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func slicing() {
	fmt.Println("=====Slicing=====")
	slice := []int{0, 1, 2, 3, 4}

	fmt.Printf("slice[:] = %v\n", slice[:])
	fmt.Printf("slice[1:] = %v\n", slice[1:])
	fmt.Printf("slice[0:3] = %v\n", slice[0:3])
}

func appendValue() {
	fmt.Println("=====Append values=====")
	x0 := []int{0, 1, 2, 3, 4}
	x1 := []int{5, 6, 7, 8, 9}
	num := 10

	// concat slices
	x2 := append(x0, x1...)

	// append value
	x3 := append(x2, num)

	// result
	fmt.Printf("x0 = %v\n", x0)
	fmt.Printf("x1 = %v\n", x1)
	fmt.Printf("x2 = %v\n", x2)
	fmt.Printf("x3 = %v\n", x3)

	// remove
	fmt.Println("=====remove=====")
	fmt.Println("remove0:", remove(x0, 0))
	fmt.Println("remove1:", remove(x0, 1))
	fmt.Println("remove2:", remove(x0, 2))
	fmt.Println("remove3:", remove(x0, 3))
	fmt.Println("remove4:", remove(x0, 4))
	fmt.Printf("x0 = %v\n", x0)
}

func remove[T any](slice []T, index int) []T {
	var result []T

	if len(slice)-1 == index {
		result = append([]T{}, slice[:index]...)
	} else {
		left := append([]T{}, slice[:index]...)
		right := append([]T{}, slice[index+1:]...)
		result = append(left, right...)
	}

	return result
}
