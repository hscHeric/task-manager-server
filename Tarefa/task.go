package main

import (
	"time"

	"github.com/google/uuid"
)

// Task representa uma tarefa com todos os detalhes.
type Task struct {
	TaskID      uuid.UUID  `json:"taskId"`         // Identificador único da tarefa
	Title       string     `json:"title"`          // Título da tarefa
	Description string     `json:"description"`    // Descrição da tarefa
	Date        *time.Time `json:"date,omitempty"` // Data da tarefa, opcional
}

// NewTask cria uma nova instância de Task com os parâmetros fornecidos.
func NewTask(title string, description string, date *time.Time) *Task {
	return &Task{
		TaskID:      uuid.New(), // Gera um novo UUID para a tarefa
		Title:       title,
		Description: description,
		Date:        date,
	}
}
