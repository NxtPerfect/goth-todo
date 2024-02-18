package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Task struct {
	Id          int32
	Title       string
	Description string
}

// date_created string
// date_modified string
// date_due string

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

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		htmlStr := fmt.Sprintf("<li><p>%d. %s</p><span>%s<span>", 3, title, description)
		tmpl, _ := template.New("T").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-task", h2)
	// http.HandleFunc("/api/tasks/0/edit", h3)

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), nil))
}
