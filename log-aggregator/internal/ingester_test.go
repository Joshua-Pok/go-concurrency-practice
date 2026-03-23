package internal

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIngester(t *testing.T) {
	jsonString := `{"level": 1, "message": "test log"}`

	req, err := http.NewRequest("POST", "/ingest", strings.NewReader(jsonString))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	writer := &MockWriter{}

	ingester := NewIngester(writer)

	ingester.HandleIngest(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}
