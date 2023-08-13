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
			case <-ch1:
				fmt.Println("receive ch1 message")
			case <-ch2:
				fmt.Println("receive ch2 message")
			}
		}
	}

	/*
		select_example_1의 동기성 처리 문제를 해결한 코드
	*/
	go first()
	go second()
	go third()

	time.Sleep(5 * time.Second)
	done <- true
}
