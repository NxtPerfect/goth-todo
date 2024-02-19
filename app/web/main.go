package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo/app/web/database"
	"todo/app/web/templates"

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

	h4 := func(w http.ResponseWriter, r *http.Request) {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		// Check in database if exists and password is correct
		conn := database.Connect()
		rows, err := conn.Query("SELECT title FROM tasks WHERE userId = (SELECT id FROM users WHERE email = $1", "admin@admin.io")
		if err != nil {
			panic(err)
		}
		conn.Close()

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
	}

	home := templates.HomePage()
	login := templates.LoginPage()
	register := templates.RegisterPage()

	http.Handle("/", templ.Handler(home))
	http.Handle("/login", templ.Handler(login))
	http.Handle("/register", templ.Handler(register))
	http.HandleFunc("/api/login", h4)
	// http.Handle("/api/tasks/0/edit", templ.Handler(login))

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), nil))
}
