package main

import "fmt"

func main() {
	// general
	x := 22
	fmt.Println("=========[General]=========")
	fmt.Printf("-->v = %v<--\n", x)
	fmt.Printf("-->#v = %#v<--\n", x)
	fmt.Printf("-->T = %T<--\n", x)
	//fmt.Printf("-->%%<--\n", x)

	// boolean
	fmt.Println("=========[Boolean]=========")
	fmt.Printf("-->v = %v<--\n", false)
	fmt.Printf("-->t = %t<--\n", false)

	// integer
	fmt.Println("=========[Integer]=========")
	fmt.Printf("-->b = %b<--\n", x)
	fmt.Printf("-->c = %c<--\n", x)
	fmt.Printf("-->d = %d<--\n", x)
	fmt.Printf("-->o = %o<--\n", x)
	fmt.Printf("-->q = %q<--\n", x)
	fmt.Printf("-->x = %x<--\n", x)
	fmt.Printf("-->X = %X<--\n", x)

	// sprint
	fmt.Println("=========[Sprint]=========")
	result := fmt.Sprintf("-->b = %b<--\n", x)
	fmt.Printf(result)
}
