package main

import "fmt"

func main() {
	resultByte, err := fmt.Println("Hello world", "Yeah")
	fmt.Println(resultByte, err)
}
