# Day 5: Loops in Go

## Overview

This session covers all aspects of loops in Go, including:
- For loops (standard, condition-only, infinite)
- Range loops
- Loop control statements (break, continue)
- Nested loops
- Loop patterns and best practices

## Demo Files

### 1. Basic Loops (01_basic_loops.go)
- Standard for loop
- Condition-only loops
- Infinite loops
- Break and continue
- Multiple variables in loop

### 2. Range-based Loops (02_range_loops.go)
- Iterating over slices
- Iterating over maps
- Iterating over strings
- Iterating over channels
- Range with blank identifier

### 3. Practical Examples (03_practical_examples.go)
- Pattern printing
- Data processing
- Search algorithms
- Collection manipulation
- Error retry loops

## Key Concepts

### Loop Types
1. Standard C-style for loop:
```go
for initialization; condition; post {
    // loop body
}
```

2. Condition-only loop (while-loop equivalent):
```go
for condition {
    // loop body
}
```

3. Infinite loop:
```go
for {
    // loop body
    // must use break to exit
}
```

4. Range-based loop:
```go
for index, value := range collection {
    // loop body
}
```

### Control Statements
- `break`: Exit the loop
- `continue`: Skip to next iteration
- `break label`: Exit to labeled statement
- `continue label`: Skip to next iteration of labeled loop

## Practice Exercises

1. **Number Series**
   - Print Fibonacci series
   - Generate prime numbers
   - Calculate factorial

2. **Pattern Printing**
   - Print pyramid patterns
   - Print number patterns
   - Print character patterns

3. **Data Processing**
   - Find maximum/minimum in slice
   - Calculate average of numbers
   - Filter elements based on condition

4. **Search and Sort**
   - Implement linear search
   - Find element in sorted array
   - Count occurrences

## Challenge Exercises

1. **Matrix Operations**
   - Matrix multiplication
   - Spiral matrix traversal
   - Matrix rotation

2. **Advanced Patterns**
   - Diamond pattern
   - Pascal's triangle
   - Alternating patterns

3. **Real-world Scenarios**
   - Retry mechanism with backoff
   - Batch processing
   - Concurrent iterations

## Common Mistakes to Avoid

1. **Infinite Loops**
   ```go
   // Incorrect
   for i := 0; i >= 0; i++ {
       // Will never end
   }

   // Correct
   for i := 0; i < maxValue; i++ {
       // Has termination condition
   }
   ```

2. **Variable Scope**
   ```go
   // Note: Loop variable scope
   for i := 0; i < 5; i++ {
       // i is only accessible here
   }
   // i is not accessible here
   ```

3. **Range Variable Reuse**
   ```go
   // Be careful with pointers in range loops
   for _, v := range slice {
       // v is reused in each iteration
   }
   ```

## Best Practices

1. **Loop Efficiency**
   - Avoid unnecessary computations in loop condition
   - Pre-calculate values when possible
   - Consider break conditions early

2. **Readability**
   - Use meaningful variable names
   - Keep loop body simple
   - Extract complex logic to functions

3. **Performance**
   - Use appropriate loop type for the task
   - Avoid unnecessary allocations
   - Consider batch processing for large datasets

## Resources

- [Go Tour - For loops](https://tour.golang.org/flowcontrol/1)
- [Effective Go - For loops](https://golang.org/doc/effective_go#for)
- [Go by Example - For loops](https://gobyexample.com/for)
- [Go by Example - Range](https://gobyexample.com/range)
