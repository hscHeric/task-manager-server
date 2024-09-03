package main

import (
	"sort"

	"github.com/google/uuid"
)

// Estrutura para resposta após adicionar uma tarefa
type AddResponse struct {
	TaskID  uuid.UUID `json:"taskId"`
	Success bool      `json:"success"`
	Error   string    `json:"error,omitempty"`
}

// Estrutura para retornar resposta
type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// Estrutura para resposta ao listar as tarefas
type ListResponse struct {
	Tasks   []*Task `json:"tasks"`
	Success bool    `json:"success"`
	Error   string  `json:"error,omitempty"`
}

type arrayTask struct {
	TaskList []*Task
}

func NewTasks() *arrayTask {
	return &arrayTask{}
}

// Função para adicionar uma nova tarefa e retornar uma resposta
func (t *arrayTask) AddTask(tsk *Task) *AddResponse {
	if tsk == nil {
		return &AddResponse{
			TaskID:  uuid.Nil,
			Success: false,
			Error:   "Tarefa não pode ser nula",
		}
	}

	t.TaskList = append(t.TaskList, tsk)

	return &AddResponse{
		TaskID:  tsk.TaskID,
		Success: true,
		Error:   "",
	}
}

// Função para remover uma tarefa
func (t *arrayTask) RemoveTask(taskID uuid.UUID) *Response {
	var taskFound bool
	for i, task := range t.TaskList {
		if task.TaskID == taskID {
			t.TaskList = append(t.TaskList[:i], t.TaskList[i+1:]...)
			taskFound = true
			break
		}
	}

	if !taskFound {
		return &Response{
			Success: false,
			Error:   "Tarefa não encontrada",
		}
	}

	return &Response{
		Success: true,
		Error:   "",
	}
}

// Função para editar uma tarefa
func (t *arrayTask) EditTask(u *UpdateTaskRequest) *Response {
	var taskFound bool
	for _, task := range t.TaskList {
		if task.TaskID == u.TaskID {
			if u.Title != nil {
				task.Title = *u.Title
			}
			if u.Description != nil {
				task.Description = *u.Description
			}
			if u.Date != nil {
				task.Date = u.Date
			}
			taskFound = true
			break
		}
	}

	if !taskFound {
		return &Response{
			Success: false,
			Error:   "Tarefa não encontrada",
		}
	}

	return &Response{
		Success: true,
		Error:   "",
	}
}

// Função para listar todas as tarefas em ordem de data
func (t *arrayTask) ListTask() *ListResponse {
	if len(t.TaskList) == 0 {
		return &ListResponse{
			Tasks:   nil,
			Success: false,
			Error:   "Nenhuma tarefa encontrada",
		}
	}

	sort.Slice(t.TaskList, func(i, j int) bool {
		if t.TaskList[i].Date == nil && t.TaskList[j].Date == nil {
			return false
		}
		if t.TaskList[i].Date == nil {
			return false
		}
		if t.TaskList[j].Date == nil {
			return true
		}
		return t.TaskList[i].Date.Before(*t.TaskList[j].Date)
	})

	return &ListResponse{
		Tasks:   t.TaskList,
		Success: true,
		Error:   "",
	}
}
