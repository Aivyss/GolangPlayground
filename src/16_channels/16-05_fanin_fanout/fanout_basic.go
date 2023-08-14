package main

import (
	"fmt"
	"math/rand"
	"time"
)

type valueTime struct {
	value int
	time  time.Duration
}

func main() {
	channelSize := 3000
	c1 := channelReservoir(channelSize)
	c2 := fanOutIn(c1)

	for i := 0; i < channelSize; i++ {
		fmt.Println(<-c2)
	}

}

func channelReservoir(size int) <-chan int {
	ch := make(chan int)

	for i := 0; i < size; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	return ch
}

func fanOutIn(c1 <-chan int) <-chan valueTime {
	output := make(chan valueTime)

	go func(ch chan<- valueTime, in <-chan int) {
		for {
			go func(ch chan<- valueTime, in <-chan int) {
				ch <- timeConsumingWork(<-in)
			}(output, c1)
		}
	}(output, c1)

	return output
}

func timeConsumingWork(n int) valueTime {
	delayTime := time.Millisecond * time.Duration(rand.Intn(500))
	time.Sleep(delayTime)

	return valueTime{
		value: n,
		time:  delayTime,
	}
}
