# Go Generics: Stack Implementation

This is a simple, type-safe **Stack** data structure built with Go Generics.

## What is a Stack?
A stack is a **LIFO** (Last In, First Out) data structure. Imagine a stack of plates: you can only add a new plate to the top, and you can only take the top one off.

## How to use
You can create a stack for any data type (int, string, custom structs) without rewriting the code.

```go
// Example for integers
intStack := Stack[int]{}
intStack.Push(10)

// Example for strings
strStack := Stack[string]{}
strStack.Push("Golang")