package controllers

import (
	"net/http"
)

type SessionMiddleware struct {
	User *Users
}

func (sm *SessionMiddleware) CheckSessionToken(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := readCookie(r, CookieSession)
		if err != nil {
			f(w, r)
			return
		}
		http.Redirect(w, r, "/users/me", http.StatusFound)
	}
}
