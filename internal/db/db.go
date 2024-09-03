package db

import (
	"database/sql"
	"time"

	"github.com/hscHeric/task-manager-server/internal/task"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseService struct {
	Db *sql.DB
}

func NewDatabaseService(db *sql.DB) *DatabaseService {
	return &DatabaseService{
		Db: db,
	}
}

func (db *DatabaseService) InsertTask(t *task.Task) error {
	date := t.Date.Format("02/01/2006")
	_, err := db.Db.Exec("INSERT INTO tasks (id, title, description, date) VALUES (?, ?, ?, ?)", t.TaskID, t.Title, t.Description, date)
	return err
}

func (db *DatabaseService) GetAllTasks() ([]*task.Task, error) {
	rows, err := db.Db.Query("SELECT id, title, description, date FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task

	for rows.Next() {
		var t task.Task
		var dateString string
		if err := rows.Scan(&t.TaskID, &t.Title, &t.Description, &dateString); err != nil {
			return nil, err
		}

		date, err := time.Parse("02/01/2006", dateString)
		if err != nil {
			return nil, err
		}
		t.Date = date

		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *DatabaseService) GetTaskByID(id string) (*task.Task, error) {
	row := db.Db.QueryRow("SELECT id, title, description, date FROM tasks WHERE id = ?", id)

	var t task.Task
	var dateString string
	if err := row.Scan(&t.TaskID, &t.Title, &t.Description, &dateString); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Nenhuma tarefa encontrada com o ID fornecido
		}
		return nil, err
	}

	date, err := time.Parse("02/01/2006", dateString)
	if err != nil {
		return nil, err
	}
	t.Date = date

	return &t, nil
}

func (db *DatabaseService) DeleteTask(id string) error {
	_, err := db.Db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
