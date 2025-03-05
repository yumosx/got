package gen

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed model.tmpl
var tmpl string

type Field struct {
	Name string
	Type string
}

type TemplateData struct {
	Pkg    string
	Model  string
	Fields []Field
}

func Emit(path string, data TemplateData) error {
	t := template.Must(template.New("struct").Parse(tmpl))
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = t.Execute(file, data)
	if err != nil {
		return err
	}

	return err
}
