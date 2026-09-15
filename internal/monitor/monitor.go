package monitor

import (
	"context"
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

type Monitor struct {
	client *http.Client
}

func NewMonitor() *Monitor {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	return &Monitor{
		client: &client,
	}
}

func (m *Monitor) CheckSite(ctx context.Context, url string, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	result := Result{URL: url}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err == nil {
		var resp *http.Response
		resp, err = m.client.Do(req)

		if err == nil {
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
		}
	}

	if err != nil {
		result.Status = "FAILED"
		result.Err = err
	}

	result.Duration = time.Since(start)
	ch <- result
}
