package api

import (
	"fmt"
	"net/http"
	"time"
	"todo/app/web/database"
	"todo/app/web/templates"
	"todo/app/web/types"
	"todo/internal"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	date_due := r.PostFormValue("date_due")
	auth_token, err := r.Cookie("auth_token")

	if err != nil {
		panic(err)
	}

	username, err := r.Cookie("username")

	if err != nil {
		panic(err)
	}

	// Check in database if exists and password is correct
	conn := database.Connect()
	rows, err := conn.Query("SELECT id FROM users WHERE username = $1;", username.Value)

	if err != nil {
		panic(err)
	}

	var real_id string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&real_id)
		if err != nil {
			panic(err)
		}
		real_token, err := internal.GenerateAuthToken(real_id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nDid id's match? %t\n", real_token == auth_token.Value)

		// Verify it's our user, if no then continue
		if real_token != auth_token.Value {
			continue
		}

		// Create new task, since the tokens matched
		_, err = conn.Query("INSERT INTO tasks (id, userid, title, description, finished, created_at, last_modified, due_at) VALUES (uuid_generate_v4(), $1, $2, $3, false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $4);", real_id, title, description, date_due)
		if err != nil {
			http.Error(w, "Couldn't create task.", http.StatusInternalServerError)
			panic(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	http.Error(w, "Invalid credentials. User doesn't exist. From error message", http.StatusUnauthorized)
	return

}

func EditTask(w http.ResponseWriter, r *http.Request) {
  task_id := r.URL.Query().Get("id")
  title := r.URL.Query().Get("title")
  description := r.URL.Query().Get("description")
  date_due := r.URL.Query().Get("date_due")
  date_created := r.URL.Query().Get("date_created")
  // fmt.Printf("%s", time.Parse("2006-02-01:T15:04:05.999999999Z07:00", date_due)
  n_date_due, err := time.Parse(time.RFC3339Nano, date_due)
  if err != nil {
    panic(err)
  }
  fmt.Printf("%s", time.Now().String())
  templates.EditForm(types.Task{Id: task_id, Title: title, Description: description, Finished: false, Date_created: date_created, Date_modified: time.Now().Format("2016-01-02"), Date_due: n_date_due.Format("2006-01-02")}).Render(r.Context(), w)
}

func RemoveTask(w http.ResponseWriter, r *http.Request) {
  // task_id := r.PostFormValue("id")
}
