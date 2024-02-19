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
	rows, err := conn.Query("SELECT title FROM tasks WHERE userId = (SELECT id FROM users WHERE email = $1);", "admin@admin.io")
	if err != nil {
		panic(err)
	}

	var title string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&title)
		if err != nil {
			panic(err)
		}
		fmt.Println("\n", title)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	if username == "admin" && password == "admin" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
  return
}
