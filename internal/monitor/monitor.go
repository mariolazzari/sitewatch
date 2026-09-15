package monitor

import (
	"net/http"
)

type Result struct {
	URL        string
	StatusCode int
	Err        error
}

func CheckSite(url string, ch chan Result) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{
			URL:        url,
			StatusCode: -1,
			Err:        err,
		}
		return
	}
	defer resp.Body.Close()

	ch <- Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		Err:        nil,
	}
}
