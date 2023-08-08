package main

import (
	"com.playground/structures"
	"fmt"
)

func main() {
	p := structures.Person{
		First: "Han",
		Last:  "Lee",
		Age:   30,
	}
	fmt.Println(p.IsAdult())
	fmt.Println(p.FullName())
}
