package main

import "fmt"

func main() {
	num := make(chan int)

	go func() {
		num <- 42 // The channel is blocked in this goroutine
	}()

	fmt.Println(<-num)
}
