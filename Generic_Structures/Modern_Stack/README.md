# Modern Generic Stack in Go

This project is a high-performance, type-safe implementation of a **Stack** data structure using **Go Generics** and **Atomic Operations**.

## Key Features
- 🚀 **Generics**: Works with any data type (`int`, `string`, custom structs) while maintaining type safety.
- ⚡ **Atomic Operations**: Uses the modern `sync/atomic` (Go 1.19+) to track operations count safely across multiple goroutines.
- 🛠 **LIFO Logic**: Standard "Last In, First Out" implementation with `Push` and `Pop` methods.

## Structure
The stack uses a slice for storage and an atomic counter for statistics:
```go
type Stack[T any] struct {
    elements []T
    ops      atomic.Int64 // Modern atomic counter
}