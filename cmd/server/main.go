package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/hscHeric/task-manager-server/internal/db" // Assuming this is your project path
	"github.com/hscHeric/task-manager-server/internal/task"
)

const dbPath = "./tasks.db"

func main() {
	fmt.Println("Aqui")
	dbConn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	defer dbConn.Close()

	dbService := db.NewDatabaseService(dbConn)

	dt := time.Now()
	newTask := task.NewTask("ola mundo", "ola mundo", dt) // Assuming this creates a task object
	if err := dbService.InsertTask(newTask); err != nil {
		log.Fatal(err)
	}

	tasks, err := dbService.GetAllTasks()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tasks:")
	for _, t := range tasks {
		fmt.Println(t)
	}
}
