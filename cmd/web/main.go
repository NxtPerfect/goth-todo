package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"todo/cmd/web/templates"

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

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("cmd/web/templates/index.html"))
		tasks := map[string][]Task{
			"Tasks": {
				{Id: 0, Title: "Task1", Description: "Descr"},
				{Id: 1, Title: "Task2", Description: "Yes very"},
				{Id: 2, Title: "Task3", Description: "Nah i'm fein"},
			},
		}
		// first parameter is response writer, second is data
		// that we'll be able to access inside our html
		tmpl.Execute(w, tasks)
	}

	// h3 := func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("cmd/web/templates/login.html"))
	// 	tmpl.Execute(w, nil)
	// }

	h4 := func(w http.ResponseWriter, r *http.Request) {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		// Check in database if exists and password is correct
		fmt.Printf("%t, %t", username == "admin", password == "admin")
	}

  tmpl := templates.loginPage()
	http.HandleFunc("/", h1)
	http.Handle("/login", templ.Handler(tmpl))
	http.HandleFunc("/api/login", h4)
	// http.Handle("/api/tasks/0/edit", templ.Handler(login))

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), nil))
}
