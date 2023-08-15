package main

import (
	log "com.playground/logger"
	"fmt"
)

/*
- LIFO (last in first out)
- also run when panic
*/
func main() {
	defer1()
	defer2()
	fmt.Println("deferred return =", defer4())
	defer3()
}

/*
evaluate value
*/
func defer1() {
	i := 0
	defer fmt.Println(i)
	i += 1

	fmt.Println("=====about to exit(defer1)=====")
}

/*
for statement
*/
func defer2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("defer: ", i)
	}

	fmt.Println("=====about to exit(defer2)=====")
}

/*
panic status
*/
func defer3() {
	fmt.Println("=====[defer3]=====")
	logger, logFile := log.GetLogger()
	/*
		defers do not care panic status.
	*/
	defer fmt.Println("defer 0")
	defer logFile.Close()
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	logger.Panicln("panic error happened")
}

/*
a deferred function increments the return value after  the surrounding function returns.
*/
func defer4() (i int) {
	defer func() { i += 1 }() // i == 1, 1 += 1 => 2

	return 1
}
