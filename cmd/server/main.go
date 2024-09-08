package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hscHeric/task-manager-server/internal/message"
	udpserver "github.com/hscHeric/task-manager-server/internal/udpServer"
)

// Assuming this is your project path

const dbPath = "./tasks.db"

func main() {
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
}
