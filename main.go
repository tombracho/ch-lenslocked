package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	method := r.Method
	fmt.Println("homeHandler executed with method", method, r.URL.Path)
	fmt.Fprint(w, "<h1>Hello, World!</h1>")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("loginHandler executed with method", method, r.URL.Path)
	fmt.Fprint(w, "<h1>LOGIN</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	method := r.Method
	fmt.Println("contactHandler executed with method", method, r.URL.Path)
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch, email me at <a href="mailto:artyomsonyx@gmail.com">`)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Status code 404: not fount")
	http.Error(w, "Page not found", http.StatusNotFound)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/":
		homeHandler(w, r)
	case "/login":
		loginHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

type Router struct{}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h2>FAQ</h2>")
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/":
		homeHandler(w, r)
	case "/login":
		loginHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
	//http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
