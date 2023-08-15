package main

import (
	log "com.playground/logger"
	"errors"
	"fmt"
	"time"
)

func main() {
	value, ok := panicTest1()
	fmt.Println(value, ok)
	fmt.Println("about to exit")

	test2, err := panicTest2()
	if err != nil {
		fmt.Println(err, test2 == 0)
	}
}

func panicTest1() (int, bool) {
	defer func() {

		if rec := recover(); rec != nil {
			fmt.Printf("type: %T, value: %v\n", rec, rec)
		}
	}() // recover
	logger, logFile := log.GetLogger()
	/*
		defers do not care panic status.
	*/
	defer logFile.Close()
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	/*
		All goroutines stopped by panic status.
	*/
	go func() {
		for {
			fmt.Println("Running another goroutine.")
			time.Sleep(time.Nanosecond * 1000)
		}
	}()

	logger.Panicln(errors.New("panic error happened"))

	/*
		dead code
	*/
	fmt.Println("after panic line")

	return 1, true
}

func panicTest2() (num int, err error) {
	defer func() {

		if rec := recover(); rec != nil {
			fmt.Printf("type: %T, value: %v\n", rec, rec)
			num = 0
			err = errors.New(rec.(string))
		}
	}() // recover
	logger, logFile := log.GetLogger()
	/*
		defers do not care panic status.
	*/
	defer logFile.Close()
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	/*
		All goroutines stopped by panic status.
	*/
	go func() {
		for {
			fmt.Println("Running another goroutine.")
			time.Sleep(time.Nanosecond * 1000)
		}
	}()

	logger.Panicln(errors.New("panic error happened"))

	/*
		dead code
	*/
	fmt.Println("after panic line")

	return 1, nil
}
