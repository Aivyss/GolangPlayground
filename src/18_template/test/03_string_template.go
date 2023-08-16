package main

import (
	"com.playground/file"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0])
	name := os.Args[1]
	fileName := os.Args[2]
	f, err := file.GetFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	tpl := `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>
					Hello World!
				</title>
			</head>
			<body>
				<h1>` + name + `</h1>
			</body>
		</html>
	`

	io.Copy(f, strings.NewReader(tpl))
}
