package controllers

import (
	"fmt"
	"net/http"

	"github.com/tombracho/ch-lenslocked/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}

	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	username := r.FormValue("username")
	user, err := u.UserService.Create(email, password, username)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "User created: %+v", user)
}

func (u Users) SignInHandler(w http.ResponseWriter, r *http.Request) {
	u.Templates.SignIn.Execute(w, nil)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	username := r.FormValue("username")
	user, err := u.UserService.Authenticate(username, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "User authenticated: %+v", user)
}
