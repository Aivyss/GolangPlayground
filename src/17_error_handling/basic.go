package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("test error")

	if err != nil {
		fmt.Println(err.Error())
	}
}
