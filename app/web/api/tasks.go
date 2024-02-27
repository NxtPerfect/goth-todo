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

// Render task edit form
func EditTaskForm(w http.ResponseWriter, r *http.Request) {
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

// Changes task in database
// assumes user got correct tasks
func EditTask(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	date_due := r.PostFormValue("date_due")
	finished := false
	date_modified := time.Now().Format("2006-01-02")

	// Query database to edit
	conn := database.Connect()
	_, err := conn.Query("UPDATE tasks SET id = $6, title = $1, description = $2, finished = $3, created_at = (SELECT created_at FROM tasks WHERE id = $6), last_modified = $4, due_at = $5 WHERE id = $6;", title, description, finished, date_modified, date_due, id)

	if err != nil {
		w.Header().Add("HX-Retarget", "closest li")
		w.Header().Add("HX-Reswap", "beforeend")
		templates.ErrorMessage("Invalid task edited.").Render(r.Context(), w)
		http.Error(w, "Bad edit request", http.StatusBadRequest)
		panic(err)
	}

  templates.Task(types.Task{Id: id, Title: title, Description: description, Finished: finished, Date_due: date_due, Date_modified: date_modified, Date_created: "1999-01-02"}, 99).Render(r.Context(), w)
}

// TODO remove current task
func RemoveTask(w http.ResponseWriter, r *http.Request) {
  id := r.PostFormValue("id")

  // TODO should check if auth token matches, to avoid deleting wrong task
  // for that, we should get the real id from db into session or some form
  // of cache
	conn := database.Connect()
	_, err := conn.Query("DELETE FROM tasks WHERE id = $1;", id)

	if err != nil {
		panic(err)
	}
}

// After canceling edit, return the task
func CancelEditTask(w http.ResponseWriter, r *http.Request) {
  id := r.PostFormValue("id")
  title := r.PostFormValue("title")
  description := r.PostFormValue("description")
  date_due := r.PostFormValue("date_due")

  templates.Task(types.Task{Id: id, Title: title, Description: description, Finished: false, Date_due: date_due, Date_modified: time.Now().Format("2006-01-02"), Date_created: time.Now().Format("2006-01-02")}, 99).Render(r.Context(), w)
}

// TODO mark task as finished in db
func CompleteTask(w http.ResponseWriter, r *http.Request) {
  id := r.PostFormValue("id")

  // !!TODO!! UNTESTED
	conn := database.Connect()
	_, err := conn.Query("UPDATE tasks SET finished = true WHERE id = $1;", id)

	if err != nil {
		panic(err)
	}
}
