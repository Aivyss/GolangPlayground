package main

import "fmt"

type sendChannel chan<- int
type receiveChannel <-chan int

func main() {
	num := make(chan int, 2)

	// send
	go foo(num, 42)

	// receive
	bar(num) // sync.WaitGroup이 없으므로 goroutine이 아닌 블로킹이 필요 (블로킹이 없을 시 메인 고루틴이 끝나 다른 고루틴이 실행되지 못함).

	fmt.Println("about to exit")

}

/*
	Channels are reference type like Slice. (DO NOT NEED POINTER)
*/
// send
func foo(c sendChannel, num int) {
	c <- num
}

// receive
func bar(c receiveChannel) {
	fmt.Println(<-c)
}
