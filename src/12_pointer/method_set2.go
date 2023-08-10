package main

import "fmt"

type SomeInterface interface {
	SomeMethod1()
	SomeMethod2()
}

type SomeType struct {
	//...
}

func (r SomeType) SomeMethod1() {
	// ...
}

func (r *SomeType) SomeMethod2() {
	// ...
}

func main() {
	/*
		SomeType is not an implementation of SomeInterface. (it has SomeMethod1)
		&SomeType (Pointer) is an implementation of SomeInterface. (it has SomeMethod1 and 2)
	*/

	//slice := []SomeInterface{SomeType{}} --> compile error
	slice := []SomeInterface{&SomeType{}}
	fmt.Println(slice)
}
