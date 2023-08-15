package main

import (
	pError "com.playground/error"
	"fmt"
)

func main() {
	foo()
}

func foo() {
	fmt.Println("=====[run foo]=====")
	bar()
}

func bar() {
	fmt.Println("=====[run bar]=====")
	err := pError.NewPlayGroundError("bar error")

	fmt.Println("err.Msg =", err.Msg)
	fmt.Println("err.Error() =", err.Error())
	fmt.Println("err.StackMsg =", err.StackMsg)
}
