package main

import (
	"com.playground/18_template/util"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob(filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "*")))
}

func main() {
	templateFileName := []string{
		"tpl1.gohtml",
		"tpl2.gohtml",
	}

	for _, fileName := range templateFileName {
		err := tpl.ExecuteTemplate(os.Stdout, fileName, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
