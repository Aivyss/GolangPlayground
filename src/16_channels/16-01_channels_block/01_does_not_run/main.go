package main

import "fmt"

/*
Just 1 Goroutine isn't concurrent state.
So channels are blocked.
*/
func main() {
	num := make(chan int)

	num <- 42          // The channel is blocked.
	fmt.Println(<-num) // You try to use the blocked channel. (NOT TO RUN)
}
