package monitor

import (
	"context"
	"net/http"
	"sync"
	"time"
)

const MAX_WAIT_TIME = 30 * time.Second

type Result struct {
	statusCode int
	Header     http.Header
	timeTaken  time.Duration
	Error      error
}

func pingSite(URL string) (Result, error) {
	start := time.Now()

	res, err := http.Head(URL)
	if err != nil {
		return Result{}, err
	}

	defer res.Body.Close()

	elapsed := time.Since(start)

	return Result{
		statusCode: res.StatusCode,
		Header:     res.Header,
		timeTaken:  elapsed,
	}, nil
}

func Monitor(urls []string, limit int) (<-chan Result, error) {
	wg := sync.WaitGroup{}

	sem := make(chan struct{}, limit)

	results := make(chan Result)

	for i, url := range urls {
		wg.Add(1)

		sem <- struct{}{}
		go func(i int, url string) {
			defer wg.Done()
			defer func() { <-sem }()

			done := make(chan Result, 1)

			go func() {
				r, err := pingSite(url)
				if err != nil {
					r.Error = err
				}
				done <- r
			}()
			select {
			case r := <-done:
				results <- r
			case <-time.After(MAX_WAIT_TIME):
				results <- Result{Error: http.ErrHandlerTimeout}
			}

		}(i, url)

	}

	go func() {
		wg.Wait()
		close(results)

	}()

	return results, nil
}
