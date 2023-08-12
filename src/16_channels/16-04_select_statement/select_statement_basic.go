package main

import "fmt"

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	// send
	go send(odd, even, quit)

	// receive
	receive(odd, even, quit)

	// finish
	fmt.Println("finished")
}

func send(odd, even, quit chan<- int) {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(odd)
	close(even)
	quit <- 0
}

func receive(odd, even, quit <-chan int) {
	for {
		select {
		case v, ok := <-odd:
			if ok {
				fmt.Println("print odd value from the odd channel:", v, ok)
			}
		case v, ok := <-even:
			if ok {
				fmt.Println("print even value from the even channel:", v, ok)
			}
		case v, ok := <-quit:
			if ok {
				fmt.Println("the receive function is finished.", v, ok)
				return
			}
		}
	}
}
