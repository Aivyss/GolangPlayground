package main

import "fmt"

/*
slice structure
*/
type item struct {
	name  string
	price int
}
type betterCart []item
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
	fmt.Printf("%T\n", crt[0])

	bc := betterCart{
		{
			name:  "itemA",
			price: 1000,
		},
		{
			name:  "itemB",
			price: 2000,
		},
	}
	fmt.Println(bc)
	fmt.Printf("%T\n", bc[0])
}
