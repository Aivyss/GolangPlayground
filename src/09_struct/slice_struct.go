package main

import "fmt"

/*
slice structure
*/
type cart []struct {
	name  string
	price int
}

func main() {
	crt := cart{
		{
			name:  "itemA",
			price: 1000,
		},
		{
			name:  "itemB",
			price: 2000,
		},
	}

	fmt.Println(crt)
}
