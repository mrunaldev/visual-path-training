---
marp: true
theme: default
paginate: true
---

# Interfaces in Go
## Week 4 - Day 17

---

# Today's Topics

1. Interface Basics
2. Interface Composition
3. Empty Interface
4. Type Assertions
5. Best Practices

---

# Interface Basics

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```

---

# Interface Implementation

- Implicit implementation
- No "implements" keyword
- Based on method signatures

```go
// Automatically implements Shape
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

---

# Interface Composition

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

---

# Empty Interface

```go
// Any type
interface{}

// Function accepting any type
func Process(v interface{}) {
    // Use type assertions or switches
}

// Slice of any type
var data []interface{}
```

---

# Type Assertions

```go
value, ok := interfaceValue.(ConcreteType)
if ok {
    // Type assertion succeeded
}

// Type switch
switch v := value.(type) {
case int:
    fmt.Printf("Integer: %d", v)
case string:
    fmt.Printf("String: %s", v)
default:
    fmt.Printf("Unknown type")
}
```

---

# Common Interfaces

- fmt.Stringer
- io.Reader
- io.Writer
- error
- sort.Interface

```go
type Stringer interface {
    String() string
}
```

---

# Best Practices

1. Keep interfaces small
2. Accept interfaces, return structs
3. Interface composition
4. Document behavior
5. Use standard interfaces

---

# Exercise Time!

1. Create custom interfaces
2. Implement multiple interfaces
3. Practice type assertions
4. Use interface composition

---

# Questions?

Let's practice with hands-on examples!
