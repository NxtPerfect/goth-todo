package types

type Task struct {
	id            int32
	title         string
	description   string
	date_created  string
	date_modified string
	date_due      string
}

type User struct {
	id       int32
	username string
	email    string
	password string
}
