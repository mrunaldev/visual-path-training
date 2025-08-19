---
marp: true
theme: gaia
paginate: true
backgroundColor: '#FFFFFF'
style: |
  .columns {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 1rem;
  }
  section {
    font-size: 1.5em;
  }
  section.lead {
    text-align: center;
  }
  code {
    background: #f0f0f0;
    border-radius: 4px;
    padding: 2px 4px;
  }
---

<!-- _class: lead -->
# Basic Go Syntax
## Day 2: Program Structure & Data Types

---

# Program Structure 📝

Every Go program has:

```go
// 1. Package declaration
package main

// 2. Import statements
import "fmt"

// 3. Functions and code
func main() {
    // Your code here
}
```

---

# Basic Data Types in Go 🔤

<div class="columns">

**Numeric Types:**
- int
- float64
- uint
- byte

**Other Types:**
- string
- bool
- rune
- complex

</div>

---

# Integer Types 🔢

```go
var age int = 25        // Regular integer
var score int8 = 100    // 8-bit integer
var population int64    // 64-bit integer

// Unsigned integers (no negative values)
var distance uint = 150
var small uint8 = 255   // 0 to 255
```

---

# Floating Point Numbers 📊

```go
var height float64 = 1.75    // 64-bit float
var weight float32 = 68.5    // 32-bit float

// Scientific notation
temperature := 1.5e3         // 1500.0
```

---

# Boolean Type ✅

```go
var isValid bool = true
var isReady = false      // Type inference

// Boolean operations
isDone := true
isStarted := false
canProceed := isDone && !isStarted
```

---

# String Type 📝

```go
var name string = "Gopher"
message := "Hello, Go!"

// Multi-line strings
description := `
    This is a
    multi-line
    string
`
```

---

# Type Conversions 🔄

```go
// Integer to Float
height := float64(175)

// Float to Integer
weight := int(68.5)

// Number to String
age := 25
ageStr := fmt.Sprintf("%d", age)

// String to Number
"123" -> strconv.Atoi()
"123.45" -> strconv.ParseFloat()
```

---

# Formatting Output 🖨️

```go
name := "Go"
age := 25
height := 1.75

// Different formatting verbs
fmt.Printf("Name: %s\n", name)     // String
fmt.Printf("Age: %d\n", age)       // Integer
fmt.Printf("Height: %.2f\n", height) // Float
```

---

# Common Format Verbs 📋

<div class="columns">

**Basic Types:**
- %s - string
- %d - integer
- %f - float
- %t - boolean

**Special Formats:**
- %v - any value
- %T - type
- %% - % sign

</div>

---

# Type Checking 🔍

```go
var value interface{} = "hello"

// Type assertions
str, ok := value.(string)
if ok {
    fmt.Println("It's a string:", str)
}

// Type switch
switch v := value.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Integer:", v)
}
```

---

# Best Practices 👍

1. Use clear type names
2. Consistent formatting
3. Appropriate types for data
4. Document type choices
5. Handle conversions safely

---

# Common Pitfalls ⚠️

- Uninitialized variables
- Wrong type conversions
- Format string mismatches
- Integer overflow
- Float precision issues

---

<!-- _class: lead -->
# Let's Code! 💻

Time for hands-on practice...

---

# Resources 📚

- [Go Basic Types](https://go.dev/tour/basics/11)
- [Format Specifiers](https://pkg.go.dev/fmt)
- [Type Conversions](https://go.dev/tour/basics/13)
- Practice Exercises
- Documentation
