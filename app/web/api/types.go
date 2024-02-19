package api

type Task struct {
	Id          int32
	Title       string
	Description string
	// date_created string
	// date_modified string
	// date_due string
}

type User struct {
	Id       int32
	username string
	email    string
	password string
}
