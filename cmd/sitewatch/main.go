package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/mariolazzari/sitewatch/internal/monitor"
)

func main() {
	start := time.Now()
	urls := []string{
		"https://example.com",
		"https://mariolazzari.it",
		"https://github.com",
		"https://this-domain-does-not-exist.example",
		"https://mariolazzari.it/missing",
	}
	resCh := make(chan monitor.Result)
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	monitor := monitor.NewMonitor()

	for _, url := range urls {
		wg.Add(1)
		log.Printf("Checking %s...\n", url)
		go monitor.CheckSite(ctx, url, resCh, &wg)
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for res := range resCh {
		log.Printf("Status: %v\n", res)
	}

	log.Printf("Total elapsed time: %s\n", time.Since(start))
}
