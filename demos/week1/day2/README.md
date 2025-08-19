# Day 2: Basic Go Syntax
## Demo Code Guide

This guide explains the demo programs for Day 2, focusing on Go's basic syntax, data types, and formatting.

## 1. Data Types (01_data_types.go)

Demonstrates basic data types in Go.

### Key Concepts:
- Integer types (int, int8, uint64)
- Floating-point types (float32, float64)
- Boolean type
- String type
- Variable declaration
- Type inference

### Running the Program:
```bash
go run 01_data_types.go
```

## 2. Type Conversion (02_type_conversion.go)

Shows how to convert between different data types.

### Key Concepts:
- Basic type conversions
- String to number conversion
- Number to string conversion
- Error handling in conversions
- Using the strconv package

### Running the Program:
```bash
go run 02_type_conversion.go
```

## 3. Formatting (03_formatting.go)

Demonstrates different ways to format and print data.

### Key Concepts:
- fmt.Printf usage
- Format verbs
- Width and padding
- Number formats
- Type information
- String operations

### Running the Program:
```bash
go run 03_formatting.go
```

## Practice Exercises

1. **Data Types Practice**
   - Create variables of different types
   - Print their values and types
   - Try different numeric types
   - Experiment with type inference

2. **Type Conversion Practice**
   - Convert between numeric types
   - Handle string conversions
   - Practice error handling
   - Try edge cases

3. **Formatting Practice**
   - Create a formatted table
   - Use different number formats
   - Practice string formatting
   - Try custom padding

## Common Issues and Solutions

1. **Type Conversion Errors**
   ```go
   // Wrong
   myInt = myFloat           // Compilation error
   
   // Right
   myInt = int(myFloat)      // Explicit conversion
   ```

2. **Format String Issues**
   ```go
   // Wrong
   fmt.Printf("%d", "string")  // Runtime error
   
   // Right
   fmt.Printf("%s", "string")  // Correct format verb
   ```

3. **Number Precision**
   ```go
   // Be aware of precision
   fmt.Printf("%.2f", 3.14159)  // Prints 3.14
   ```

## Best Practices

1. **Variable Declaration**
   - Use short declaration (:=) when possible
   - Be explicit about types when needed
   - Use meaningful variable names

2. **Type Conversion**
   - Always handle conversion errors
   - Use appropriate types for data
   - Be careful with numeric conversions

3. **Formatting**
   - Use appropriate format verbs
   - Consider alignment and readability
   - Document complex format strings

## Challenge Exercises

1. **Calculator Program**
   - Accept two numbers as strings
   - Convert to appropriate types
   - Perform basic operations
   - Format output nicely

2. **Data Formatter**
   - Create a program that formats data as a table
   - Use padding and alignment
   - Include different data types
   - Handle errors gracefully

## Next Steps

1. Review all example code
2. Complete practice exercises
3. Try challenge exercises
4. Read about basic operators
5. Prepare for Day 3
