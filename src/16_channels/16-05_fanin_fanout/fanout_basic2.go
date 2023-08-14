package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type valueTime2 struct {
	value int
	time  time.Duration
}

func main() {
	channelSize := 3000
	c1 := make(chan int)
	c2 := make(chan valueTime2)
	go channelReservoir2(c1, channelSize)
	go fanOutIn2(c1, c2)

	for value := range c2 {
		fmt.Println(value)
	}
}

func channelReservoir2(ch chan<- int, size int) {
	for i := 0; i < size; i++ {
		ch <- i
	}
	close(ch)
}

func fanOutIn2(c1 <-chan int, output chan<- valueTime2) {
	var wg sync.WaitGroup
	for value := range c1 {
		wg.Add(1)

		go func(ch chan<- valueTime2, value int) {
			ch <- timeConsumingWork2(value)
			wg.Done()
		}(output, value)
	}

	wg.Wait()
	close(output)
}

func timeConsumingWork2(n int) valueTime2 {
	delayTime := time.Millisecond * time.Duration(rand.Intn(500))
	time.Sleep(delayTime)

	return valueTime2{
		value: n,
		time:  delayTime,
	}
}
