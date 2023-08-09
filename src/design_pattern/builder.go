package main

import "fmt"

type stringBox struct {
	value string
}

/*
builder pattern
usually, builder method returns pointer type value because of saving the memory.
*/
func newStringBox(s string) *stringBox {
	return &stringBox{
		value: s,
	}
}

func main() {
	box := newStringBox("testBox")
	fmt.Println(*box)
}
