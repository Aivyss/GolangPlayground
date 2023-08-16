package main

import (
	"com.playground/18_template/util"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

var tpl4 *template.Template

func init() {
	tpl4 = template.Must(template.ParseGlob(filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "*")))
}

func main() {
	/*
		struct field of data must be public.
	*/
	err := tpl4.ExecuteTemplate(os.Stdout, "tpl4.gohtml", struct {
		Title string
		Name  string
	}{
		Title: "New title",
		Name:  "tpl4, Ayano",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("=====about to exit=====")
}
