package main

import (
	"com.playground/structures"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	p := structures.Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	id, _ := uuid.NewUUID()
	s := structures.Student{
		Person: structures.Person{
			First: "Minsu",
			Last:  "Kim",
			Age:   15,
		},
		StudentId: id.String(),
	}

	id, _ = uuid.NewUUID()
	w := structures.Worker{
		Person: structures.Person{
			First: "YongHee",
			Last:  "Park",
			Age:   27,
		},
		Salary:         3000,
		Id:             id.String(),
		JobDescription: "Back Office",
	}

	humans := []structures.Human{p, s, w}
	checkJobDescription(humans...)
}

func checkJobDescription(humans ...structures.Human) {
	for _, human := range humans {
		var result string

		switch human.(type) {
		case structures.Person:
			result = fmt.Sprintf("None, %v", human.(structures.Person).FullName())
		case structures.Student:
			result = fmt.Sprintf("Student, %v", human.(structures.Student).FullName())
		case structures.Worker:
			result = fmt.Sprintf("Worker, %v", human.(structures.Worker).Person.FullName())
		default:
			result = "unknown error"
		}

		fmt.Println(result)
	}
}
