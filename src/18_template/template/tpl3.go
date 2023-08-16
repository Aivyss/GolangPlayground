package main

import (
	"com.playground/18_template/util"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

var tpl3 *template.Template

func init() {
	tpl3 = template.Must(template.ParseGlob(filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "*")))
}

func main() {
	err := tpl3.ExecuteTemplate(os.Stdout, "tpl3.gohtml", "Han Lee")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("=====about to exit=====")
}
