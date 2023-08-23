package main

import (
	"com.playground/21_http_server/templateLoad"
	"net/http"
)

func main() {
	template := templateLoad.LoadTemplate()
	mux := http.NewServeMux()

	mux.HandleFunc("/first", func(writer http.ResponseWriter, request *http.Request) {
		template.ExecuteTemplate(writer, "tpl1.gohtml", nil)
	})
	mux.HandleFunc("/second", func(writer http.ResponseWriter, request *http.Request) {
		template.ExecuteTemplate(writer, "tpl2.gohtml", nil)
	})

	http.ListenAndServe(":8080", mux)
}
