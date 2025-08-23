---
marp: true
theme: default
paginate: true
---

# Go Functions Basics
## Week 2 - Day 6

---

# What are Functions?

- Building blocks of Go programs
- Reusable pieces of code
- Perform specific tasks
- Can take inputs and return outputs

---

# Function Declaration

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

---

# Basic Function Example

```go
func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

// Usage
message := greet("Gopher")
fmt.Println(message) // Output: Hello, Gopher!
```

---

# Multiple Parameters

```go
func add(a, b int) int {
    return a + b
}

// Usage
sum := add(5, 3)
fmt.Printf("5 + 3 = %d\n", sum)
```

---

# Variadic Functions

- Take variable number of arguments
- Use ... syntax
- Arguments become a slice inside function

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

---

# Multiple Return Values

```go
func swap(x, y string) (string, string) {
    return y, x
}

// Usage
first, second := swap("hello", "world")
```

---

# Best Practices

1. Use meaningful function names
2. Keep functions focused and small
3. Document your functions
4. Return early for error conditions
5. Use named returns when appropriate

---

# Exercise Time!

Try creating functions that:
1. Calculate area of different shapes
2. Process strings (reverse, count chars)
3. Work with slices of numbers

---

# Questions?

Let's practice with some hands-on coding!
