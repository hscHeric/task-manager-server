package db

import (
	"encoding/json"
	"fmt"
)

type Dispacher struct {
	skeleton *Skeleton
}

// Função para inicializar o despachante com o Skeleton
func NewDispacher(skeleton *Skeleton) *Dispacher {
	return &Dispacher{skeleton: skeleton}
}

// Método Invoke para despachar chamadas de método com base em um comando e payload
func (d *Dispacher) Invoke(command string, args []byte) ([]byte, error) {
	var response *Response

	// Seleciona o método a ser chamado com base no comando
	switch command {
	case "InsertTask":
		response = d.skeleton.InsertTask(args)
	case "GetAllTasks":
		response = d.skeleton.GetAllTasks()
	case "GetTaskByID":
		response = d.skeleton.GetTaskByID(args)
	case "DeleteTask":
		response = d.skeleton.DeleteTask(args)
	default:
		response = NewResponse(nil, fmt.Errorf("comando não reconhecido: %s", command))
	}

	// Serializa a resposta para JSON
	serializedResponse, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("falha ao serializar a resposta: %w", err)
	}

	return serializedResponse, nil
}
