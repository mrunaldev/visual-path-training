# Concurrent Web Scraper Demo

This demo implements a concurrent web scraper that showcases several advanced Go concurrency patterns.

## Patterns Demonstrated

### 1. Worker Pool Pattern
- Uses a configurable number of worker goroutines
- Implements a semaphore to control concurrent requests
- Demonstrates efficient resource utilization

### 2. Fan-out/Fan-in Pattern
- Fan-out: Distributes URLs to multiple worker goroutines
- Fan-in: Combines results from all workers into a single channel
- Shows how to manage multiple concurrent data streams

### 3. Rate Limiting
- Implements a configurable requests-per-second limit
- Uses time.Ticker for precise timing control
- Demonstrates how to prevent overwhelming target servers

### 4. Context Usage
- Implements graceful shutdown via context cancellation
- Shows proper timeout handling
- Demonstrates resource cleanup

## Code Structure

- `scraper.go`: Core scraper implementation with all concurrency patterns
- `main.go`: Example usage of the scraper

## Running the Demo

1. Build the project:
   ```bash
   go build
   ```

2. Run the executable:
   ```bash
   ./webscraper
   ```

3. To stop gracefully, press Ctrl+C

## Implementation Details

### Scraper Configuration
- Number of concurrent workers
- Rate limit (requests per second)
- Timeout per request

### Error Handling
- Proper error propagation through channels
- Context cancellation handling
- Graceful shutdown support

### Resource Management
- Controlled concurrency through semaphores
- Proper cleanup of resources
- Memory-efficient streaming of results

## Best Practices Demonstrated

1. Proper channel usage
2. Context-based cancellation
3. Graceful shutdown handling
4. Rate limiting implementation
5. Error propagation
6. Resource cleanup
