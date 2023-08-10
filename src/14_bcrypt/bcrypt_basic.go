package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

/*
you need to implement the package below
go get -u golang.org/x/crypto/...
*/
func main() {
	// generate bcrypt password
	pw := "testPw"
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)

	if err != nil {
		bs = nil
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	// compare bcrypt hashed password with plain password
	err = bcrypt.CompareHashAndPassword(bs, []byte("testPw"))

	if err != nil {
		fmt.Println("incorrect password")
	} else {
		fmt.Println("correct password")
	}

	err = bcrypt.CompareHashAndPassword(bs, []byte("wrong password"))

	if err != nil {
		fmt.Println("incorrect password")
	} else {
		fmt.Println("correct password")
	}
}
