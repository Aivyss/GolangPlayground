package main

import "fmt"

func main() {
	// For Clause
	fmt.Println("=====[Basic Statement]=====")
	for i := 0; i < 5; i += 1 {
		fmt.Println(i)
	}

	// conditional statement
	fmt.Println("=====[Conditional Statement]=====")
	x := 0

	for x < 5 {
		fmt.Println(x)

		x += 1
	}

	x = 0
	for {
		if x >= 5 {
			break
		} else {
			fmt.Println(x)
		}

		x += 1
	}

	fmt.Println("=====[Example1]=====")
	var slice []string
	for i := 33; i <= 122; i += 1 {
		slice = append(slice, string(rune(i)))
	}

	fmt.Println(slice)
}
