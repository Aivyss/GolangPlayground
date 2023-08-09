package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human1 interface {
	pr1() string
}

type human2 interface {
	pr2() string
}

type human3 interface {
	pr1() string
	pr2() string
}

func (p person) pr1() string {
	return fmt.Sprintf("name :%v, age: %v (%v) and my job: None", p.first+" "+p.last, p.age, p.age >= 20)
}

func (p *person) pr2() string {
	return fmt.Sprintf("name :%v, age: %v (%v) and my job: None", p.first+" "+p.last, p.age, p.age >= 20)
}

func doPr1(h human1) string {
	return h.pr1()
}

func doPr2(h human2) string {
	return h.pr2()
}

func doPr3(h human3) (string, string) {
	return h.pr1(), h.pr2()
}

func main() {
	p := person{
		first: "Han",
		last:  "Lee",
		age:   30,
	}

	/*
		doPr1 : (r Receiver) => value, pointer
		doPr2 : (r *Receiver) => pointer only
	*/
	fmt.Println("doPr1(p) =", doPr1(p))
	fmt.Println("doPr1(&p) =", doPr1(&p))
	//fmt.Println("doPr2(p) =", doPr2(p)) --> compile error
	fmt.Println("doPr2(&p) =", doPr2(&p))

	/*
		if an interface method has more than one of pointer receiver,
		you just have to use pointer not value.
	*/
	v1, v2 := doPr3(&p)
	fmt.Println("v1 =", v1)
	fmt.Println("v2 =", v2)
}
