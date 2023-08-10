package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
I make race condition with shared memory (counter value) for test.
YOU DO NOT CODE LIKE THIS IN YOUR REAL PROJECTS.
*/
func main() {
	fmt.Println("CPUs =", runtime.NumCPU())
	fmt.Println("Goroutines =", runtime.NumGoroutine())

	counter := 0
	const gs = 1000 // count of goroutine
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched() // yield CPU resources to another goroutine.
			v += 1
			counter = v
			wg.Done()
		}()

		fmt.Println("Goroutines =", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("counter =", counter)
}
