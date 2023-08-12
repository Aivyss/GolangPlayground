package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs =", runtime.NumCPU())
	fmt.Println("Goroutines =", runtime.NumGoroutine())

	counter := 0
	const gs = 1000 // count of goroutine
	var wg sync.WaitGroup
	var mutx sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mutx.Lock()
			v := counter
			runtime.Gosched() // yield CPU resources to another goroutine.
			v += 1
			counter = v
			mutx.Unlock()
			wg.Done()
		}()

		fmt.Println("Goroutines =", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("counter =", counter)
}
