package task

import (
	"time"

	"github.com/google/uuid"
)

// Task representa uma tarefa com todos os detalhes.
type Task struct {
	Date        time.Time `json:"date,omitempty"` // Data da tarefa, opcional
	TaskID      string    `json:"taskId"`         // Identificador único da tarefa
	Title       string    `json:"title"`          // Título da tarefa
	Description string    `json:"description"`    // Descrição da tarefa
}

// NewTask cria uma nova instância de Task com os parâmetros fornecidos.
func NewTask(title string, description string, date time.Time) *Task {
	return &Task{
		TaskID:      uuid.NewString(), // Gera um novo UUID para a tarefa
		Title:       title,
		Description: description,
		Date:        date,
	}
}

func (t *Task) GetDateString() string {
	return t.Date.Format("02/01/2006")
}
