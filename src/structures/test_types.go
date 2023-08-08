package structures

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func (p Person) IsAdult() bool {
	return p.Age >= 20
}

func (p Person) Pr() string {
	var adult string

	if p.IsAdult() {
		adult = "adult"
	} else {
		adult = "child"
	}

	return fmt.Sprintf("name :%v, age: %v (%v) and my job: None", p.FullName(), p.Age, adult)
}

func (p Person) FullName() string {
	return p.First + " " + p.Last
}

type Student struct {
	Person    // no field name => embedded structure
	StudentId string
}

func (s Student) Pr() string {
	var adult string

	if s.IsAdult() {
		adult = "adult"
	} else {
		adult = "child"
	}

	return fmt.Sprintf("name :%v, age: %v (%v) and my job: %v", s.FullName(), s.Age, adult, "Student")
}

type Worker struct {
	Person         Person
	JobDescription string
	Salary         int
	Id             string
}

func (s Worker) Pr() string {
	var adult string

	if s.Person.IsAdult() {
		adult = "adult"
	} else {
		adult = "child"
	}

	return fmt.Sprintf("name :%v, age: %v (%v) and my job: %v", s.Person.FullName(), s.Person.Age, adult, "Worker")
}

type Human interface {
	Pr() string
}
