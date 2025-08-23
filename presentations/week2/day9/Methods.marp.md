---
marp: true
theme: default
paginate: true
---

# Methods in Go
## Week 2 - Day 9

---

# Today's Topics

1. Method Declaration
2. Value vs Pointer Receivers
3. Method Sets
4. Best Practices

---

# Method Declaration

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Usage
rect := Rectangle{10, 5}
area := rect.Area()
```

---

# Value Receivers

- Receives a copy of the type
- Can't modify the original value
- Good for read-only operations
- Thread-safe

```go
func (c Counter) Count() int {
    return c.value
}
```

---

# Pointer Receivers

- Receives pointer to the type
- Can modify the original value
- More efficient for large structures
- Common in practice

```go
func (c *Counter) Increment() {
    c.value++
}
```

---

# When to Use Each

Value Receivers:
- Small structures
- Immutable operations
- No internal pointers

Pointer Receivers:
- Need to modify receiver
- Large structures
- Consistency with other methods

---

# Method Sets

- All methods declared on a type
- Determines interface implementation
- Influenced by receiver type

```go
type Stringer interface {
    String() string
}

func (p Person) String() string {
    return p.Name
}
```

---

# Method Chaining

```go
func (s String) ToUpper() String {
    // implementation
    return result
}

// Usage
result := str.ToUpper().Trim().Reverse()
```

---

# Best Practices

1. Be consistent with receiver types
2. Use pointer receivers for mutating methods
3. Keep methods focused and small
4. Document receiver behavior
5. Follow Go naming conventions

---

# Exercise Time!

1. Create types with methods
2. Practice with both receiver types
3. Implement method sets
4. Build method chains

---

# Questions?

Let's practice with hands-on examples!
