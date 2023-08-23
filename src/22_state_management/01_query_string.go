package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	router.GET("/test1", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		// get from request
		value := req.FormValue("q")

		// create response
		_, _ = io.WriteString(w, value)
		header := w.Header()
		header.Set("Content-Type", "text/plain")
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}
