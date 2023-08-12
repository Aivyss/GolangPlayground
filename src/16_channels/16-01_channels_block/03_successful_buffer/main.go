package main

import "fmt"

func main() {
	num := make(chan int, 1) // create channel buffer

	num <- 42
	//num <- 42 // The channel is blocked.
	fmt.Println(<-num)
}
