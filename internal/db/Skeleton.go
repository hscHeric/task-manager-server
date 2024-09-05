package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hscHeric/task-manager-server/internal/task"
)

const dbPath = "./tasks.db"

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
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

// Response é a estrutura de resposta padrão para todas as operações do Skeleton
type Response struct {
	Data  interface{} `json:"data,omitempty"` // Pode ser qualquer tipo: Task, lista de Tasks, etc.
	Error string      `json:"error,omitempty"`
}

// NewResponse cria uma nova resposta padronizada com dados ou erro
func NewResponse(data interface{}, err error) *Response {
	if err != nil {
		return &Response{Error: err.Error()}
	}
	return &Response{Data: data}
}

func (s *Skeleton) InsertTask(data []byte) *Response {
	var req request
	dbConn := initDB()
	defer dbConn.Close()

	// Desserializa o JSON para a estrutura `request`
	if err := json.Unmarshal(data, &req); err != nil {
		return NewResponse(nil, fmt.Errorf("falha ao desserializar o JSON da requisição: %w", err))
	}

	// Cria uma nova tarefa com os dados da requisição
	newTask := task.NewTask(req.Title, req.Description, req.Date)
	dbService := NewDatabaseService(dbConn)

	if err := dbService.InsertTask(newTask); err != nil {
		return NewResponse(nil, fmt.Errorf("falha ao inserir a tarefa no banco de dados: %w", err))
	}

	return NewResponse("Tarefa inserida com sucesso", nil)
}

func (s *Skeleton) GetAllTasks() *Response {
	dbConn := initDB()
	defer dbConn.Close()
	dbService := NewDatabaseService(dbConn)

	tasks, err := dbService.GetAllTasks()
	if err != nil {
		return NewResponse(nil, fmt.Errorf("falha ao obter as tarefas: %w", err))
	}

	return NewResponse(tasks, nil)
}

type ID struct {
	TaskID string `json:"taskId"`
}

func (s *Skeleton) GetTaskByID(data []byte) *Response {
	dbConn := initDB()
	defer dbConn.Close()
	dbService := NewDatabaseService(dbConn)

	var id ID
	err := json.Unmarshal(data, &id)
	if err != nil {
		return NewResponse(nil, fmt.Errorf("erro ao desserializar o ID: %w", err))
	}

	task, err := dbService.GetTaskByID(id.TaskID)
	if err != nil {
		return NewResponse(nil, fmt.Errorf("erro ao obter a tarefa: %w", err))
	}

	return NewResponse(task, nil)
}

func (s *Skeleton) DeleteTask(data []byte) *Response {
	dbConn := initDB()
	defer dbConn.Close()
	dbService := NewDatabaseService(dbConn)

	var id ID
	if err := json.Unmarshal(data, &id); err != nil {
		return NewResponse(nil, fmt.Errorf("erro ao desserializar o ID: %w", err))
	}

	err := dbService.DeleteTask(id.TaskID)
	if err != nil {
		return NewResponse(nil, fmt.Errorf("erro ao excluir a tarefa: %w", err))
	}

	return NewResponse("Tarefa excluída com sucesso", nil)
}
