package main

import (
	"fmt"
)

func main() {
	s := "Hello world!"
	ss := "안녕하세요"
	bs := []byte(s)
	bss := []byte(ss)

	fmt.Println("======[conversion to byte array]=====")
	fmt.Printf("type: %T, value: %v\n", s, s)
	fmt.Printf("type: %T, value: %v\n", bs, bs)
	fmt.Printf("type: %T, value: %v\n", ss, ss)
	fmt.Printf("type: %T, value: %v\n", bss, bss)

	fmt.Println("======[str array, s[i]]=====")
	for i := 0; i < len(s); i++ {
		fmt.Printf("UTF-8: %#U, value: %v\n", s[i], s[i])
	}
	fmt.Println("======[str array, ss[i]]=====")
	for i := 0; i < len(ss); i++ {
		fmt.Printf("UTF-8: %#U, value: %v\n", ss[i], ss[i])
	}

	fmt.Println("======[str array, range s]=====")
	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println("======[str array, range ss]=====")
	for i, v := range ss {
		fmt.Println(i, v)
	}
}
