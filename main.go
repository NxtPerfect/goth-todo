package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})
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
  http.HandleFunc("/tasks/edit", func (w http.ResponseWriter, r *http.Request)  {
    id := r.Form.Get("id")
    title := r.Form.Get("title")
    description := r.Form.Get("description")
    finished := r.Form.Get("finished")
    date_due := r.Form.Get("date_due")
    date_created := r.Form.Get("date_created")
    var task types.Task
    task.Id = id
    task.Title = title
    task.Description = description
    var err error
    task.Finished, err = strconv.ParseBool(finished)
    if err != nil {
      panic(err)
    }
    task.Date_due = date_due
    task.Date_created = date_created
    task.Date_modified = time.Now().Format("2006-01-02")
    templates.EditForm(task).Render(r.Context(), w)
  })
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/register", api.Register)
	http.HandleFunc("/api/logout", api.Logout)
	http.HandleFunc("/api/add", api.AddTask)
	http.HandleFunc("/api/tasks/edit", api.EditTask)
	http.HandleFunc("/api/tasks/remove", api.RemoveTask)
  http.Handle("/*", templ.Handler(notFound))

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), nil))
}
