package monitor

import (
	"net/http"
	"sync"
	"time"
)

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

			result, err := pingSite(url)
			if err != nil {
				result.Error = err
			}
			results <- result

		}(i, url)

	}

	go func() {
		wg.Wait()
		close(results)

	}()

	return results, nil
}
