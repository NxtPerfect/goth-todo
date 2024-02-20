package types

type Task struct {
	Id            string
	Title         string
	Description   string
	Finished      bool
	Date_created  string
	Date_modified string
	Date_due      string
}

type User struct {
	Id       int32
	Username string
	Email    string
	Password string
}
