package main

import "fmt"

func main() {
	num := make(chan int)

	// send
	go func(c chan<- int) {
		const cnt = 100
		for i := 0; i < cnt; i++ {
			c <- i
		}

		close(c)
	}(num)

	// receive (range)
	for value := range num {
		fmt.Println(value)
	}

	fmt.Println("about to exit")
}
