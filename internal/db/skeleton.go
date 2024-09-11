package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hscHeric/task-manager-server/internal/task"
	_ "github.com/mattn/go-sqlite3" // Importa o driver SQLite
)

const dbPath = "/home/gabriel/Downloads/SD_ENVIO_03/task-manager-server/cmd/server/tasks.db"

type Skeleton struct{}

func NewSkeleton() *Skeleton {
	return &Skeleton{}
}

func initDB() *sql.DB {
	dbConn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(fmt.Errorf("erro ao abrir o banco de dados: %w", err))
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatal(fmt.Errorf("erro ao conectar ao banco de dados: %w", err))
	}

	fmt.Println("Conectado ao banco de dados com sucesso")
	return dbConn
}

type request struct {
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type ID struct {
	TaskID string `json:"taskId"`
}

func (s *Skeleton) InsertTask(data []byte) ([]byte, error) {
	var req request
	dbConn := initDB()
	defer dbConn.Close()

	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("falha ao desserializar o JSON da requisição: %w", err)
	}

	newTask := task.NewTask(req.Title, req.Description, req.Date)
	dbService := NewDatabaseService(dbConn)

	if err := dbService.InsertTask(newTask); err != nil {
		return nil, fmt.Errorf("falha ao inserir a tarefa no banco de dados: %w", err)
	}

	return json.Marshal("Tarefa inserida com sucesso")
}

func (s *Skeleton) GetAllTasks() ([]byte, error) {
	dbConn := initDB()
	defer dbConn.Close()
	dbService := NewDatabaseService(dbConn)

	tasks, err := dbService.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("falha ao obter as tarefas: %w", err)
	}

	return json.Marshal(tasks)
}

func (s *Skeleton) GetTaskByID(data []byte) ([]byte, error) {
	dbConn := initDB()
	defer dbConn.Close()
	dbService := NewDatabaseService(dbConn)

	var id ID
	if err := json.Unmarshal(data, &id); err != nil {
		return nil, fmt.Errorf("erro ao desserializar o ID: %w", err)
	}

	t, err := dbService.GetTaskByID(id.TaskID)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter a tarefa: %w", err)
	}

	return json.Marshal(t)
}

func (s *Skeleton) DeleteTask(data []byte) ([]byte, error) {
	dbConn := initDB()
	defer dbConn.Close()
	dbService := NewDatabaseService(dbConn)

	var id ID
	if err := json.Unmarshal(data, &id); err != nil {
		return nil, fmt.Errorf("erro ao desserializar o ID: %w", err)
	}

	if err := dbService.DeleteTask(id.TaskID); err != nil {
		return nil, fmt.Errorf("erro ao excluir a tarefa: %w", err)
	}

	return json.Marshal("Tarefa excluída com sucesso")
}

