---
marp: true
theme: default
paginate: true
---

# Synchronization in Go
## Week 4 - Day 16

---

# Today's Topics

1. WaitGroups
2. Mutex
3. RWMutex
4. sync.Map
5. Best Practices

---

# WaitGroup

```go
var wg sync.WaitGroup

wg.Add(n)    // Add counter
wg.Done()    // Decrement counter
wg.Wait()    // Wait for zero

// Example
wg.Add(1)
go func() {
    defer wg.Done()
    // Do work
}()
```

---

# Mutex

```go
var mu sync.Mutex

mu.Lock()
// Critical section
mu.Unlock()

// Using defer
mu.Lock()
defer mu.Unlock()
// Critical section
```

---

# Thread-safe Counter

```go
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}
```

---

# RWMutex

- Multiple readers
- Single writer
- Better performance for read-heavy workloads

```go
var rwmu sync.RWMutex

// Read lock
rwmu.RLock()
defer rwmu.RUnlock()

// Write lock
rwmu.Lock()
defer rwmu.Unlock()
```

---

# sync.Map

- Thread-safe map
- No explicit locking
- Good for cache-like scenarios

```go
var m sync.Map

m.Store(key, value)
value, ok := m.Load(key)
m.Delete(key)

m.Range(func(key, value interface{}) bool {
    // Process entries
    return true
})
```

---

# Common Patterns

1. Protect shared resources
2. Coordinate goroutines
3. Handle race conditions
4. Implement thread-safe types

---

# Best Practices

1. Use minimal critical sections
2. Prefer RWMutex for read-heavy
3. Avoid nested locks
4. Document locking requirements
5. Use sync.Map appropriately

---

# Exercise Time!

1. Create thread-safe counter
2. Implement shared cache
3. Use WaitGroups
4. Handle race conditions

---

# Questions?

Let's practice with hands-on examples!
