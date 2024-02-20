package database

import (
	"database/sql"
	"fmt"
	"os"
	"todo/app/web/types"
	"todo/internal"

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
	rows, err := conn.Query("SELECT id FROM users WHERE username = $1", username)
	if err != nil {
		panic(err)
	}

	var id string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		generated_token, err := internal.GenerateAuthToken(id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nDid tokens match? %t\n", authtoken == generated_token)

		// If it's the password, setup auth token and send to user
		if authtoken == generated_token {
			// authtoken, err := uuid.Parse(authtoken)
			row, err := conn.Query("SELECT * FROM tasks WHERE userid = $1;", id)
			if err != nil {
				panic(err)
			}

			var tasks []types.Task
			// For each task, append to array, reject user_id
			for row.Next() {
				var task types.Task
				err = row.Scan(&task.Id, new(interface{}), &task.Title, &task.Description, &task.Finished, &task.Date_created, &task.Date_modified, &task.Date_due)
				if err != nil {
					panic(err)
				}
				if err != nil {
					panic(err)
				}
				tasks = append(tasks, task)
			}
			return tasks
		}
	}

	return []types.Task{}
}

// Inserts new tasks for user
// func SetTasks(userId string, tasks []api.types.Tasks) {
// 	// insert tasks to db
// }

// Modifies tasks for user
func ModifyTasks(userId string, taskid string) {
	//
}
