package main

import (
	"time"

	"github.com/google/uuid"
)

// UpdateTaskRequest representa uma solicitação para atualizar uma tarefa.
type UpdateTaskRequest struct {
	TaskID      uuid.UUID  `json:"taskId"`                // Identificador da tarefa a ser atualizada
	Title       *string    `json:"title,omitempty"`       // Novo título da tarefa, opcional
	Description *string    `json:"description,omitempty"` // Nova descrição da tarefa, opcional
	Date        *time.Time `json:"date,omitempty"`        // Nova data da tarefa, opcional
}

// NewUpdateTaskRequest cria uma nova instância de UpdateTaskRequest.
func NewUpdateTaskRequest(taskId uuid.UUID, title *string, description *string, date *time.Time) *UpdateTaskRequest {
	return &UpdateTaskRequest{
		TaskID:      taskId,
		Title:       title,
		Description: description,
		Date:        date,
	}
}

// RemoveTaskRequest representa uma solicitação para remover uma tarefa.
type RemoveTaskRequest struct {
	TaskID uuid.UUID `json:"taskId"` // Identificador da tarefa a ser removida
}

// NewRemoveTaskRequest cria uma nova instância de RemoveTaskRequest.
func NewRemoveTaskRequest(taskId uuid.UUID) *RemoveTaskRequest {
	return &RemoveTaskRequest{
		TaskID: taskId,
	}
}
