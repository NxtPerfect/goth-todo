package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"

	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	godotenv.Load()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	fmt.Printf(psqlInfo)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	conn.Ping()

	return conn
}

// Returns all tasks for user
func GetTasks(email string, conn *sql.DB) *sql.Rows {
	rows, err := conn.Query("SELECT * FROM tasks WHERE userId = (SELECT id FROM users WHERE email = $1", email)
	if err != nil {
		panic(err)
	}
	return rows
}

// Inserts new tasks for user
// func SetTasks(userId string, tasks []api.types.Tasks) {
// 	// insert tasks to db
// }

// Modifies tasks for user
func ModifyTasks(userId string, taskid string) {
	//
}
