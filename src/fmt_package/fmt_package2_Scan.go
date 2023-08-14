package main

import "fmt"

func main() {
	var strings []string
	flag := true

	size := 0
	cnt := 0
	fmt.Print("input size :")
	_, err := fmt.Scan(&size)

	for flag && cnt < size {
		cnt += 1
		var str string
		fmt.Print("input your value :")
		_, err = fmt.Scan(&str)

		if err != nil {
			fmt.Println(err.Error())
			flag = false
		} else {
			strings = append(strings, str)
		}
	}

	fmt.Println(strings)
}
