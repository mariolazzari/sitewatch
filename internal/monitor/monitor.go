package monitor

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Status     string
	URL        string
	StatusCode int
	Err        error
	Duration   time.Duration
}

func CheckSite(url string, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	result := Result{URL: url}

	resp, err := http.Get(url)
	if err != nil {
		result.Status = "FAILED"
		result.Err = err
		result.Duration = time.Since(start)

		ch <- result
		return
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode < 300:
		result.Status = "OK"
	case resp.StatusCode < 400:
		result.Status = "WARNING"
	default:
		result.Status = "ERROR"
	}

	result.StatusCode = resp.StatusCode
	result.Duration = time.Since(start)

	ch <- result
}
