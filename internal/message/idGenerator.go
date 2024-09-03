package message

import (
	"math"
	"sync"
)

type IDGenerator struct {
	mu sync.Mutex
	id int
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{id: math.MaxInt}
}

func (gen *IDGenerator) GetNextID() int {
	gen.mu.Lock()
	defer gen.mu.Unlock()
	gen.id--
	return gen.id
}
