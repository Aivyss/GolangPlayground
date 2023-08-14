package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	r := strings.NewReader("Han Lee")

	_, _ = io.Copy(f, r)
	fmt.Println("about to exit")

	_ = f.Close() // you can also use defer f.Close to above. But I positioned this here due to warnings
}
