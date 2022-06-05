package tcp_server

import "sync"

// ServerStore interface: MemoryStore right now, possible Redis in the future
type ServerStore interface {
	Set(clientId string, task string)
	GetAndDelete(clientId string) (string, bool)
}

type MemoryStore struct {
	// TODO: clean up old sessions in goroutine (out of scope this task)
	mx sync.RWMutex
	m  map[string]string // clientId -> task (use sync.Map if you have more than 32 core and many clients)
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		m: make(map[string]string),
	}
}

func (s *MemoryStore) Set(clientId string, task string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.m[clientId] = task
}

func (s *MemoryStore) GetAndDelete(clientId string) (string, bool) {
	s.mx.Lock()
	task, ok := s.m[clientId]
	if ok {
		delete(s.m, clientId)
	}
	s.mx.Unlock()
	return task, ok
}
