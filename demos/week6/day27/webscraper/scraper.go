package main

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

// Page represents a web page
type Page struct {
	URL     string
	Content string
}

// Result represents a processing result
type Result struct {
	Page  Page
	Error error
}

// Scraper handles concurrent web scraping
type Scraper struct {
	workers   int
	rateLimit int
	client    *http.Client
	semaphore chan struct{}
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
}

// NewScraper creates a new web scraper
func NewScraper(workers, rateLimit int, timeout time.Duration) *Scraper {
	ctx, cancel := context.WithCancel(context.Background())

	return &Scraper{
		workers:   workers,
		rateLimit: rateLimit,
		semaphore: make(chan struct{}, workers),
		ctx:       ctx,
		cancel:    cancel,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Scrape concurrently scrapes multiple URLs
func (s *Scraper) Scrape(urls []string) <-chan Result {
	// Create channels
	urlChan := s.generator(urls)

	// Fan out to workers
	workers := make([]<-chan Result, s.workers)
	for i := 0; i < s.workers; i++ {
		workers[i] = s.worker(urlChan)
	}

	// Fan in results
	return s.fanIn(workers...)
}

// generator creates a channel of URLs with rate limiting
func (s *Scraper) generator(urls []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		// Create rate limiter
		ticker := time.NewTicker(time.Second / time.Duration(s.rateLimit))
		defer ticker.Stop()

		for _, url := range urls {
			select {
			case <-s.ctx.Done():
				return
			case <-ticker.C:
				select {
				case out <- url:
				case <-s.ctx.Done():
					return
				}
			}
		}
	}()

	return out
}

// worker processes URLs and returns results
func (s *Scraper) worker(urls <-chan string) <-chan Result {
	results := make(chan Result)

	go func() {
		defer close(results)

		for url := range urls {
			// Acquire semaphore
			select {
			case s.semaphore <- struct{}{}:
			case <-s.ctx.Done():
				return
			}

			// Release semaphore when done
			defer func() { <-s.semaphore }()

			// Fetch page
			content, err := s.fetch(url)

			select {
			case results <- Result{
				Page: Page{
					URL:     url,
					Content: content,
				},
				Error: err,
			}:
			case <-s.ctx.Done():
				return
			}
		}
	}()

	return results
}

// fanIn combines multiple result channels into one
func (s *Scraper) fanIn(channels ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	out := make(chan Result)

	// Start goroutine for each input channel
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan Result) {
			defer wg.Done()
			for r := range c {
				select {
				case out <- r:
				case <-s.ctx.Done():
					return
				}
			}
		}(ch)
	}

	// Close output channel when all inputs are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// fetch retrieves content from a URL
func (s *Scraper) fetch(url string) (string, error) {
	req, err := http.NewRequestWithContext(s.ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Stop gracefully stops the scraper
func (s *Scraper) Stop() {
	s.cancel()
}

// Wait waits for all workers to finish
func (s *Scraper) Wait() {
	s.wg.Wait()
}
