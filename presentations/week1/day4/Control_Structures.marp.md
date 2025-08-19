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
# Control Structures in Go
## Day 4: Making Decisions in Code

---

# Control Flow Basics 🔄

Programs need to:
- Make decisions
- Execute different paths
- Handle conditions
- Process alternatives

Control structures help us do this!

---

# If Statement 🔍

Basic syntax:
```go
if condition {
    // code to execute if true
}
```

Example:
```go
age := 18
if age >= 18 {
    fmt.Println("You can vote!")
}
```

---

# If-Else Statement 🔀

```go
if condition {
    // code for true case
} else {
    // code for false case
}
```

Example:
```go
score := 75
if score >= 60 {
    fmt.Println("Pass")
} else {
    fmt.Println("Fail")
}
```

---

# If with Initialization 💫

Go special feature:
```go
if value := compute(); value > 10 {
    fmt.Println("High value:", value)
} else {
    fmt.Println("Low value:", value)
}
// value is not accessible here
```

---

# Multiple Conditions ➕

```go
if condition1 {
    // code for condition1
} else if condition2 {
    // code for condition2
} else if condition3 {
    // code for condition3
} else {
    // default code
}
```

---

# Nested If Statements 🎯

```go
if outerCondition {
    if innerCondition {
        // nested code
    }
}
```

Example:
```go
if age >= 18 {
    if hasID {
        fmt.Println("Can enter")
    }
}
```

---

# Switch Statement 🔄

Basic syntax:
```go
switch value {
case option1:
    // code for option1
case option2:
    // code for option2
default:
    // default code
}
```

---

# Switch with Multiple Cases 📋

```go
switch day {
case "Monday", "Tuesday", "Wednesday",
     "Thursday", "Friday":
    fmt.Println("Weekday")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Invalid day")
}
```

---

# Switch with Conditions 🎯

```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}
```

---

# Switch with Fallthrough 🔽

```go
switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
}
```

---

# Type Switch 🔄

```go
var i interface{} = 42

switch v := i.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
case bool:
    fmt.Printf("Boolean: %t\n", v)
default:
    fmt.Printf("Unknown type\n")
}
```

---

# Best Practices 👍

1. Keep conditions simple
2. Use clear boolean expressions
3. Avoid deep nesting
4. Consider switch over long if-else chains
5. Use meaningful case ordering

---

# Common Patterns 📝

```go
// Early return
if err != nil {
    return err
}

// Guard clause
if !isValid {
    return
}

// Default value
if value == "" {
    value = "default"
}
```

---

# Tips for Clean Code 🧹

1. Use positive conditions
2. Avoid complex nesting
3. Extract complex conditions
4. Use switch for multiple options
5. Keep blocks short and focused

---

<!-- _class: lead -->
# Let's Code! 💻

Time for hands-on practice...

---

# Resources 📚

- [Go Control Structures](https://tour.golang.org/flowcontrol/1)
- [Effective Go - Control Structures](https://golang.org/doc/effective_go#control-structures)
- Practice Exercises
- Documentation
