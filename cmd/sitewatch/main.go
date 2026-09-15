package main

import (
	"log"
	"time"

	"github.com/mariolazzari/sitewatch/internal/monitor"
)

func main() {
	start := time.Now()
	url := "https://example.com"

	log.Printf("Checking %s...\n", url)
	statusCode, err := monitor.CheckSite(url)
	if err != nil {
		log.Fatalf("error checking %s: %s\n", url, err)
	}

	log.Printf("Status: %d\n", statusCode)
	log.Printf("Total elapsed time: %s\n", time.Since(start))
}
