package main

import (
	"bufio"
	"fmt"
	"strings"
)

/*
bufio.ScanRunes => 문자별로 나누는 함수
type: func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
*/
func main() {
	s := "test message\nsecond line message\n third line message\n"
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
