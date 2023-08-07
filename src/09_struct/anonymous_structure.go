package main

import "fmt"

func main() {
	anonym := struct {
		first string
		last  string
		age   int
	}{
		first: "Han",
		last:  "Lee",
		age:   30,
	}

	fmt.Printf("type : %T\n", anonym)
	fmt.Println(anonym)
}
