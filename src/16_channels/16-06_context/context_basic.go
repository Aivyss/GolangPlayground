package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctxTest()
	withCancelTest()
	cancelExample1()
	cancelExample2()
}

func ctxTest() {
	fmt.Println("=====[ctxTest]=====")
	ctx := context.Background()

	fmt.Println("context =", ctx)
	fmt.Println("context err =", ctx.Err())
	fmt.Printf("context type =%T\n", ctx)
}

func withCancelTest() {
	fmt.Println("=====[withCancelTest]=====")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context =", ctx)
	fmt.Println("context err =", ctx.Err())
	fmt.Printf("context type =%T\n", ctx)
	fmt.Println("cancel =", cancel)
	fmt.Printf("cancel type =%T\n", cancel)

	cancel()
	fmt.Println("---after cancel--- ")
	fmt.Println("context =", ctx)
	fmt.Println("context err =", ctx.Err())
	fmt.Printf("context type =%T\n", ctx)
	fmt.Println("cancel =", cancel)
	fmt.Printf("cancel type =%T\n", cancel)
}

func cancelExample1() {
	fmt.Println("=====[cancelExample1]=====")
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error =", ctx.Err())
	fmt.Println("numGoroutine =", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n += 1
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working =", n)
			}

		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error =", ctx.Err())
	fmt.Println("numGoroutine =", runtime.NumGoroutine())

	fmt.Println("---about to cancel context---")
	cancel()
	fmt.Println("---cancelled context---")
	time.Sleep(time.Second * 2)
	fmt.Println("error =", ctx.Err())
	fmt.Println("numGoroutine =", runtime.NumGoroutine())
}

func cancelExample2() {
	fmt.Println("=====[cancelExample2]=====")
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1

		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n += 1
				}
			}
		}()

		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range gen(ctx) {
		if v == 5 {
			break
		} else {
			fmt.Println("Working =", v)
		}
	}

	fmt.Println("about to exit")
}
