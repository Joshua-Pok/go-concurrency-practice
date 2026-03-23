package internal

import (
	"testing"
	"time"
)

func TestMockWriter(t *testing.T) {

	sliceOfLogs := []Log{
		{
			Level:     1,
			Message:   "Hello",
			Timestamp: time.Now(),
		},
		{
			Level:     1,
			Message:   "World",
			Timestamp: time.Now(),
		},
	}

	mw := MockWriter{}

	mw.Write(sliceOfLogs)

	if len(mw.Logs) != 2 {
		t.Error("What happened??")
	}

}
