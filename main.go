package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/tombracho/ch-lenslocked/controllers"
	"github.com/tombracho/ch-lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	//homeTpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	//homeTpl = views.Must(homeTpl, err)

	//contactTpl, err := views.Parse(filepath.Join("templates", "contact.gohtml"))
	//contactTpl = views.Must(contactTpl, err)

	//faqTpl, err := views.Parse(filepath.Join("templates", "faq.gohtml"))
	//faqTpl = views.Must(faqTpl, err)

	//r.Method(http.MethodGet, "/", controllers.Static{Template: homeTpl})
	//r.Method(http.MethodGet, "/contact", controllers.Static{Template: contactTpl})
	//r.Method(http.MethodGet, "/faq", controllers.Static{Template: faqTpl})

	r.Get("/", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))
	r.Get("/faq", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))
	r.Get("/login", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "login.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
