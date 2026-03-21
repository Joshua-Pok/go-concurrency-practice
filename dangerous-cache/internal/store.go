package main

import "fmt"

type Cache[T any] interface {
	Get(key string) (T, bool)
	Set(key string, value T) error
	Delete(key string) error
}

type Store[T any] struct {
	data map[string]T
}

func NewStore[T any]() Cache[T] {
	return &Store[T]{
		data: make(map[string]T),
	}
}

func (s *Store[T]) Get(key string) (T, bool) {
	val, ok := s.data[key]
	if !ok {
		var zero T //initialize a zero value
		return zero, false
	}

	return val, true
}

func (s *Store[T]) Set(key string, value T) error {
	s.data[key] = value
	return nil
}

func (s *Store[T]) Delete(key string) error {
	if _, ok := s.data[key]; !ok {
		return fmt.Errorf("key %q not found", key)
	}

	delete(s.data, key)
	return nil
}
