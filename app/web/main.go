package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"todo/app/web/api"
	"todo/app/web/templates"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

type Task struct {
	Id                 int32
	Title              string
	Description        string
	date_created       string
	date_last_modified string
	date_due           string
}

func readUsernameCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usernameCookie, err := r.Cookie("username")
		if err == nil {
			ctx := context.WithValue(r.Context(), "username", usernameCookie.Value)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	godotenv.Load()

	// username := r.Cookie("username")

	// home := templates.HomePage()
	login := templates.LoginPage()
	register := templates.RegisterPage()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username, err := r.Cookie("username")
		if err != nil {
			fmt.Printf("%s \n", err)
			templates.HomePage("").Render(r.Context(), w)
			return
		}
		templates.HomePage(username.Value).Render(r.Context(), w)
	})
	http.Handle("/login", templ.Handler(login))
	http.Handle("/register", templ.Handler(register))
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/register", api.Register)

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), nil))
}
