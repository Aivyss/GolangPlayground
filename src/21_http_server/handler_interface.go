package main

import (
	"fmt"
	"log"
	"net/http"
)

type testType int

func (t testType) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Fprintf(rw, "Handler Test")
}
func main() {
	var handler http.Handler = testType(1)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalln(err)
	}
}
