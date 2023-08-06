package main

import "fmt"

func main() {
	dict := map[string]string{
		"졸리다":  "眠い",
		"배고프다": "お腹ペコペコ",
		"지우다":  "消す",
	}

	fmt.Println("result:", dict)
	fmt.Println("result:", dict["졸리다"])
	fmt.Println("result:", dict["배고프다"])
	fmt.Println("result:", dict["없는값"] == "") // zero value

	// check zero value (real value or zero value by ok variables)
	v1, ok1 := dict["졸리다"]
	v2, ok2 := dict["없는값"]

	fmt.Println(v1, ok1)
	fmt.Println(v2, ok2)

	// range
	for key, value := range dict {
		fmt.Println(key, ":", value)
	}

	// delete
	delete(dict, "지우다")
	fmt.Println("result:", dict)
}
