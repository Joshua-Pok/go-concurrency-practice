package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"

	"github.com/Joshua-Pok/log-aggregator/internal"
)

func main() {

	wg := sync.WaitGroup{}

	writer := &internal.MockWriter{}

	ingester := internal.NewIngester(writer, 50)

	ingester.Start(10, &wg)

	for i := 0; i < 20; i++ {
		body := `{"level": 1, "message": "test log"}`
		req, _ := http.NewRequest("POST", "/ingest", strings.NewReader(body))
		rr := httptest.NewRecorder()
		ingester.HandleIngest(rr, req)
		fmt.Printf("sent log %d, status: %d\n", i, rr.Code)
	}

	ingester.Stop(&wg)

	fmt.Println(ingester.Stats())

}
