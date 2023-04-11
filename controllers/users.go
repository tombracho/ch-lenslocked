package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form submission.", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "<p>The username: %s</p>", r.PostFormValue("username"))
	fmt.Fprintf(w, "<p>The email: %s</p>", r.FormValue("email"))
	fmt.Println(r.PostForm.Get("password"))
	fmt.Fprintf(w, "<p>The passowrd: %s</p>", r.PostForm.Get("passowrd"))
}
