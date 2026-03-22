package internal

import (
	"fmt"
	"sync"
)

type SafeStore[T any] struct {
	mu   sync.RWMutex
	data map[string]T
}

func NewSafeStore[T any]() Cache[T] {
	return &SafeStore[T]{
		mu:   sync.RWMutex{},
		data: make(map[string]T),
	}
}

func (s *SafeStore[T]) Get(key string) (T, bool) {
	//must lock here also, read during concurrent write will still crash

	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	if !ok {
		var zero T
		return zero, false
	}

	return val, true

}

func (s *SafeStore[T]) Set(key string, val T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = val
	return nil

}

func (s *SafeStore[T]) Delete(key string) error {
	s.mu.Lock()

	defer s.mu.Unlock()

	if _, ok := s.data[key]; !ok {
		return fmt.Errorf("Key %q not found", key)
	}

	delete(s.data, key)
	return nil

}
