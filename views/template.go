package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Parse(filepath string) (Template, error) {
	htmlTpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("Parsing template: %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
