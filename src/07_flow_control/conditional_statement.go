package main

import (
	"fmt"
	"math/rand"
)

func main() {
	cnt := 0

	for cnt < 10 {
		if x := rand.Float64(); x > 0.5 {
			// you can't use y variable because out of scope
			fmt.Printf("[if block] x = %f\n", x)
		} else if y := rand.Float64(); x+y > 1 {
			fmt.Printf("[else if block] x = %f / y = %f\n", x, y)
		} else {
			fmt.Printf("[else block] x = %f / y = %f\n", x, y)
		}

		// you can't use x, y variable because out of scope
		cnt += 1
	}
}
