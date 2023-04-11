package controllers

import (
	"html/template"
	"net/http"

	"github.com/tombracho/ch-lenslocked/views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "How old are u?",
			Answer:   "I am 25 years old",
		},
		{
			Question: "What is your name",
			Answer:   "My name is Artem",
		},
		{
			Question: "Where is your office?",
			Answer:   "Our entire team is remote!",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
