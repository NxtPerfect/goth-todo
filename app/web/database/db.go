package database

import (
	"database/sql"
	"fmt"
	"os"
	"todo/app/web/types"

	_ "github.com/lib/pq"

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
func GetTasks(username string, authtoken string, conn *sql.DB) []types.Task {
	rows, err := conn.Query("SELECT authtoken, id FROM users WHERE username = $1", username)
	if err != nil {
		panic(err)
	}

	var token string
	var id string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&token, &id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nDid tokens match? %t\n", authtoken == token)

		// If it's the password, setup auth token and send to user
		if authtoken == token {
			row, err := conn.Query("SELECT * FROM tasks WHERE id = $1;", id)
			if err != nil {
				panic(err)
			}

			var tasks []types.Task
			// Since there can be only one record as email is unique
			// we can call row.Next once and get id and username
			row.Next()
			row.Scan(&tasks)
			return tasks
		}
	}

	return
}

// Inserts new tasks for user
// func SetTasks(userId string, tasks []api.types.Tasks) {
// 	// insert tasks to db
// }

// Modifies tasks for user
func ModifyTasks(userId string, taskid string) {
	//
}
