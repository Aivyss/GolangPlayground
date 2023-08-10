package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("This function is invoked before main function.")
}

func main() {
	fmt.Println("OS =", runtime.GOOS)
	fmt.Println("Architecture =", runtime.GOARCH)
	fmt.Println("CPUs =", runtime.NumCPU())
	fmt.Println("Goroutine =", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPUs =", runtime.NumCPU())
	fmt.Println("Goroutine =", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("[foo]", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("[bar]", i)
	}
}
