package main

import (
	"fmt"
	"sync"
)

var singleton *singletonStruct
var once sync.Once

type singletonStruct struct {
	value string
}

func getValue(value string) *singletonStruct {
	once.Do(func() { singleton = &singletonStruct{value} })

	return singleton
}

func main() {
	value1 := getValue("first calling")
	value2 := getValue("second calling")

	fmt.Println("value1 =", *value1)
	fmt.Println("value2 =", *value2)
}
