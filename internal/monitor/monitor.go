package monitor

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Err        error
	Duration   time.Duration
}

func CheckSite(url string, ch chan Result, wg *sync.WaitGroup) {
	start := time.Now()

	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {

		ch <- Result{
			URL:        url,
			StatusCode: -1,
			Err:        err,
			Duration:   time.Since(start),
		}
		return
	}
	defer resp.Body.Close()

	ch <- Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		Err:        nil,
		Duration:   time.Since(start),
	}
}
