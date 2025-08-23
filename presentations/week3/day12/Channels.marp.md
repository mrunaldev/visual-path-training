---
marp: true
theme: default
paginate: true
---

# Channels in Go
## Week 3 - Day 12

---

# Today's Topics

1. Channel Basics
2. Channel Types
3. Select Statement
4. Communication Patterns

---

# What are Channels?

- Type-safe communication pipes
- Send and receive values
- Synchronize goroutines
- FIFO (First In, First Out)

```go
ch := make(chan int)    // Unbuffered
ch := make(chan int, 5) // Buffered
```

---

# Channel Operations

```go
// Send
ch <- value

// Receive
value := <-ch

// Close
close(ch)

// Range
for v := range ch {
    // Use v
}
```

---

# Channel Types

1. Unbuffered
   - Synchronous communication
   - Blocks until ready

2. Buffered
   - Asynchronous up to capacity
   - Blocks when full

```go
unbuffered := make(chan int)
buffered := make(chan int, 5)
```

---

# Channel Directions

```go
// Bidirectional
chan T

// Send-only
chan<- T

// Receive-only
<-chan T

func producer(out chan<- int)
func consumer(in <-chan int)
```

---

# The Select Statement

```go
select {
case v1 := <-ch1:
    // Use v1
case ch2 <- v2:
    // v2 sent
case <-time.After(timeout):
    // Handle timeout
default:
    // Optional default case
}
```

---

# Common Patterns

1. Generator Pattern
2. Fan-out/Fan-in
3. Pipeline Pattern
4. Worker Pools
5. Timeouts
6. Done Channels

---

# Best Practices

1. Document channel ownership
2. Be clear about buffering
3. Handle channel closure
4. Use select for timeouts
5. Clean up resources

---

# Exercise Time!

1. Create producer/consumer
2. Implement worker pool
3. Practice with select
4. Handle multiple channels

---

# Questions?

Let's dive into hands-on practice!
