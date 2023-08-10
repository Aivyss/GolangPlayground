package main

import (
	"fmt"
	"sort"
)

func main() {
	// integer sort
	nums := []int{6, 34, 9, 123, 7, 1}
	sort.Ints(nums)
	fmt.Println(nums)

	// string sort
	strs := []string{"relationship", "apply", "people", "golang", "management", "apple", "昼間", "ポテト", "新宿", "か", "ぬ", "あ"}
	sort.Strings(strs)
	fmt.Println(strs)
}
