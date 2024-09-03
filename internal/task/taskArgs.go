package task

// RemoveTaskRequest representa uma solicitação para remover uma tarefa.
type IDTaskRequest struct {
	TaskID string `json:"taskId"` // Identificador da tarefa a ser removida
}

// NewRemoveTaskRequest cria uma nova instância de RemoveTaskRequest.
func NewRemoveTaskRequest(taskId string) *IDTaskRequest {
	return &IDTaskRequest{
		TaskID: taskId,
	}
}
