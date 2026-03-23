package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type MetricStore struct {
	totalReceived atomic.Uint64
	totalWritten  atomic.Uint64
}

type Ingester struct {
	logChan chan []Log
	writer  LogWriter
	metrics MetricStore
}

func NewIngester(w LogWriter, bufferSize int) *Ingester {
	return &Ingester{writer: w, logChan: make(chan []Log, bufferSize)}

}

func (i *Ingester) HandleIngest(w http.ResponseWriter, r *http.Request) {

	log := &Log{}

	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		fmt.Errorf("Error decoding log: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	i.metrics.totalReceived.Add(1)

	select {
	case i.logChan <- []Log{*log}: //success! we were able to process it
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "server too busy", http.StatusServiceUnavailable)

	}

}

func (i *Ingester) startWorker(wg *sync.WaitGroup) {
	defer wg.Done()

	batch := make([]Log, 0, 100)

	ticker := time.NewTicker(100 * time.Millisecond)

	defer ticker.Stop()

	for {
		select {
		case logs, ok := <-i.logChan: //channel is closed
			if !ok {
				if len(batch) > 0 {
					i.writer.Write(batch)
				}
				return
			}

			batch = append(batch, logs...)

			if len(batch) >= 100 { //batch is full
				i.writer.Write(batch)
				i.metrics.totalWritten.Add(uint64(len(batch)))
				batch = batch[:0] //clear batch
			}
		case <-ticker.C: //100ms passes
			if len(batch) > 0 {
				i.writer.Write(batch)
				i.metrics.totalWritten.Add(uint64(len(batch)))
				batch = batch[:0]
			}
		}
	}

}

func (i *Ingester) Start(workerCount int, wg *sync.WaitGroup) {
	for j := 0; j < workerCount; j++ {
		wg.Add(1)
		go i.startWorker(wg)
	}
}

func (i *Ingester) Stop(wg *sync.WaitGroup) {
	close(i.logChan) // stop producers from sending anymore shit
	wg.Wait()        //wait for workers to finish their shit
}

func (i *Ingester) Stats() (received, written uint64) {
	return i.metrics.totalReceived.Load(), i.metrics.totalWritten.Load()
}
