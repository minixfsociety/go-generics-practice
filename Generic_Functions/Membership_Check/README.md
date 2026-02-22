# Project 3: Membership Check (Contains)

### Overview
This project introduces the built-in `comparable` constraint in Go. It demonstrates how to check for the existence of an element in a slice of any compatible type.

### Key Concepts Learned:
* **`comparable` Constraint**: A built-in interface that strictly allows types supporting the equality operators (`==` and `!=`). This includes booleans, numbers, strings, and even pointers or structs (as long as their fields are also comparable).
* **Search Logic**: Implementing a standard linear search algorithm that works across different data types.
* **Return Timing**: Understanding that `return false` must happen only after the entire collection has been exhausted, while `return true` can happen as soon as a match is found.

### Implementation Snippet
```go
func Contains[T comparable](list []T, target T) bool {
    for _, v := range list {
        if v == target {
            return true
        }
    }
    return false
}