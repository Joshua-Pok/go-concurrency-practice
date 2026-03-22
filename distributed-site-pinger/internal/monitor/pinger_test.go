package monitor

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestSmoke(t *testing.T) {
	if false {
		t.Error("what even is the meaning of life")
	}
}

func TestPingSite(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }))

	defer srv.Close()

	res, err := pingSite(srv.URL)
	if err != nil {
		t.Errorf("Error occured during ping: %v", err)
	}

	if res.statusCode != 200 {
		t.Error("Uneexpected Status code:", res.statusCode)
	}

	if res.timeTaken < 0 {
		t.Error("no fucking way it took negative time")
	}
}

func TestMonitorConcurrency(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wg := sync.WaitGroup{}

		w.WriteHeader(http.StatusOK)
	}))

}
