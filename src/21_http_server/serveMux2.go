package main

import (
	"com.playground/21_http_server/templateLoad"
	"net/http"
)

var tpl = templateLoad.LoadTemplate()

func thirdHandlerFunc(writer http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(writer, "tpl3.gohtml", nil)
}

func main() {

	var firstHandlerFunc http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		tpl.ExecuteTemplate(writer, "tpl1.gohtml", nil)
	}
	var secondHandlerFunc http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		tpl.ExecuteTemplate(writer, "tpl2.gohtml", nil)
	}

	http.Handle("/first", firstHandlerFunc)
	http.HandleFunc("/second", secondHandlerFunc)
	http.Handle("/third", http.HandlerFunc(thirdHandlerFunc))

	// if you give nil to ListenAndServe, DefaultServeMux is used.
	http.ListenAndServe(":8080", nil)
}
