---
marp: true
theme: default
class: invert
paginate: true
---

# Day 5: Loops in Go
Visual Path Go Training

---

# Overview

- Types of Loops in Go
- For Loop Variations
- Range Loops
- Control Statements
- Best Practices
- Common Patterns

---

# Types of Loops in Go

Go simplifies looping by using only the `for` keyword, but with different patterns:

1. Standard C-style for loop
2. Condition-only loop (while-loop equivalent)
3. Infinite loop
4. Range-based loop

---

# Standard For Loop

```go
for initialization; condition; post {
    // loop body
}
```

Example:
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

---

# Condition-only Loop (While Loop)

```go
for condition {
    // loop body
}
```

Example:
```go
sum := 1
for sum < 100 {
    sum *= 2
}
```

---

# Infinite Loop

```go
for {
    // loop body
    if condition {
        break
    }
}
```

Example:
```go
count := 0
for {
    fmt.Println(count)
    count++
    if count >= 5 {
        break
    }
}
```

---

# Range Loop

Used to iterate over elements in:
- Arrays/Slices
- Maps
- Strings
- Channels

```go
for index, value := range collection {
    // loop body
}
```

---

# Range Examples

```go
// Slice
numbers := []int{1, 2, 3}
for i, num := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", i, num)
}

// Map
ages := map[string]int{"Alice": 25, "Bob": 30}
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

---

# Control Statements

1. `break`
   - Exits the loop immediately

2. `continue`
   - Skips to next iteration

3. Labels
   - Control outer loops from inner loops

---

# Break and Continue Examples

```go
// Break
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Print(i, " ")
}

// Continue
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Print(i, " ")
}
```

---

# Labeled Loops

```go
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i*j > 4 {
                break outer
            }
            fmt.Printf("(%d,%d) ", i, j)
        }
    }
```

---

# Common Patterns

1. Loop with Multiple Variables
```go
for i, j := 0, 10; i < 5; i, j = i+1, j+2 {
    fmt.Printf("i: %d, j: %d\n", i, j)
}
```

2. Collection Processing
```go
for _, num := range numbers {
    sum += num
}
```

---

# Error Handling in Loops

```go
for attempt := 1; attempt <= maxRetries; attempt++ {
    err := someOperation()
    if err == nil {
        break
    }
    time.Sleep(time.Second * time.Duration(attempt))
}
```

---

# Best Practices

1. Choose the right loop type
   - Use range for collections
   - Use condition-only for while-like loops
   - Use infinite loop for event loops

2. Keep loops simple
   - Extract complex logic to functions
   - Avoid deep nesting

---

# Best Practices (cont.)

3. Performance considerations
   - Pre-calculate values when possible
   - Avoid unnecessary allocations
   - Use break early pattern

4. Readability
   - Use meaningful variable names
   - Keep indentation clean
   - Comment complex logic

---

# Common Mistakes

1. Infinite Loops
```go
// Wrong
for i := 0; i >= 0; i++ {
    // Never ends
}

// Right
for i := 0; i < maxValue; i++ {
    // Has end condition
}
```

---

# Common Mistakes (cont.)

2. Variable Scope
```go
// Range variable reuse
for _, v := range slice {
    go func() {
        fmt.Println(v) // Wrong: v is reused
    }()
}

// Correct
for _, v := range slice {
    v := v // Create new variable
    go func() {
        fmt.Println(v)
    }()
}
```

---

# Practical Examples

1. Data Processing
```go
func processNumbers(numbers []int) (min, max, sum int) {
    if len(numbers) == 0 {
        return
    }
    min, max = numbers[0], numbers[0]
    for _, n := range numbers {
        if n < min { min = n }
        if n > max { max = n }
        sum += n
    }
    return
}
```

---

# Practical Examples (cont.)

2. Pattern Printing
```go
func printPyramid(height int) {
    for i := 0; i < height; i++ {
        for j := 0; j < height-i-1; j++ {
            fmt.Print(" ")
        }
        for k := 0; k <= i*2; k++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}
```

---

# Exercise Time!

Try these exercises:
1. Print Fibonacci sequence
2. Find prime numbers up to N
3. Create a multiplication table
4. Implement bubble sort
5. Count word frequency in a text

---

# Key Takeaways

1. Go simplifies loops with a single `for` keyword
2. Range is powerful for collection iteration
3. Control flow statements add flexibility
4. Labels help manage nested loops
5. Choose the right loop type for your needs

---

# Questions?

- What's the difference between break and continue?
- When should you use a range loop?
- How do you handle infinite loops safely?
- What are best practices for loop performance?

---

# Additional Resources

- [Go Tour - For loops](https://tour.golang.org/flowcontrol/1)
- [Effective Go - For loops](https://golang.org/doc/effective_go#for)
- [Go by Example - Range](https://gobyexample.com/range)
- Practice exercises in `/demos/week1/day5/`
