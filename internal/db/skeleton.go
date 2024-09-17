package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/hscHeric/task-manager-server/internal/task"
	_ "github.com/mattn/go-sqlite3" // Importa o driver SQLite
)

const dbPath = "/home/gabriel/Downloads/SD_ENVIO_03/task-manager-server/cmd/server/tasks.db"

type Skeleton struct {
	db      *sql.DB
	initErr error
}

var (
	instance *Skeleton
	once     sync.Once
)

func NewSkeleton() *Skeleton {
	once.Do(func() {
		dbConn, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			instance = &Skeleton{
				initErr: fmt.Errorf("erro ao abrir o banco de dados: %w", err),
			}
			return
		}

		if err = dbConn.Ping(); err != nil {
			instance = &Skeleton{
				initErr: fmt.Errorf("erro ao conectar ao banco de dados: %w", err),
			}
			return
		}

		fmt.Println("Conectado ao banco de dados com sucesso")

		instance = &Skeleton{
			db: dbConn,
		}
	})

	return instance
}

func (s *Skeleton) GetInitError() error {
	return s.initErr
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

	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("falha ao desserializar o JSON da requisição: %w", err)
	}

	newTask := task.NewTask(req.Title, req.Description, req.Date)
	dbService := NewDatabaseService(s.db)

	if err := dbService.InsertTask(newTask); err != nil {
		return nil, fmt.Errorf("falha ao inserir a tarefa no banco de dados: %w", err)
	}

	return json.Marshal("Tarefa inserida com sucesso")
}

func (s *Skeleton) GetAllTasks() ([]byte, error) {
	dbService := NewDatabaseService(s.db)

	tasks, err := dbService.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("falha ao obter as tarefas: %w", err)
	}

	return json.Marshal(tasks)
}

func (s *Skeleton) GetTaskByID(data []byte) ([]byte, error) {
	var id ID
	if err := json.Unmarshal(data, &id); err != nil {
		return nil, fmt.Errorf("erro ao desserializar o ID: %w", err)
	}

	dbService := NewDatabaseService(s.db)
	t, err := dbService.GetTaskByID(id.TaskID)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter a tarefa: %w", err)
	}

	return json.Marshal(t)
}

func (s *Skeleton) DeleteTask(data []byte) ([]byte, error) {
	var id ID
	if err := json.Unmarshal(data, &id); err != nil {
		return nil, fmt.Errorf("erro ao desserializar o ID: %w", err)
	}

	dbService := NewDatabaseService(s.db)
	if err := dbService.DeleteTask(id.TaskID); err != nil {
		return nil, fmt.Errorf("erro ao excluir a tarefa: %w", err)
	}

	return json.Marshal("Tarefa excluida com sucesso")
}
