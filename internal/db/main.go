package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hscHeric/task-manager-server/internal/db"
	"github.com/hscHeric/task-manager-server/internal/task"
)

func insertTask(d *db.Dispacher) {
	taskData := `{"title": "Nova Tarefa", "description": "Descrição da nova tarefa", "date": "2024-09-05T00:00:00Z"}`

	resp, err := d.Invoke("InsertTask", []byte(taskData))
	if err != nil {
		log.Fatalf("Erro ao inserir tarefa: %v", err)
	}

	// Desserializando a resposta
	var response db.Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Fatalf("Erro ao desserializar a resposta: %v", err)
	}

	if response.Error != "" {
		log.Fatalf("Erro na resposta: %s", response.Error)
	}

	fmt.Println("Tarefa inserida com sucesso.")
}

func getAllTasks(d *db.Dispacher) {
	resp, err := d.Invoke("GetAllTasks", nil)
	if err != nil {
		log.Fatalf("Erro ao obter todas as tarefas: %v", err)
	}

	var response db.Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Fatalf("Erro ao desserializar a resposta: %v", err)
	}

	if response.Error != "" {
		log.Fatalf("Erro na resposta: %s", response.Error)
	}

	// Desserializa os dados da resposta
	var tasks []task.Task

	// Se response.Data é uma interface, precisamos converter para um []interface{}
	switch data := response.Data.(type) {
	case []interface{}:
		// Converta []interface{} para []byte
		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("Erro ao converter []interface{} para []byte: %v", err)
		}

		// Desserialize os dados para uma lista de tarefas
		err = json.Unmarshal(dataBytes, &tasks)
		if err != nil {
			log.Fatalf("Erro ao desserializar as tarefas: %v", err)
		}
	default:
		log.Fatalf("Tipo de dados inesperado: %T", response.Data)
	}

	fmt.Println("Tasks:")
	for _, t := range tasks {
		fmt.Printf("ID: %s, Title: %s, Description: %s, Date: %s\n",
			t.TaskID, t.Title, t.Description, t.GetDateString())
	}
}
func getTaskByID(d *db.Dispacher, taskID string) {
	idData := fmt.Sprintf(`{"taskId": "%s"}`, taskID)

	resp, err := d.Invoke("GetTaskByID", []byte(idData))
	if err != nil {
		log.Fatalf("Erro ao obter tarefa por ID: %v", err)
	}

	var response db.Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Fatalf("Erro ao desserializar a resposta: %v", err)
	}

	if response.Error != "" {
		log.Fatalf("Erro na resposta: %s", response.Error)
	}

	// Desserializa os dados da resposta
	var task task.Task

	// Usa type assertion para garantir que response.Data é do tipo correto
	switch data := response.Data.(type) {
	case map[string]interface{}:
		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("Erro ao converter map[string]interface{} para JSON: %v", err)
		}
		err = json.Unmarshal(dataBytes, &task)
		if err != nil {
			log.Fatalf("Erro ao desserializar a tarefa: %v", err)
		}
	default:
		log.Fatalf("Tipo de dados inesperado: %T", response.Data)
	}

	fmt.Println("Task:")
	fmt.Printf("ID: %s\nTitle: %s\nDescription: %s\nDate: %s\n",
		task.TaskID, task.Title, task.Description, task.GetDateString())
}

func deleteTask(d *db.Dispacher, taskID string) {
	idData := fmt.Sprintf(`{"taskId": "%s"}`, taskID)

	resp, err := d.Invoke("DeleteTask", []byte(idData))
	if err != nil {
		log.Fatalf("Erro ao excluir tarefa: %v", err)
	}

	// Desserializando a resposta
	var response db.Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Fatalf("Erro ao desserializar a resposta: %v", err)
	}

	if response.Error != "" {
		log.Fatalf("Erro na resposta: %s", response.Error)
	}

	fmt.Println("Tarefa excluída com sucesso.")
}

func main() {
	skeleton := db.NewSkeleton()
	dispacher := db.NewDispacher(skeleton)

	// Listar todas as tarefas
	//insertTask(dispacher)
	//getAllTasks(dispacher)
	//getTaskByID(dispacher, "e076de67-d7f8-4ef3-8da5-4428389d49df")
	//deleteTask(dispacher, "2f68a6f9-8368-4e95-8fc3-c7d9a8566164")
	getAllTasks(dispacher)
}
