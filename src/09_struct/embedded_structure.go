package main

import (
	"com.playground/src/structures"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	// embedded structure
	id, _ := uuid.NewUUID()
	p := structures.Person{
		First: "Han",
		Last:  "Lee",
		Age:   30,
	}
	std := structures.Student{
		Person:    p, // embedded structure
		StudentId: id.String(),
	}

	fmt.Println(std)
	fmt.Println(std.Person)
	fmt.Println(std.First, std.Last, std.Age, std.StudentId) // you can access fields of a embedded structure directly.

	// not embedded structure
	id, _ = uuid.NewUUID()
	worker := structures.Worker{
		Person: structures.Person{
			First: "Han",
			Last:  "Lee",
			Age:   30,
		},
		Id:             id.String(),
		Salary:         1000,
		JobDescription: "Application Engineer",
	}

	fmt.Println(worker)
	fmt.Println(worker.Id, worker.Salary, worker.JobDescription)
	fmt.Println(worker.Person)
}
