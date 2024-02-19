package api

import (
	"net/http"
	"todo/app/web/database"
)

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
  email := r.PostFormValue("email")
	password := r.PostFormValue("password")
  confirm_password := r.PostFormValue("confirm_password")

  if password != confirm_password {
    http.Error(w, "Passwords don't match", http.StatusBadRequest)
    return
  }

  // check if email is email, verify lengths > 0, verify alphanumeric
  // then insert into database
  // hash password
  conn := database.Connect()

	_, err := conn.Query("INSERT INTO users (id, username, email, password) VALUES ( uuid_generate_v4(), $1, $2, $3 )", username, email, password)
	if err != nil {
		panic(err)
	}

  http.Redirect(w, r, "/", http.StatusSeeOther)
  return
}
