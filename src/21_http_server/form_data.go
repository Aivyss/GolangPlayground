package main

import (
	"com.playground/21_http_server/templateLoad"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type handler int

func (h handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // parse query string and form body
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(req.Method)

	//tpl1.ExecuteTemplate(rw, "formData.gohtml", req.PostForm) -> just form data
	_ = tpl1.ExecuteTemplate(rw, "formData.gohtml", req.Form) // -> form data + query string data
}

var tpl1 *template.Template

func init() {
	tpl1 = templateLoad.LoadTemplate()
}

func main() {
	err := http.ListenAndServe(":8080", handler(1))
	if err != nil {
		log.Fatalln(err)
	}
}
