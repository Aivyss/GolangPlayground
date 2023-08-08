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

	doPR(p, s, w)

	// before override
	fmt.Println(s.Person.Pr())
}

func doPR(humans ...structures.Human) {
	for _, v := range humans {
		fmt.Println(v.Pr())
	}
}
