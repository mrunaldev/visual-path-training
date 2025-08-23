---
marp: true
theme: default
paginate: true
---

# Generics in Go
## Week 4 - Day 18

---

# Today's Topics

1. Generic Functions
2. Generic Types
3. Type Constraints
4. Practical Examples

---

# Generic Functions

```go
func Map[T, U any](items []T, f func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = f(item)
    }
    return result
}

// Usage
numbers := []int{1, 2, 3}
squares := Map(numbers, func(x int) int {
    return x * x
})
```

---

# Generic Types

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    // Implementation
}
```

---

# Type Constraints

```go
// Basic constraint
[T any]

// Comparable types
[K comparable]

// Custom constraint
type Number interface {
    int | float64 | float32
}

func Sum[T Number](values []T) T {
    // Implementation
}
```

---

# Type Parameters

```go
type Pair[K, V any] struct {
    Key   K
    Value V
}

func NewPair[K, V any](key K, value V) Pair[K, V] {
    return Pair[K, V]{key, value}
}
```

---

# Generic Interfaces

```go
type Container[T any] interface {
    Add(T)
    Get() T
    Contains(T) bool
}

type List[T any] struct {
    // Implementation
}
```

---

# Practical Examples

1. Generic Collections
2. Data Structures
3. Utility Functions
4. Result Types

```go
type Result[T any] struct {
    Value T
    Error error
}
```

---

# Best Practices

1. Use meaningful constraints
2. Keep type parameters clear
3. Document generic behavior
4. Consider performance
5. Test thoroughly

---

# Exercise Time!

1. Create generic stack
2. Implement generic map
3. Use type constraints
4. Build data structures

---

# Questions?

Let's practice with hands-on examples!
