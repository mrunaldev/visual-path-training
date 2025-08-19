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
# Variables, Constants & Operators
## Day 3: Building Blocks of Go Programs

---

# Variable Declaration 📝

Three ways to declare variables:

```go
// 1. Full declaration
var age int = 25

// 2. Type inference
var name = "Gopher"

// 3. Short declaration (:=)
score := 95
```

---

# Variable Scope 🎯

```go
var globalVar = "I'm global"

func main() {
    localVar := "I'm local"
    
    if true {
        blockVar := "I'm block-scoped"
        // blockVar only exists here
    }
    // localVar exists here
}
// globalVar exists everywhere
```

---

# Constants 🔒

```go
// Single constant
const Pi = 3.14159

// Multiple constants
const (
    StatusOK = 200
    StatusNotFound = 404
    MaxAge = 100
)

// iota for enumeration
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
)
```

---

# Basic Operators ➗

<div class="columns">

**Arithmetic:**
- `+` (add)
- `-` (subtract)
- `*` (multiply)
- `/` (divide)
- `%` (remainder)

**Comparison:**
- `==` (equal)
- `!=` (not equal)
- `<` (less)
- `>` (greater)
- `<=`, `>=`

</div>

---

# Assignment Operators ✍️

```go
x := 10

x += 5   // x = x + 5
x -= 3   // x = x - 3
x *= 2   // x = x * 2
x /= 4   // x = x / 4
x %= 3   // x = x % 3

// Increment/Decrement
x++      // x = x + 1
x--      // x = x - 1
```

---

# Logical Operators 🔄

```go
a, b := true, false

// AND operator
result1 := a && b    // false

// OR operator
result2 := a || b    // true

// NOT operator
result3 := !a        // false

// Combined
result4 := (a && !b) // true
```

---

# Operator Precedence 📊

1. `()` Parentheses
2. `*`, `/`, `%`
3. `+`, `-`
4. `==`, `!=`, `<`, `<=`, `>`, `>=`
5. `&&` (AND)
6. `||` (OR)

```go
result := 5 + 3 * 2  // 11, not 16
result = (5 + 3) * 2 // 16
```

---

# Introduction to Arrays 📦

```go
// Fixed-size array
var numbers [5]int

// Initialize with values
colors := [3]string{"red", "green", "blue"}

// Array with implicit size
scores := [...]int{95, 89, 78, 92}

// Accessing elements
first := colors[0]    // "red"
colors[1] = "yellow" // Modify element
```

---

# Array Operations 🔄

```go
// Length of array
numbers := [5]int{1, 2, 3, 4, 5}
length := len(numbers)  // 5

// Iterate over array
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}

// Range-based loop
for index, value := range numbers {
    fmt.Printf("%d: %d\n", index, value)
}
```

---

# Multi-dimensional Arrays 🎲

```go
// 2D array
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Accessing elements
element := matrix[1][2]  // 6
```

---

# Common Practices 👍

1. Use short declaration when possible
2. Keep variables close to usage
3. Use meaningful names
4. Group related constants
5. Comment complex operations

---

# Naming Conventions 📝

```go
// Good variable names
userAge := 25
firstName := "John"
maxValue := 100

// Good constant names
const (
    MaxConnections = 100
    DefaultTimeout = 30
    StatusSuccess  = "OK"
)
```

---

# Identifiers Rules ⚡

- Must start with letter/underscore
- Can contain letters, digits, underscores
- Case-sensitive
- Cannot use keywords

```go
valid_name := 1
userName123 := "John"
_temp := "temp"
// 123name := "invalid"  // Not allowed
```

---

<!-- _class: lead -->
# Let's Code! 💻

Time for hands-on practice...

---

# Resources 📚

- [Go Variables](https://tour.golang.org/basics/8)
- [Go Operators](https://go.dev/ref/spec#Operators)
- [Go Arrays](https://go.dev/tour/moretypes/6)
- Practice Exercises
- Documentation
