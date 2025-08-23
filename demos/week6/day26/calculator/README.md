# Go Testing Demo

This directory contains example code demonstrating various testing techniques in Go:

1. **Unit Testing**
   - Table-driven tests
   - Helper functions
   - Test fixtures
   - Assertions and error reporting

2. **Benchmarking**
   - Performance measurement
   - Multiple benchmarks
   - ResetTimer usage

3. **Fuzzing**
   - Property-based testing
   - Seed corpus
   - Fuzz targets

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks
go test -bench=.

# Run fuzz tests
go test -fuzz=FuzzCalculator_Add

# Run tests with race detection
go test -race ./...
```

## Project Structure

```
calculator/
├── calculator.go         # Main implementation
├── calculator_test.go    # Unit tests and benchmarks
└── calculator_fuzz_test.go  # Fuzz tests
```

## Best Practices Demonstrated

1. Table-driven tests for multiple test cases
2. Helper functions for test setup
3. Clear test names and error messages
4. Benchmarks for performance-critical code
5. Fuzz testing for robustness
6. Proper error handling and boundary cases

## Concepts Covered

1. Basic arithmetic operations
2. Memory operations
3. Error handling
4. Test organization
5. Performance testing
6. Property-based testing
