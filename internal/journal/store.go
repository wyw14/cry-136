package journal

import (
	"sync"

	"github.com/wyw14/cry-136/internal/model"
)

type Store struct {
	mu     sync.RWMutex
	events []model.Event
}

func NewStore() *Store { return &Store{} }

func (s *Store) Append(event model.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
}

func (s *Store) Events() []model.Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]model.Event(nil), s.events...)
}

func (s *Store) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.events)
}
