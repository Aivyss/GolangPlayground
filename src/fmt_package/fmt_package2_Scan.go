package main

import "fmt"

func main() {
	var strings []string
	size := 0
	cnt := 0

	fmt.Print("input size :")
	_, err := fmt.Scan(&size)

	if err != nil {
		fmt.Println("about to exit due to error: ", err.Error())
		return
	}

	for cnt < size {
		cnt += 1
		var str string
		fmt.Print("input your value :")
		_, err = fmt.Scan(&str)

		if err != nil {
			fmt.Println(err.Error())
			cnt -= 1
		} else {
			strings = append(strings, str)
		}
	}

	fmt.Println(strings)
}
