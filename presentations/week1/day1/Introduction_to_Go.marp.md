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
  .small-text {
    font-size: 0.8em;
  }
  code {
    background: #f0f0f0;
    border-radius: 4px;
    padding: 2px 4px;
  }
---

<!-- _class: lead -->
# Introduction to Programming with Go
## Day 1: Getting Started

---

# What is Programming? 🤔

Programming is giving instructions to a computer to:
- Solve problems
- Automate tasks
- Process data
- Create applications

Think of it as writing a recipe for the computer to follow!

---

# Programming Languages 🌍

<div class="columns">

**Types:**
- Low-level
- High-level
- Compiled
- Interpreted

**Examples:**
- Python
- Java
- C++
- JavaScript
- Go

</div>

---

# Why Go? 🚀

![bg right:40%](https://go.dev/images/gophers/pilot-bust.svg)

- Simple & Clear syntax
- Fast compilation
- Built-in concurrency
- Great standard library
- Growing ecosystem
- Strong community

---

# Go's History 📚

- Created by Google (2009)
- Creators:
  - Robert Griesemer
  - Rob Pike
  - Ken Thompson
- First release: 2012
- Current version: 1.21

---

# Key Features of Go ⭐

<div class="columns">

**Language Features:**
- Static typing
- Fast compilation
- Garbage collection
- Built-in testing
- Cross-platform

**Developer Benefits:**
- Easy to learn
- Clean syntax
- Great tooling
- Fast development
- Strong performance

</div>

---

# Setting Up Go 🛠️

1. Download Go:
   - Visit go.dev/dl
   - Choose your OS
   - Follow installer

2. Verify installation:
```bash
go version
```

3. Set up workspace:
```bash
mkdir my-go-projects
cd my-go-projects
```

---

# Your First Go Program 👋

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Save as `hello.go` and run:
```bash
go run hello.go
```

---

# Understanding the Code 🔍

```go
package main       // Program entry point
```

```go
import "fmt"      // Standard library package
```

```go
func main() {     // Main function
    // Your code here
}
```

---

# Go Tools 🔧

Essential commands:
```bash
go run    # Run program
go build  # Compile program
go fmt    # Format code
go test   # Run tests
```

---

# IDE Setup 💻

Recommended: VS Code
Extensions:
- Go extension
- Code Runner
- Go Test Explorer

Settings:
- Auto-formatting
- Auto-imports
- Code completion

---

# Best Practices 📝

1. Clear naming
2. Consistent formatting
3. Good documentation
4. Version control
5. Regular testing

---

<!-- _class: lead -->
# Let's Code! 💻

Time for hands-on practice...

---

# Resources 📚

- [Go Documentation](https://go.dev/doc)
- [Go Tour](https://tour.golang.org)
- [Go by Example](https://gobyexample.com)
- Course Repository
- Practice Exercises
