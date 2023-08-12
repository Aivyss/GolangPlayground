package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs =", runtime.NumCPU())
	fmt.Println("Goroutines =", runtime.NumGoroutine())

	var counter int64 = 0
	const gs = 1000 // count of goroutine
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() // yield CPU resources to another goroutine.
			fmt.Println("counter =", atomic.LoadInt64(&counter))
			wg.Done()
		}()

		fmt.Println("Goroutines =", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("====[after Wait]====")
	fmt.Println("counter =", counter)
}
