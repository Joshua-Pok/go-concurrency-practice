package internal

import (
	"sync"
	"time"
)

type Log struct {
	Level     int
	Message   string
	Timestamp time.Time
}

type LogWriter interface {
	Write([]Log) error
}

type MockWriter struct {
	mu   sync.Mutex
	Logs []Log
}

func (m *MockWriter) Write(newLogs []Log) error {
	m.mu.Lock()

	defer m.mu.Unlock()

	m.Logs = append(m.Logs, newLogs...)
	return nil
}
