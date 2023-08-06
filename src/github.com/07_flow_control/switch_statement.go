package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch x := "testString"; x {
	case "a":
		fmt.Println("case a")
	default:
		fmt.Println("default case", x)
	}

	for i := 0; i < 1000; i++ {
		number := rand.Float64()

		switch {
		case number > 0.7:
			fmt.Printf("x = %f is larger than 0.7\n", number)
		case number > 0.5:
			fmt.Printf("x = %f is larger than 0.5\n", number)
		case number > 0.3:
			fmt.Printf("x = %f is larger than 0.3\n", number)
		case number > 0.1:
			fmt.Printf("x = %f is larger than 0.1\n", number)
			fallthrough
		case false:
			fmt.Println("fallthrough test", number)
		default:
			fmt.Println("default case:", number)

		}
	}
}
