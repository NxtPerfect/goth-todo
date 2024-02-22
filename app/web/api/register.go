package api

import (
	"net/http"
	"todo/app/web/database"
	"todo/internal"

	"github.com/goware/emailx" // Email validation
)

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
  email := r.PostFormValue("email")
	password := r.PostFormValue("password")
  confirm_password := r.PostFormValue("confirm_password")

  if password != confirm_password {
    http.Error(w, "Passwords don't match.", http.StatusBadRequest)
    return
  }

  err := emailx.Validate(email)
  if err != nil {
    if err == emailx.ErrInvalidFormat {
      http.Error(w, "Invalid email format.", http.StatusBadRequest)
    }
    if err == emailx.ErrUnresolvableHost {
      http.Error(w, "Unresolvable email host.", http.StatusBadRequest)
    }
    panic(err)
  }

  // TODO: check if email is email, verify lengths > 0, verify alphanumeric
  // then insert into database
  // TODO: hash password
  conn := database.Connect()

	_, err = conn.Query("INSERT INTO users (id, username, email, password) VALUES ( uuid_generate_v4(), $1, $2, $3 )", username, emailx.Normalize(email), password)
	if err != nil {
		panic(err)
	}
  
  // Get generated user id
  row, err := conn.Query("SELECT id FROM users WHERE email = $1", emailx.Normalize(email))

  cookie := internal.CreateUsernameCookie(username)
  http.SetCookie(w, &cookie)

  // Get userid from rows
  var userid string
  defer row.Close()
  row.Next()
  row.Scan(&userid)

  cookie = internal.CreateAuthtokenCookie(userid)
  http.SetCookie(w, &cookie)
  http.Redirect(w, r, "/", http.StatusSeeOther)
  return
}
