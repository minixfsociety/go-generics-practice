# Project 2: Comparisons (Max)

### Overview
This project explores the limitations of the `any` constraint and demonstrates how to handle types that support comparison operators like `>` and `<`. 

### Key Concepts Learned:
* **Ordered Constraints**: Not all types in Go support comparison (e.g., you cannot check if one `bool` is "greater" than another). We created an `Ordered` interface to limit `T` to types that support `>`.
* **String Comparison**: In Go, strings can be compared using relational operators. The comparison is done lexicographically (by byte value/alphabetical order).
* **Multiple Arguments of the same Generic Type**: Defining `(a, b T)` ensures that both arguments passed to the function are of the exact same type at runtime.

### Implementation Snippet
```go
type Ordered interface {
    int | float64 | string
}

func Max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}