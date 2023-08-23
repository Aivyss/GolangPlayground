package main

import (
	"com.playground/21_http_server/templateLoad"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = templateLoad.LoadTemplate()
}

var firstHandle httprouter.Handle = func(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := tpl.ExecuteTemplate(writer, "tpl1.gohtml", nil)

	if err != nil {
		notFoundPageHandlerFunc(writer, req)
	}
}
var secondHandle httprouter.Handle = func(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := tpl.ExecuteTemplate(writer, "tpl2.gohtml", nil)

	if err != nil {
		notFoundPageHandlerFunc(writer, req)
	}
}

var notFoundPageHandlerFunc = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	_ = tpl.ExecuteTemplate(w, "notFound.gohtml", nil)
})

func main() {
	router := httprouter.New()
	router.GET("/first", firstHandle)
	router.GET("/second", secondHandle)
	router.GET("/third/:name", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		name := params.ByName("name")
		query, err := url.ParseQuery(req.URL.RawQuery) // you get query strings not path variables.
		if err != nil {
			log.Println(err)
		}
		fmt.Println("qs =", query.Get("qs"))

		err = tpl.ExecuteTemplate(w, "tpl3.gohtml", name)
		if err != nil {
			notFoundPageHandlerFunc(w, req)
		}
	})

	// 404 settings
	router.NotFound = notFoundPageHandlerFunc

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}
