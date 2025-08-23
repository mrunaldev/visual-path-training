---
marp: true
theme: default
paginate: true
---

# Code Organization in Go
## Week 2 - Day 10

---

# Today's Topics

1. Packages
2. Imports
3. Multi-file Programs
4. Best Practices

---

# Packages

- Basic unit of code organization
- One directory = one package
- Package name = last element of import path
- `main` package is special

```go
package calculator

func Add(a, b float64) float64 {
    return a + b
}
```

---

# Package Visibility

- Uppercase = Exported (public)
- Lowercase = Unexported (private)

```go
// Visible outside package
func Calculate() {}

// Only visible within package
func calculate() {}
```

---

# Imports

```go
import (
    "fmt"               // Standard library
    "github.com/user/pkg"  // External package
    "myapp/pkg"           // Local package
)

// Dot import (avoid)
import . "fmt"

// Aliased import
import f "fmt"
```

---

# Package Organization

```
myapp/
  ├── main.go
  ├── config/
  │   └── config.go
  ├── models/
  │   └── user.go
  └── utils/
      └── helper.go
```

---

# Multi-file Programs

- Files in same package share scope
- No explicit import between files
- Order-independent compilation

```go
// user.go
package models

type User struct {...}

// address.go
package models

type Address struct {...}
```

---

# Testing Organization

- Test files end in _test.go
- Can be in same or different package
- Use `testing` package

```go
package calc_test

import "testing"

func TestAdd(t *testing.T) {...}
```

---

# Best Practices

1. One package per directory
2. Clear package names
3. Organize by functionality
4. Keep packages focused
5. Use meaningful imports

---

# Exercise Time!

1. Create a simple package
2. Write package tests
3. Organize multi-file program
4. Practice imports

---

# Questions?

Let's practice with hands-on examples!
