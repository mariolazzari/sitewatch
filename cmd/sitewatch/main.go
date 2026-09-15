package main

import (
	"log"
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
	}

	for _, url := range urls {
		log.Printf("Checking %s...\n", url)

		statusCode, err := monitor.CheckSite(url)
		if err != nil {
			log.Printf("error checking %s: %s\n", url, err)
		}

		log.Printf("Status: %d\n", statusCode)
	}

	log.Printf("Total elapsed time: %s\n", time.Since(start))
}
