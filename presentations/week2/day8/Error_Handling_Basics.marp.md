---
marp: true
theme: default
paginate: true
---

# Error Handling in Go
## Week 2 - Day 8

---

# Today's Topics

1. Error Handling Patterns
2. panic
3. recover
4. Best Practices

---

# Error Handling Basics

- Errors are values in Go
- Multiple return values pattern
- error interface

```go
type error interface {
    Error() string
}
```

---

# Common Error Patterns

```go
// Return error as second value
func doSomething() (Result, error) {
    if something wrong {
        return nil, errors.New("something went wrong")
    }
    return result, nil
}

// Check error
if err != nil {
    // handle error
}
```

---

# Custom Error Types

```go
type CustomError struct {
    Code    int
    Message string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}
```

---

# panic

- For unrecoverable situations
- Stops normal execution
- Runs deferred functions
- Crashes program if not recovered

```go
func dangerous() {
    if badThing {
        panic("something terrible happened")
    }
}
```

---

# recover

- Used to handle panics
- Must be called in deferred function
- Returns panic value
- Returns nil if no panic

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Printf("Recovered: %v", r)
    }
}()
```

---

# Best Practices

1. Use errors for expected problems
2. Use panic for unrecoverable situations
3. Always recover in package APIs
4. Document error conditions
5. Keep error messages clear and actionable

---

# Exercise Time!

1. Create custom error types
2. Implement error handling
3. Practice with panic/recover
4. Build robust error handling patterns

---

# Questions?

Let's practice with some hands-on examples!
