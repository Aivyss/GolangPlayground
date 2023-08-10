package main

import (
	"com.playground/structures"
	"fmt"
	"sort"
)

// byAgeSort sort method 2
type byAgeSort []structures.Person

func (b byAgeSort) Len() int {
	return len(b)
}

func (b byAgeSort) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b byAgeSort) Less(i int, j int) bool {
	return b[i].Age < b[j].Age
}

func main() {
	people := []structures.Person{
		{
			First: "Han",
			Last:  "Lee",
			Age:   30,
		},
		{
			First: "James",
			Last:  "Bond",
			Age:   32,
		},
		{
			First: "YeongHee",
			Last:  "Lee",
			Age:   14,
		},
		{
			First: "CheolSu",
			Last:  "Kim",
			Age:   17,
		},
	}

	// prepare sort slice
	test1 := append([]structures.Person{}, people...)
	test2 := append([]structures.Person{}, people...)
	test3 := append([]structures.Person{}, people...)
	fmt.Println("people =", people)
	fmt.Println("test1 =", test1)
	fmt.Println("test2 =", test2)

	// sort method 1
	sort.Slice(test1, func(i, j int) bool {
		return test1[i].Age < test1[j].Age
	})

	// sort method 2
	sort.Sort(byAgeSort(test2))

	// string sort
	sort.Slice(test3, func(i, j int) bool {
		return test3[i].First < test3[j].First
	})

	// result
	fmt.Println("people =", people)
	fmt.Println("test1 =", test1)
	fmt.Println("test2 =", test2)
	fmt.Println("test3 =", test3)
}
