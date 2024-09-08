package udpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/hscHeric/task-manager-server/internal/message"
)

type CacheEntry struct {
	Addr *net.UDPAddr
	ID   int
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
