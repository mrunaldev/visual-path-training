---
marp: true
theme: default
paginate: true
---

# Testing and Benchmarking in Go
## Advanced Go Programming - Day 26

---

# Overview

1. Unit Testing in Go
2. Table-Driven Tests
3. Benchmarking
4. Test Coverage
5. Best Practices

---

# Unit Testing in Go

- Tests live alongside your code in `*_test.go` files
- Uses the `testing` package
- Run with `go test`

```go
// math.go
func Add(a, b int) int {
    return a + b
}

// math_test.go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

---

# Table-Driven Tests

- Efficient way to test multiple cases
- Easy to add new test cases
- Clear pattern of inputs and expected outputs

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"mixed numbers", -2, 3, 1},
        {"zeros", 0, 0, 0},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := Add(tc.a, tc.b)
            if got != tc.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    tc.a, tc.b, got, tc.expected)
            }
        })
    }
}
```

---

# Test Fixtures and Helpers

```go
func setupTestCase(t *testing.T) func() {
    // Setup code
    t.Log("Setting up test case")
    
    return func() {
        // Teardown code
        t.Log("Tearing down test case")
    }
}

func TestWithFixture(t *testing.T) {
    teardown := setupTestCase(t)
    defer teardown()
    
    // Test code here
}
```

---

# Testing HTTP Handlers

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/hello", nil)
    w := httptest.NewRecorder()
    
    HelloHandler(w, req)
    
    resp := w.Result()
    body, _ := io.ReadAll(resp.Body)
    
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status OK; got %v", resp.Status)
    }
    
    if string(body) != "Hello, World!" {
        t.Errorf("Expected 'Hello, World!'; got %v", string(body))
    }
}
```

---

# Benchmarking

- Measures performance of your code
- Uses `testing.B` type
- Run with `go test -bench=.`

```go
func BenchmarkAdd(b *testing.B) {
    // Reset the timer to ignore setup time
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

---

# Benchmark Examples

```go
func BenchmarkStringConcat(b *testing.B) {
    var str string
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        str += "x"
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    var builder strings.Builder
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        builder.WriteString("x")
    }
}
```

---

# Test Coverage

- Measures how much code is tested
- Run with `go test -cover`
- Generate HTML report:
  ```bash
  go test -coverprofile=coverage.out
  go tool cover -html=coverage.out
  ```

---

# Best Practices

1. Test files should end with `_test.go`
2. Test functions should start with `Test`
3. Use table-driven tests for multiple cases
4. Test both success and failure cases
5. Use subtests for better organization
6. Keep tests readable and maintainable
7. Don't test unexported functions directly

---

# Test Organization

```
mypackage/
  ├── mycode.go
  ├── mycode_test.go
  ├── testdata/
  │   ├── input.json
  │   └── expected.json
  └── internal/
      └── testutil/
          └── helpers.go
```

---

# Mocking in Tests

```go
type DataStore interface {
    Get(id string) (string, error)
    Set(id, value string) error
}

type MockStore struct {
    data map[string]string
}

func (m *MockStore) Get(id string) (string, error) {
    if val, ok := m.data[id]; ok {
        return val, nil
    }
    return "", errors.New("not found")
}
```

---

# Testing Concurrent Code

```go
func TestConcurrent(t *testing.T) {
    const workers = 4
    const iterations = 1000
    
    var wg sync.WaitGroup
    wg.Add(workers)
    
    for i := 0; i < workers; i++ {
        go func() {
            defer wg.Done()
            for j := 0; j < iterations; j++ {
                // Test concurrent operations
            }
        }()
    }
    
    wg.Wait()
}
```

---

# Advanced Testing Topics

1. Race Detection
   ```bash
   go test -race
   ```

2. Fuzzing
   ```go
   func FuzzAdd(f *testing.F) {
       f.Add(1, 2)  // Seed corpus
       f.Fuzz(func(t *testing.T, a int, b int) {
           Add(a, b)
       })
   }
   ```

3. Integration Tests
4. Performance Testing

---

# Resources

1. Go Testing Package Documentation
2. Go Blog - Testing Articles
3. Example Test Patterns
4. Benchmarking Guidelines

---

# Questions?

1. How to write effective tests?
2. When to use table-driven tests?
3. How to measure test coverage?
4. Best practices for benchmarking?
