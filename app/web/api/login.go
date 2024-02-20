package api

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"todo/app/web/database"
	"todo/internal"

	"github.com/joho/godotenv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// Check in database if exists and password is correct
	conn := database.Connect()
	rows, err := conn.Query("SELECT password FROM users WHERE email = $1;", email)

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
		fmt.Printf("\nDid passwords match? %t\n", password == real_pass)

		// If it's the password, setup auth token and send to user
		if password == real_pass {
			row, err := conn.Query("SELECT id, username FROM users WHERE email = $1;", email)
			if err != nil {
				panic(err)
			}

			var userid string
			var username string
			// Since there can be only one record as email is unique
			// we can call row.Next once and get id and username
			row.Next()
			row.Scan(&userid, &username)
			expiration := time.Now().Add(30 * 24 * time.Hour)
			token, _ := internal.GenerateAuthToken(userid, os.Getenv("AUTH_TOKEN_SEED"))
			cookie := http.Cookie{Name: "auth_token", Value: token, Expires: expiration, HttpOnly: true, SameSite: http.SameSiteLaxMode, Path: "/"}
			http.SetCookie(w, &cookie)
			cookie = http.Cookie{Name: "username", Value: username, Expires: expiration, SameSite: http.SameSiteLaxMode, Path: "/"}
			http.SetCookie(w, &cookie)
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
