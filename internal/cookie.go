package internal

import (
	"net/http"
	"time"
)

func CreateUsernameCookie(username string) http.Cookie {
  expiration := time.Now().Add(30 * 24 * time.Hour)
  cookie := http.Cookie{Name: "username", Value: username, Expires: expiration, SameSite: http.SameSiteLaxMode, Path: "/"}
  return cookie
}

func CreateAuthtokenCookie (userid string) http.Cookie {
  expiration := time.Now().Add(30 * 24 * time.Hour)
  token, _ := GenerateAuthToken(userid)
  cookie := http.Cookie{Name: "auth_token", Value: token, Expires: expiration, HttpOnly: true, SameSite: http.SameSiteLaxMode, Path: "/"}
  return cookie
}
