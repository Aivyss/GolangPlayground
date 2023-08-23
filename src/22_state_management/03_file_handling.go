package main

import (
	"com.playground/18_template/util"
	"com.playground/21_http_server/templateLoad"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// define variables
	router := httprouter.New()
	tpl := templateLoad.LoadTemplate()

	// routing
	router.GET(
		"/favicon.ico",
		func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
			http.NotFoundHandler().ServeHTTP(w, req)
		})

	router.GET("/", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		_ = tpl.ExecuteTemplate(w, "fileFormData.gohtml", nil)
	})

	router.POST("/upload", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		// get file from form data
		file, header, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// file information
		fmt.Println("file =", file)
		fmt.Println("header.Size(byte) =", header.Size)
		fmt.Println("header.Header =", header.Header)
		fmt.Println("header.Filename =", header.Filename)
		fmt.Println("err =", err)

		// read the file
		bs, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s := string(bs)

		// save the file to server
		dst, err := os.Create(filepath.Join(util.GetGoPath(), "file", header.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// render view
		_ = tpl.ExecuteTemplate(w, "fileFormDataRead.gohtml", s)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}
}
