package main

import (
	"log"

	udpserver "github.com/hscHeric/task-manager-server/internal/udpServer"
)

func main() {

	// Inicializa o servidor UDP
	server, err := udpserver.NewUDPServer("localhost:12345")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor UDP: %v", err)
	}

	// Inicia o servidor
	go func() {
		defer server.Close() // Garante que o socket seja fechado ao final
		server.Start()
	}()

	// Aguarda um sinal para encerrar o servidor (simulação de trabalho)
	select {}
}
