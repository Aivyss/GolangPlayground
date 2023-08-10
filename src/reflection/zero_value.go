package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := struct{ A int }{0}
	field := reflect.ValueOf(s).Field(0)

	// Use of IsZero() method
	fmt.Println(field.IsZero())

	s = struct{ A int }{1}
	field = reflect.ValueOf(s).Field(0)

	// Use of IsZero() method
	fmt.Println(field.IsZero())

	s = struct{ A int }{}
	field = reflect.ValueOf(s).Field(0)

	// Use of IsZero() method
	fmt.Println(field.IsZero())

	/*
		true
		false
		true
		-> golang do not check initialization... =_=
	*/
}
