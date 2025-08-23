---
marp: true
theme: default
paginate: true
---

# Memory Management in Go
## Week 3 - Day 14

---

# Today's Topics

1. Pointers
2. Memory Allocation
3. Garbage Collection
4. Best Practices

---

# Pointers Basics

```go
x := 42
ptr := &x    // Get address
value := *ptr // Dereference

// Function parameters
func modify(x *int) {
    *x = 100
}
```

---

# Stack vs Heap

Stack:
- Function local variables
- Fixed size
- Fast allocation/deallocation

Heap:
- Dynamic allocation
- Variable size
- Garbage collected

---

# Memory Allocation

```go
// Stack allocation
x := 42

// Heap allocation
ptr := new(int)
*ptr = 42

// Make for slices/maps/channels
slice := make([]int, 0, 10)
```

---

# Garbage Collection

- Automatic memory management
- Mark and sweep algorithm
- Concurrent collection
- Stop-the-world phases

```go
runtime.GC()      // Force GC
runtime.ReadMemStats(&stats)
```

---

# Memory Leaks

Common causes:
1. Forgotten goroutines
2. Unclosed resources
3. Growing slices/maps
4. Global variables
5. Circular references

---

# Resource Management

```go
type Resource struct {
    // ...
}

func (r *Resource) Close() error {
    // Cleanup
}

defer resource.Close()
```

---

# Best Practices

1. Use pointers judiciously
2. Close resources properly
3. Monitor memory usage
4. Avoid memory leaks
5. Profile when needed

---

# Performance Tips

1. Minimize allocations
2. Reuse objects
3. Use sync.Pool
4. Right-size containers
5. Batch operations

---

# Exercise Time!

1. Practice with pointers
2. Monitor memory usage
3. Handle resources
4. Find memory leaks

---

# Questions?

Let's practice with hands-on examples!
