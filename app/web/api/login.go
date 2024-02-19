package api

import (
	"fmt"
	"net/http"
	"todo/app/web/database"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// Check in database if exists and password is correct
	conn := database.Connect()
	rows, err := conn.Query("SELECT password FROM users WHERE username = $1;", username)
	// rows, err := conn.Query("SELECT title FROM tasks WHERE userId = (SELECT id FROM users WHERE username = $1);", username)
	if err != nil {
		panic(err)
	}

	var real_pass string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&real_pass)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nDid passwords match? %t", password == real_pass)

    if password == real_pass {
      http.Redirect(w, r, "/", http.StatusSeeOther)
      return
    }
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

  http.Error(w, "Invalid credentials", http.StatusUnauthorized)
  return

}
