# Day 4: Control Structures in Go
## Demo Code Guide

This guide walks through the demo programs for Day 4, focusing on control structures in Go.

## 1. If Statements (01_if_statements.go)

Demonstrates various forms of if statements in Go.

### Key Concepts:
- Basic if statements
- If-else statements
- If with initialization
- Nested if statements
- Complex conditions
- Function conditions

### Running the Program:
```bash
go run 01_if_statements.go
```

## 2. Switch Statements (02_switch_statements.go)

Shows different ways to use switch statements in Go.

### Key Concepts:
- Basic switch
- Multiple cases
- Fallthrough
- Type switch
- Switch with initialization
- Switch without expression

### Running the Program:
```bash
go run 02_switch_statements.go
```

## 3. Practical Examples (03_practical_examples.go)

Real-world examples using control structures.

### Key Concepts:
- User validation
- Role-based access control
- Settings management
- Temperature classification
- Password strength checker

### Running the Program:
```bash
go run 03_practical_examples.go
```

## Practice Exercises

1. **Grade Calculator**
```go
// Create a program that:
// - Takes a numerical score (0-100)
// - Returns letter grade (A, B, C, D, F)
// - Handles invalid inputs
```

2. **Login System**
```go
// Create a simple login system that:
// - Checks username and password
// - Has different access levels
// - Handles invalid attempts
```

3. **Weather Advisor**
```go
// Create a program that:
// - Takes temperature and conditions
// - Suggests activities
// - Provides weather warnings
```

## Challenge Exercises

1. **Advanced Calculator**
```go
// Create a calculator that:
// - Handles multiple operations
// - Uses switch for operations
// - Validates input
// - Handles errors
```

2. **Game State Machine**
```go
// Create a game state system:
// - Different states (menu, playing, paused)
// - State transitions
// - Invalid state handling
```

## Common Issues and Solutions

1. **Condition Complexity**
   ```go
   // Bad
   if a > b && b > c && c > d && d > e {
       // Hard to read
   }

   // Good
   isValid := a > b && b > c
   isInRange := c > d && d > e
   if isValid && isInRange {
       // Clearer
   }
   ```

2. **Switch Fallthrough**
   ```go
   // Be careful with fallthrough
   switch x {
   case 1:
       fmt.Println("1")
       fallthrough    // Will execute next case regardless
   case 2:
       fmt.Println("2")
   }
   ```

3. **Nested If Statements**
   ```go
   // Bad
   if a {
       if b {
           if c {
               // Too deep
           }
       }
   }

   // Good
   if !a || !b || !c {
       return
   }
   // Main logic here
   ```

## Best Practices

1. **Keep Conditions Simple**
   - Break down complex conditions
   - Use intermediate variables
   - Consider extracting functions

2. **Switch vs If-Else**
   - Use switch for multiple options
   - Use if-else for binary decisions
   - Consider readability

3. **Error Handling**
   - Check errors first
   - Return early
   - Provide clear messages

## Next Steps

1. Complete all exercises
2. Try challenge exercises
3. Review best practices
4. Read about loops (next topic)
5. Practice combining control structures

## Additional Resources

- [Go Tour - Flow Control](https://tour.golang.org/flowcontrol/1)
- [Effective Go - Control Structures](https://golang.org/doc/effective_go#control-structures)
- [Go by Example - If/Else](https://gobyexample.com/if-else)
- [Go by Example - Switch](https://gobyexample.com/switch)
