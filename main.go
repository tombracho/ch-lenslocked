package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/tombracho/ch-lenslocked/controllers"
	"github.com/tombracho/ch-lenslocked/migrations"
	"github.com/tombracho/ch-lenslocked/models"
	"github.com/tombracho/ch-lenslocked/templates"
	"github.com/tombracho/ch-lenslocked/views"
)

func main() {
	//setup a database connection
	cfg := models.DefaultPostgresConfig()
	fmt.Println(cfg.String())
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//apply migrations
	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	//setup our model services
	userServices := models.UserService{
		DB: db,
	}
	sessionServices := models.SessionService{
		DB: db,
	}

	//setup middlewares
	userMw := controllers.UserMiddleware{
		SessionService: &sessionServices,
	}

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		//TODO: Fix this before deploying
		csrf.Secure(false),
	)

	//setup controllers
	usersC := controllers.Users{
		UserService:    &userServices,
		SessionService: &sessionServices,
	}

	//Create and setup router
	r := chi.NewRouter()

	r.Use(csrfMw)
	r.Use(userMw.SetUser)

	//setup routes
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)

	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signin", usersC.SignInHandler)
	r.Post("/signin", usersC.SignIn)

	usersC.Templates.CurrentUser = views.Must(views.ParseFS(templates.FS, "user.gohtml", "tailwind.gohtml"))
	r.Route("/users/me", func(r chi.Router) {
		r.Use(userMw.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})

	r.Post("/signout", usersC.SignOut)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
