---
marp: true
theme: default
paginate: true
---

# Data Structures in Go
## Week 3 - Day 13

---

# Today's Topics

1. Built-in Data Structures
2. Structs
3. Custom Data Structures
4. Best Practices

---

# Arrays and Slices

Arrays:
```go
var numbers [5]int
numbers := [5]int{1, 2, 3, 4, 5}
```

Slices:
```go
slice := []int{1, 2, 3}
slice = append(slice, 4)
```

---

# Maps

```go
// Declaration
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}

// Operations
ages["Carol"] = 35    // Add
delete(ages, "Bob")   // Remove
age, exists := ages["Alice"] // Check
```

---

# Structs

```go
type Person struct {
    Name    string
    Age     int
    Address Address
}

type Address struct {
    Street string
    City   string
}
```

---

# Working with Structs

```go
person := Person{
    Name: "John",
    Age:  30,
}

// Methods
func (p Person) FullName() string {
    return p.Name
}

// Pointer receivers
func (p *Person) SetAge(age int) {
    p.Age = age
}
```

---

# Custom Data Structures

1. Stacks
2. Queues
3. Linked Lists
4. Trees
5. Graphs

---

# Implementation Example

```go
type Stack struct {
    items []interface{}
}

func (s *Stack) Push(item interface{}) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
    // Implementation
}
```

---

# Thread Safety

```go
type SafeStack struct {
    lock  sync.Mutex
    items []interface{}
}

func (s *SafeStack) Push(item interface{}) {
    s.lock.Lock()
    defer s.lock.Unlock()
    // Implementation
}
```

---

# Best Practices

1. Choose appropriate structure
2. Consider thread safety
3. Use interfaces
4. Handle edge cases
5. Document behavior

---

# Exercise Time!

1. Implement custom stack
2. Create thread-safe queue
3. Work with nested structs
4. Practice with maps/slices

---

# Questions?

Let's practice with hands-on examples!
