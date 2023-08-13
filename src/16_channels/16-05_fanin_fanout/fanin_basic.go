package main

import (
	"fmt"
	"sync"
)

/*
fanin design pattern
*/
func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan string)

	// send channels
	go send(even, odd)
	go receive(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("about to exit")

	// fan in general pattern
	chan1 := make(chan int)
	chan2 := make(chan int)
	output := make(chan int)
	go send(chan1, chan2)
	go fanin(output, chan1, chan2)

	for v := range output {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func send(even, odd chan<- int) {
	defer close(even)
	defer close(odd)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
}

func receive(even, odd <-chan int, fanIn chan<- string) {
	var wg sync.WaitGroup
	defer close(fanIn)
	wg.Add(2)

	go func() {
		for v := range even {
			fanIn <- fmt.Sprintf("type = [even], \tvalue = %v", v)
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanIn <- fmt.Sprintf("type = [odd], \tvalue = %v", v)
		}
		wg.Done()
	}()

	wg.Wait()
}

func fanin[T any](output chan<- T, channels ...<-chan T) {
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(channel <-chan T) {
			for value := range channel {
				output <- value
			}

			wg.Done()
		}(channel)
	}

	wg.Wait()
	close(output)
}
