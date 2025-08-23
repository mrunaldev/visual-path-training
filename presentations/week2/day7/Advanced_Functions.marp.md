---
marp: true
theme: default
paginate: true
---

# Advanced Functions in Go
## Week 2 - Day 7

---

# Today's Topics

1. Anonymous Functions
2. Multiple Return Values
3. defer keyword
4. init and main functions
5. Underscore Operator

---

# Anonymous Functions

```go
// Direct invocation
func() {
    fmt.Println("Hello from anonymous function!")
}()

// Assigned to variable
greeting := func(name string) string {
    return "Hello, " + name
}
```

---

# Closures

- Functions that capture their environment
- Access variables from outer scope
- Maintain their own state

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

---

# Multiple Return Values

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
```

---

# The defer Keyword

- Delays execution until surrounding function returns
- LIFO (Last In, First Out) order
- Common for cleanup operations

```go
func processFile() {
    f, err := os.Open("file.txt")
    if err != nil {
        return
    }
    defer f.Close()
    // Process file...
}
```

---

# init() Function

- Called before main()
- Multiple init functions allowed
- Used for initialization
- Cannot be called directly

```go
func init() {
    // Initialize resources...
}
```

---

# Underscore Operator

- Ignores values in assignments
- Used in imports for side effects
- Common in range loops

```go
for _, value := range slice {
    // Use value without index
}
```

---

# Best Practices

1. Use anonymous functions sparingly
2. Keep closures simple
3. Check errors from multiple returns
4. Use defer for cleanup
5. Don't abuse init()

---

# Exercise Time!

1. Create a closure-based counter
2. Implement error handling with multiple returns
3. Practice using defer
4. Experiment with init functions

---

# Questions?

Let's move to hands-on practice!
