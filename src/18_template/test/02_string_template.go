package main

import (
	"com.playground/snippet/file"
	"fmt"
	"io"
	"strings"
)

func main() {
	f, err := file.GetSimpleFile("index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	name := "Han Lee"

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
