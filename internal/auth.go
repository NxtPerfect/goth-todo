package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

func GenerateAuthToken(userId string, seed string) (string, error) {
	h := sha256.New()

	_, err := h.Write([]byte(userId + seed))
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func ValidateAuthToken(userId string, authtoken string, seed string) bool {
	validtoken, _ := GenerateAuthToken(userId, seed)
	return validtoken == authtoken
}

func GetUsernameCookie(r http.Request) string {
	usernameCookie, err := r.Cookie("username")
	if err != nil {
		panic(err)
	}
	return usernameCookie.Value
}
