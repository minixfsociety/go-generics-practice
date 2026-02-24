# Project 4: Data Filtering (Filter)

### Overview
This project demonstrates the power of combining Generics with Higher-Order Functions. We implemented a universal `Filter` function that can process any slice based on a custom logic (predicate).

### Key Concepts Learned:
* **Function as a Parameter**: Using `func(T) bool` as an argument allows the caller to define the filtering logic dynamically.
* **`any` Constraint**: Since the filtering logic is delegated to the callback function, the `Filter` function itself doesn't need to know anything about the type `T`.
* **Dynamic Slicing**: Building a new slice of type `T` using the `append` function within a generic context.
* **Anonymous Functions (Closures)**: Passing "on-the-fly" functions in `main` to handle specific filtering tasks (e.g., checking for even numbers or string length).

### Implementation Snippet
```go
func Filter[T any](input []T, check func(T) bool) []T {
    var result []T
    for _, v := range input {
        if check(v) {
            result = append(result, v)
        }
    }
    return result
}