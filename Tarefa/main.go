package main

import (
	"fmt"
	"time"
)

// Função para listar todas as tarefas e imprimir a resposta
func listarTarefas(taskManager *arrayTask) {
	fmt.Println("\nListando tarefas...")
	listResponse := taskManager.ListTask()
	fmt.Println("List Response:")
	for _, task := range listResponse.Tasks {
		fmt.Printf("Título: %s\nDescrição: %s\nData: %v\n\n",
			task.Title, task.Description, task.Date.Format("02/01/2006"))
	}
	fmt.Println("Success:", listResponse.Success)
	fmt.Println("Error:", listResponse.Error)
}

func main() {
	taskManager := NewTasks()

	date1 := time.Date(2020, 5, 6, 11, 45, 04, 0, time.UTC)
	date2 := time.Date(2021, 7, 6, 11, 45, 04, 0, time.UTC)
	date3 := time.Date(2000, 5, 8, 11, 45, 04, 0, time.UTC)

	// Criando tarefas
	task1 := NewTask("Tarefa 1", "Descrição da tarefa 1", &date1)
	task2 := NewTask("Tarefa 2", "Descrição da tarefa 2", &date2)
	task3 := NewTask("Tarefa 3", "Descrição da tarefa 3", &date3)

	// Adiciona as três tarefas
	fmt.Println("Adicionando tarefas...")
	taskManager.AddTask(task1)
	taskManager.AddTask(task2)
	taskManager.AddTask(task3)

	listarTarefas(taskManager)

	// Edita a segunda tarefa
	newTitle := "Título Atualizado"
	newDescription := "Descrição Atualizada"
	fmt.Println("\nEditando a segunda tarefa...")
	editRequest := NewUpdateTaskRequest(task2.TaskID, &newTitle, &newDescription, nil)
	editResponse := taskManager.EditTask(editRequest)
	fmt.Println("Edit Response:", editResponse)

	listarTarefas(taskManager)

	// Remove a terceira tarefa
	fmt.Println("\nRemovendo a terceira tarefa...")
	taskManager.RemoveTask(task3.TaskID)

	listarTarefas(taskManager)
}
