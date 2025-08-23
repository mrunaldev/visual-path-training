---
marp: true
theme: default
paginate: true
---

# Concurrency Patterns in Go
## Week 3 - Day 15

---

# Today's Topics

1. Pipeline Pattern
2. Fan-out/Fan-in Pattern
3. Worker Pool Pattern
4. Best Practices

---

# Pipeline Pattern

```go
// Generator -> Processing -> Output

nums := Generator(1, 2, 3)
squares := Square(nums)
evens := Filter(squares)

for n := range evens {
    fmt.Println(n)
}
```

---

# Pipeline Components

1. Source/Generator
2. Processing stages
3. Sink/Consumer

```go
func Generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
```

---

# Fan-out/Fan-in Pattern

Fan-out:
- Distribute work to multiple goroutines
- Process in parallel

Fan-in:
- Combine multiple results
- Single output channel

---

# Fan-out Implementation

```go
func FanOut(input <-chan int, n int) []<-chan int {
    channels := make([]<-chan int, n)
    for i := 0; i < n; i++ {
        channels[i] = worker(input)
    }
    return channels
}
```

---

# Fan-in Implementation

```go
func FanIn(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    wg.Add(len(channels))
    for _, ch := range channels {
        go func(c <-chan int) {
            for n := range c {
                out <- n
            }
            wg.Done()
        }(ch)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

---

# Worker Pool Pattern

1. Fixed number of workers
2. Task queue
3. Result collection
4. Graceful shutdown

---

# Worker Pool Implementation

```go
type WorkPool struct {
    workers  int
    tasks    chan Task
    results  chan Result
    done     chan struct{}
}

func (wp *WorkPool) Start() {
    for i := 0; i < wp.workers; i++ {
        go wp.worker()
    }
}
```

---

# Best Practices

1. Clear ownership of channels
2. Proper error handling
3. Context for cancellation
4. Resource cleanup
5. Avoid goroutine leaks

---

# Exercise Time!

1. Build a pipeline
2. Implement fan-out/fan-in
3. Create worker pool
4. Handle cancellation

---

# Questions?

Let's practice with hands-on examples!
