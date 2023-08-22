package main

import (
	"com.playground/21_http_server/templateLoad"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type handler2 int

func (h handler2) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // parse query string and form body
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		Url           *url.URL
		Header        http.Header
		ContentLength int64
		Submissions1  map[string][]string
		Submissions2  map[string][]string
	}{
		Method:        req.Method,
		Url:           req.URL,
		Header:        req.Header,
		ContentLength: req.ContentLength,
		Submissions1:  req.Form,
		Submissions2:  req.PostForm,
	}

	responseHeader := rw.Header()
	// Set Method doesn't work after rw.WriteHeader
	responseHeader.Set("Custom-Key", "My Custom Value")
	responseHeader.Set("Content-Type", "text/html; charset=utf-8")
	// ExecuteTemplate runs rw.WriteHeader method
	_ = tpl2.ExecuteTemplate(rw, "methodAndHeader.gohtml", data)

}

var tpl2 *template.Template

func init() {
	tpl2 = templateLoad.LoadTemplate()
}

func main() {
	err := http.ListenAndServe(":8080", handler2(1))
	if err != nil {
		log.Fatalln(err)
	}
}
