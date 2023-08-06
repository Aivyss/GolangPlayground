package main

import "fmt"

func main() {
	slice1 := []int{0, 1, 2, 3, 4, 5}
	slice2 := append([]int{}, slice1...) // deep copy

	fmt.Printf("slice1[:3] = %v\n", slice1[:3])
	fmt.Printf("slice1[3+1:] = %v\n", slice1[3+1:])
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("append(slice1[:3], slice1[3+1:]...) = %v\n", append(slice1[:3], slice1[3+1:]...))
	fmt.Printf("slice1 = %v\n", slice1) // changed element.
	fmt.Printf("slice2 = %v\n", slice2)

	/*
		slice1[:3] is a partial array of slice1.
		So, you append some elements by append method based on slice1[:3] then slice1 is modified too.
		But slice2 was created based on new slice([]int{}, new Array), so it works like deep copy.
	*/
}
