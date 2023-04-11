package controllers

import (
	"net/http"
)

//type Static struct {
//	Template Template
//}

//func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	static.Template.Execute(w, nil)
//}

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
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
