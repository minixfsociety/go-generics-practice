Project 1: Basic Arithmetic (Sum)
Overview
This is the first project in the Go Generics learning path. It demonstrates how to create a universal sum function that works with multiple numeric types without code duplication.

Key Concepts Learned:
Type Parameters: Using the [T Numbers] syntax to define a generic placeholder for types.
Type Sets (Interfaces): Using the | (Union) operator inside an interface to restrict which types are allowed.
Zero Value Initialization: Declaring var sum T which automatically starts at the "zero value" of the specific type (e.g., 0 for integers, 0.0 for floats).
Type Inference: Observing how the Go compiler automatically detects the type of T from the function arguments, making explicit type calling like sum[float64](...) optional in most cases.
Implementation Snippet
Go
type Numbers interface {
    int64 | float64
}

func sum[T Numbers](numbers []T) T {
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum
}
How to Run
Bash
go run main.go