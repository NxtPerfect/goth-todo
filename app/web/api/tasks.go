package api

import (
	"fmt"
	"net/http"
	"todo/app/web/database"
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
	rows, err := conn.Query("SELECT id FROM users WHERE username = $1;", username)

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
		// TODO: NOT TESTED
		_, err = conn.Query("INSERT INTO tasks (id, userid, title, description, finished, created_at, last_modified, due_at) VALUES (uuid_generate_v4(), $1, $2, $3, false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $4);", real_id, title, description, date_due)
		if err != nil {
			http.Error(w, "Couldn't create task.", http.StatusInternalServerError)
			panic(err)
		}

		http.Redirect(w, r, "/", http.StatusAccepted)
		return
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	http.Error(w, "Invalid credentials. User doesn't exist.", http.StatusUnauthorized)
	return

}
