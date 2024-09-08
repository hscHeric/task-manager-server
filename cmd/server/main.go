package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hscHeric/task-manager-server/internal/db"
	"github.com/hscHeric/task-manager-server/internal/task"
)

// Assuming this is your project path

const dbPath = "./tasks.db"

func insertTask(d *db.Dispacher, args []byte) {

	resp, err := d.Invoke("InsertTask", []byte(args))
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
	fmt.Printf("ID: %sTitle: %sDescription: %sDate: %s\n",
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
	/*
		idGen := message.NewIDGenerator()
		mm := message.NewMessage("ola", "ola", []byte("ola"), http.StatusCreated, http.StatusAccepted, idGen)
		m, _ := json.Marshal(mm)
		fmt.Println(string(m))

		server, err := udpserver.NewUDPServer(":1234")
		if err != nil {
			fmt.Println("Error ao criar o servidor", err)
			return
		}
		defer server.Close()

		for {
			msg, _, err := server.GetRequest()
			if err != nil {
				fmt.Println("Erro ao receber a mensagem:", err)
				continue
			}

			if msg != nil {
				fmt.Println("ok")
			}
		}
	*/

	skeleton := db.NewSkeleton()
	dispacher := db.NewDispacher(skeleton)
	//taskData := `{"title": "Nova Tarefa", "description": "Descrição da nova tarefa", "date": "2024-09-05T00:00:00Z"}`

	// Listar todas as tarefas
	//insertTask(dispacher, []byte(taskData))
	getAllTasks(dispacher)
	getTaskByID(dispacher, "e076de67-d7f8-4ef3-8da5-4428389d49df")
	deleteTask(dispacher, "fb72cf34-5937-4962-919e-1e28fc6c2796")
	getAllTasks(dispacher)
}
