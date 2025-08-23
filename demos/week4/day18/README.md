# Day 18: Generics in Go

This module covers Go's implementation of generics, which allows writing flexible and type-safe code that can work with different types while maintaining compile-time type safety.

## Topics Covered

1. Generic Types
   - Type parameters
   - Constraints
   - Type inference

2. Generic Functions
   - Basic syntax
   - Multiple type parameters
   - Constraints and interfaces

3. Data Structures
   - Generic Stack
   - Generic Set
   - Type-safe collections

4. Utilities
   - Generic math operations
   - Generic sorting
   - Collection operations (Map, Filter, Reduce)

## Examples

1. `01_generic_stack.go`
   - Generic stack implementation
   - Type-safe push/pop operations
   - Works with any type

2. `02_generic_utils.go`
   - Generic utility functions
   - Type constraints
   - Common operations

## Key Concepts

1. Type Parameters
   ```go
   func Map[T, U any](items []T, transform func(T) U) []U
   ```

2. Constraints
   ```go
   type Number interface {
       ~int | ~float64 | ~complex128
   }
   ```

3. Type Inference
   ```go
   result := Map([]int{1, 2, 3}, func(x int) string {
       return fmt.Sprint(x)
   })
   ```

## Best Practices

1. Use Generics When:
   - Implementing data structures
   - Writing algorithms
   - Creating reusable utilities

2. Avoid Generics When:
   - Interface satisfaction is sufficient
   - Type assertions work better
   - Code becomes less readable

3. Design Tips:
   - Keep constraints minimal
   - Use meaningful type parameter names
   - Document type requirements

## Exercises

1. Generic Binary Tree
   - Implement a generic binary tree
   - Add comparison operations
   - Traverse and search

2. Generic Sorting
   - Create a generic sorting function
   - Support custom comparators
   - Handle different types

## Resources

1. Documentation
   - [Go Generics Guide](https://go.dev/doc/tutorial/generics)
   - [Type Parameters Proposal](https://go.googlesource.com/proposal/+/master/design/43651-type-parameters.md)

2. Examples
   - [Standard Library Examples](https://pkg.go.dev/golang.org/x/exp/slices)
   - [Container Package](https://pkg.go.dev/golang.org/x/exp/container)

3. Articles
   - [Go Blog: Introduction to Generics](https://go.dev/blog/intro-generics)
   - [Practical Go Generics](https://go.dev/blog/when-generics)
