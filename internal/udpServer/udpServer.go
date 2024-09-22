package udpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/hscHeric/task-manager-server/internal/db"
	"github.com/hscHeric/task-manager-server/internal/message"
)

var ignoreFirstResponse = true //variavel global para o caso de teste de mensagem perdida

type CacheEntry struct {
	Addr *net.UDPAddr
	ID   int64
}

type UDPServer struct {
	conn  *net.UDPConn
	cache map[string]*CacheEntry
	mu    sync.Mutex
}

func NewUDPServer(addr string) (*UDPServer, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}

	return &UDPServer{
		conn:  conn,
		cache: make(map[string]*CacheEntry),
	}, nil
}

func (s *UDPServer) GetRequest() (*message.Message, *net.UDPAddr, error) {
	buffer := make([]byte, 1024)
	n, addr, err := s.conn.ReadFromUDP(buffer)
	if err != nil {
		return nil, nil, err
	}

	var msg message.Message
	err = json.Unmarshal(buffer[:n], &msg)
	if err != nil {
		return nil, nil, err
	}

	//ignorar primeira a resposta se for igual a -Teste mensagem perdida-.
	if ignoreFirstResponse {
		print("ok")
	}
	if msg.T == 3 && ignoreFirstResponse {
		log.Println("Ignorando a mensagem de teste de mensagem perdida.")
		ignoreFirstResponse = false // Ignora apenas a primeira vez
		return nil, nil, fmt.Errorf("simulação de perda de mensagem: %d", msg.T)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	cacheKey := addr.String()

	if entry, exists := s.cache[cacheKey]; exists {
		if entry.ID == msg.ID {
			log.Println("Mensagem duplicada recebida:", msg.ID, "de", addr)
			return nil, nil, fmt.Errorf("duplicated message: %d", msg.ID)
		}
	}

	s.cache[cacheKey] = &CacheEntry{
		Addr: addr,
		ID:   msg.ID,
	}

	log.Println("Mensagem recebida de:", addr)
	return &msg, addr, nil
}

func (s *UDPServer) SendResponse(addr *net.UDPAddr, response *message.Message) error {
	data, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("erro ao codificar a mensagem JSON: %v", err)
	}

	_, err = s.conn.WriteToUDP(data, addr)
	if err != nil {
		return err
	}

	log.Println("Resposta enviada para: ", addr.String())
	return nil
}

func (s *UDPServer) Close() error {
	return s.conn.Close()
}

func (s *UDPServer) handleRequest(m *message.Message) []byte {
	skeleton := db.NewSkeleton()
	if err := skeleton.GetInitError(); err != nil {
		fmt.Printf("Erro na inicialização: %v\n", err)
		return nil
	}

	dispatcher := db.NewDispacher(skeleton)
	param := db.NewParametros(m.ObjReference, m.MethodID, m.Args)
	response, err := dispatcher.Invoke(param)

	if err != nil {
		log.Printf("Erro ao obter resposta: %v", err)
		return nil
	}

	return response
}

func (s *UDPServer) Start() {
	log.Println("Servidor UDP iniciado e aguardando mensagens...")

	for {
		msg, addr, err := s.GetRequest()
		if err != nil {
			log.Printf("Erro ao receber mensagem: %v", err)
			continue
		}

		log.Println(msg.ObjReference)
		log.Printf("Mensagem recebida do cliente %s: %+v", addr.String(), msg)
		response := s.handleRequest(msg)

		if response == nil {
			shutdownMessage := &message.Message{

				ObjReference: msg.ObjReference,
				MethodID:     msg.MethodID,
				Args:         nil,
				T:            1,
				ID:           msg.ID,
				StatusCode:   500,
			}

			err = s.SendResponse(addr, shutdownMessage)
			if err != nil {
				log.Printf("Erro ao enviar mensagem de encerramento: %v", err)
			}

		}

		msg.Args = response
		msg.T = 1
		msg.StatusCode = 200
		err = s.SendResponse(addr, msg)
		if err != nil {
			log.Printf("Erro ao enviar resposta: %v", err)
		}
	}

}
