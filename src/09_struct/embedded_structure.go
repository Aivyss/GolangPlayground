package main

import (
	"com.playground/src/structures"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id, _ := uuid.NewUUID()
	p := structures.Person{
		First: "Han",
		Last:  "Lee",
		Age:   30,
	}
	std := structures.Student{
		Person:    p,
		StudentId: id.String(),
	}

	fmt.Println(std)
}
