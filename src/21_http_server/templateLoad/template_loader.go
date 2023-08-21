package templateLoad

import (
	"com.playground/18_template/util"
	"html/template"
	"path/filepath"
)

func LoadTemplate() *template.Template {
	return template.Must(template.ParseGlob(filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "*")))
}
