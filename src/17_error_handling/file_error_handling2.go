package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

	_ = f.Close()
}
