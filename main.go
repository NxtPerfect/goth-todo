package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "strconv"
	// "time"
	"todo/app/web/api"
	"todo/app/web/database"
	"todo/app/web/templates"
	"todo/app/web/types"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	login := templates.LoginPage()
	register := templates.RegisterPage()
	tos := templates.TosPage()
	notFound := templates.NotFoundPage()
	showAddForm := false

  homePage := func(w http.ResponseWriter, r *http.Request) {
		username, err := r.Cookie("username")
		if err != nil || username == nil {
			fmt.Printf("Couldn't get username, assuming guest %s \n", err)
			templates.HomePage("", []types.Task{}).Render(r.Context(), w)
			return
		}

		auth_token, err := r.Cookie("auth_token")
		if err != nil || auth_token == nil {
			fmt.Printf("No auth token for user")
			templates.HomePage("", []types.Task{}).Render(r.Context(), w)
			return
		}
		tasks := database.GetTasks(username.Value, auth_token.Value, database.Connect())
		templates.HomePage(username.Value, tasks).Render(r.Context(), w)
	}

	http.HandleFunc("/", homePage)
	http.Handle("/login", templ.Handler(login))
	http.Handle("/register", templ.Handler(register))
	http.Handle("/tos", templ.Handler(tos))
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if !showAddForm {
			form := templates.AddForm()
			showAddForm = true
			form.Render(r.Context(), w)
		}
		return
	})
	http.HandleFunc("/tasks/edit", api.EditTaskForm)
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/register", api.Register)
	http.HandleFunc("/api/logout", api.Logout)
	http.HandleFunc("/api/add", api.AddTask)
	http.HandleFunc("/api/tasks/edit", api.EditTask)
	http.HandleFunc("/api/tasks/edit/cancel", api.CancelEditTask)
	http.HandleFunc("/api/tasks/remove", api.RemoveTask)
  http.HandleFunc("/api/tasks/complete", api.CompleteTask)
	http.Handle("/*", templ.Handler(notFound))

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), nil))
}
