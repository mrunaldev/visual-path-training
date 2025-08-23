---
marp: true
theme: default
paginate: true
---

# Advanced Concurrency Patterns in Go
## Advanced Go Programming - Day 27

---

# Overview

1. Worker Pools
2. Fan-out/Fan-in
3. Rate Limiting
4. Context Patterns
5. Pipeline Patterns
6. Best Practices

---

# Worker Pools

Pattern for concurrent task processing:
- Fixed number of workers
- Shared work queue
- Load balancing

```go
func worker(id int, jobs <-chan Task, results chan<- Result) {
    for job := range jobs {
        results <- process(job)
    }
}

// Create worker pool
jobs := make(chan Task, 100)
results := make(chan Result, 100)
for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
}
```

---

# Worker Pool Implementation

```go
type Pool struct {
    workers  int
    tasks    chan Task
    results  chan Result
    done     chan struct{}
}

func NewPool(workers int) *Pool {
    return &Pool{
        workers: workers,
        tasks:   make(chan Task),
        results: make(chan Result),
        done:    make(chan struct{}),
    }
}

func (p *Pool) Start() {
    for i := 0; i < p.workers; i++ {
        go p.worker(i)
    }
}
```

---

# Fan-out Pattern

Distributing work across multiple goroutines:

```go
func fanOut(input <-chan Task, workers int) []<-chan Result {
    outputs := make([]<-chan Result, workers)
    for i := 0; i < workers; i++ {
        outputs[i] = worker(input)
    }
    return outputs
}

func worker(input <-chan Task) <-chan Result {
    output := make(chan Result)
    go func() {
        defer close(output)
        for task := range input {
            output <- process(task)
        }
    }()
    return output
}
```

---

# Fan-in Pattern

Combining multiple channels into one:

```go
func fanIn(ctx context.Context, channels ...<-chan Result) <-chan Result {
    var wg sync.WaitGroup
    multiplexed := make(chan Result)

    // Start goroutine for each input channel
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan Result) {
            defer wg.Done()
            for {
                select {
                case r, ok := <-c:
                    if !ok {
                        return
                    }
                    multiplexed <- r
                case <-ctx.Done():
                    return
                }
            }
        }(ch)
    }

    // Close multiplexed channel when all inputs are done
    go func() {
        wg.Wait()
        close(multiplexed)
    }()

    return multiplexed
}
```

---

# Rate Limiting

Control the rate of operations:

```go
func rateLimiter(input <-chan Task, rps int) <-chan Task {
    limiter := time.Tick(time.Second / time.Duration(rps))
    output := make(chan Task)
    
    go func() {
        defer close(output)
        for task := range input {
            <-limiter    // Wait for tick
            output <- task
        }
    }()
    
    return output
}

// With bursting
func burstRateLimiter(input <-chan Task, rps, burst int) <-chan Task {
    limiter := rate.NewLimiter(rate.Limit(rps), burst)
    output := make(chan Task)
    
    go func() {
        defer close(output)
        for task := range input {
            limiter.Wait(context.Background())
            output <- task
        }
    }()
    
    return output
}
```

---

# Context Patterns

Managing goroutine lifecycles:

```go
func worker(ctx context.Context, tasks <-chan Task) {
    for {
        select {
        case task := <-tasks:
            process(task)
        case <-ctx.Done():
            return
        }
    }
}

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
go func() {
    time.Sleep(5 * time.Second)
    cancel()
}()
```

---

# Pipeline Pattern

Series of stages connected by channels:

```go
func pipeline(input <-chan Task) <-chan Result {
    // Stage 1: Validation
    validated := validate(input)
    
    // Stage 2: Processing
    processed := process(validated)
    
    // Stage 3: Enrichment
    enriched := enrich(processed)
    
    return enriched
}

func validate(in <-chan Task) <-chan Task {
    out := make(chan Task)
    go func() {
        defer close(out)
        for task := range in {
            if isValid(task) {
                out <- task
            }
        }
    }()
    return out
}
```

---

# Generator Pattern

Creating a stream of data:

```go
func generator(done <-chan struct{}) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := 0; ; i++ {
            select {
            case out <- i:
            case <-done:
                return
            }
        }
    }()
    return out
}

// Usage
done := make(chan struct{})
nums := generator(done)
for i := 0; i < 5; i++ {
    fmt.Println(<-nums)
}
close(done)
```

---

# Error Handling in Concurrent Code

```go
type Result struct {
    Value interface{}
    Error error
}

func worker(tasks <-chan Task) <-chan Result {
    results := make(chan Result)
    go func() {
        defer close(results)
        for task := range tasks {
            value, err := process(task)
            results <- Result{value, err}
        }
    }()
    return results
}

// Error handling
for result := range results {
    if result.Error != nil {
        log.Printf("Error: %v", result.Error)
        continue
    }
    // Use result.Value
}
```

---

# Graceful Shutdown

```go
type Server struct {
    workers  []*Worker
    tasks    chan Task
    shutdown chan struct{}
    wg       sync.WaitGroup
}

func (s *Server) Shutdown() {
    close(s.shutdown) // Signal workers to stop
    s.wg.Wait()      // Wait for workers to finish
}

func (s *Server) worker() {
    defer s.wg.Done()
    for {
        select {
        case task := <-s.tasks:
            process(task)
        case <-s.shutdown:
            return
        }
    }
}
```

---

# Resource Management

```go
// Semaphore pattern
type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
    return make(chan struct{}, n)
}

func (s Semaphore) Acquire() {
    s <- struct{}{}
}

func (s Semaphore) Release() {
    <-s
}

// Usage
sem := NewSemaphore(3)
sem.Acquire()
// Do work
sem.Release()
```

---

# Best Practices

1. Always use buffered channels appropriately
2. Always handle channel closure properly
3. Use select with default case when needed
4. Implement proper shutdown mechanisms
5. Handle errors gracefully
6. Use context for cancellation
7. Monitor goroutine leaks
8. Control concurrent resource usage

---

# Common Pitfalls

1. Goroutine leaks
2. Race conditions
3. Deadlocks
4. Channel closure issues
5. Resource exhaustion
6. Error propagation
7. Uncontrolled concurrency

---

# Debugging Tools

1. Race Detector
   ```bash
   go run -race main.go
   ```

2. pprof
   ```go
   import _ "net/http/pprof"
   ```

3. Trace Tool
   ```go
   trace.Start(os.Stdout)
   defer trace.Stop()
   ```

---

# Questions?

1. When to use different patterns?
2. How to handle errors in concurrent code?
3. Best practices for graceful shutdown?
4. How to debug concurrent programs?
