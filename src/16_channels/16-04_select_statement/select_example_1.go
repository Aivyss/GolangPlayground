package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	done := make(chan bool)
	first := func() {
		for {
			time.Sleep(100 * time.Millisecond)
			ch1 <- true
		}
	}
	second := func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- true
		}
	}
	third := func() {
		for {
			select {
			case <-done:
				fmt.Println("exit gracefully")

				return
			default:
				<-ch1
				fmt.Println("receive ch1 message")
				<-ch2
				fmt.Println("receive ch2 message")
			}
		}
	}

	/*
		first goroutine의 100 밀리초의 고루틴은 의미가 없음. 왜냐하면 third goroutine의 default 구문에서 동기적으로 작동하는데
		ch2에서 500 밀리초가 소요되므로 병목현상이 발생.
	*/
	go first()
	go second()
	go third()

	time.Sleep(5 * time.Second)
	done <- true
	time.Sleep(5 * time.Second)
}
