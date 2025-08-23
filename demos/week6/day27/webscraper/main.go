package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	// Example URLs to scrape
	urls := []string{
		"https://golang.org",
		"https://blog.golang.org",
		"https://pkg.go.dev",
		// Add more URLs as needed
	}

	// Create scraper with:
	// - 3 workers
	// - 2 requests per second rate limit
	// - 10 second timeout per request
	scraper := NewScraper(3, 2, 10*time.Second)

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start scraping in a separate goroutine
	results := scraper.Scrape(urls)

	// Process results
	go func() {
		for result := range results {
			if result.Error != nil {
				fmt.Printf("Error scraping %s: %v\n", result.Page.URL, result.Error)
				continue
			}

			// Print first 100 characters of content
			content := strings.TrimSpace(result.Page.Content)
			if len(content) > 100 {
				content = content[:100] + "..."
			}
			fmt.Printf("Successfully scraped %s: %s\n", result.Page.URL, content)
		}
	}()

	// Wait for interrupt signal
	<-sigChan
	fmt.Println("\nShutting down gracefully...")

	// Stop the scraper
	scraper.Stop()
	scraper.Wait()
	fmt.Println("Shutdown complete")
}
