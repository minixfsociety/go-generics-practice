# Generic Interfaces: Storage System 📦

This project demonstrates how to use **Generic Interfaces** in Go to create a flexible and pluggable storage system.

## The Problem
Normally, if you want a storage for `int` and another for `string`, you have to write two different interfaces or use `interface{}` (which is unsafe). 

## The Solution: Generic Interface
We created a single interface that defines **how** to work with any type `T`:

```go
type Repository[T any] interface {
    Save(T)     // Save something of type T
    Get() T     // Get something of type T
    Clear()     // Reset to zero value
}