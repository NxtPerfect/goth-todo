package main

import (
	"log"
	"net/http"
	"os"
	"todo/app/web/api"
	"todo/app/web/templates"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

type Task struct {
	Id          int32
	Title       string
	Description string
	// date_created string
	// date_modified string
	// date_due string
}

func main() {
	godotenv.Load()

	home := templates.HomePage()
	login := templates.LoginPage()
	register := templates.RegisterPage()

	http.Handle("/", templ.Handler(home))
	http.Handle("/login", templ.Handler(login))
	http.Handle("/register", templ.Handler(register))
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/register", api.Register)

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), nil))
}
