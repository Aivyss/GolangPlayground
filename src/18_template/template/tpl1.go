package main

import (
	"com.playground/18_template/util"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	testOneTemplate()
	testMultipleTemplates()
	testParseGlob()
}

func testOneTemplate() {
	fmt.Println("=====[testOneTemplate]=====")
	templatePath := filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "tpl1.gohtml")

	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func testMultipleTemplates() {
	fmt.Println("=====[testMultipleTemplates]=====")
	templatePaths := []string{
		filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "tpl1.gohtml"),
		filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "tpl2.gohtml"),
	}

	tpl, err := template.ParseFiles(templatePaths...)
	if err != nil {
		log.Fatalln(err)
	}

	for _, path := range templatePaths {
		splits := strings.Split(path, "/")
		templateFileName := splits[len(splits)-1]

		err = tpl.ExecuteTemplate(os.Stdout, templateFileName, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func testParseGlob() {
	fmt.Println("=====[testParseGlob]=====")
	templateFileName := []string{
		"tpl1.gohtml",
		"tpl2.gohtml",
	}
	tpl, err := template.ParseGlob(filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "*.gohtml"))
	if err != nil {
		log.Fatalln(err)
	}

	for _, fileName := range templateFileName {
		err = tpl.ExecuteTemplate(os.Stdout, fileName, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
