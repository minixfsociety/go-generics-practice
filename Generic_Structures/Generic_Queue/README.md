# Generic Queue in Go 📦

This project implements a **Queue** data structure using Go Generics. 
A Queue follows the **FIFO** (First In, First Out) principle — the first element added to the queue is the first one to be removed.

## Features
- **Generic Type Support**: Works with `int`, `string`, `structs`, etc.
- **Pointer-based Methods**: Uses pointers to ensure the original queue is modified, not a copy.
- **Safe Operations**: Methods like `Dequeue` and `Peek` return a boolean to safely handle empty queues.

## Logic Overview
* **Enqueue**: Adds an element to the end of the slice using `append`.
* **Dequeue**: Retrieves the first element (index `0`) and "slices" it off using `elements[1:]`.
* **Peek**: Look at the first element without removing it.
* **Size/IsEmpty**: Helper methods to track the state of the queue.

## Example Usage

```go
q := &Queue[string]{}

q.Enqueue("First Task")
q.Enqueue("Second Task")

val, ok := q.Dequeue()
if ok {
    fmt.Println("Processed:", val) // Output: First Task
}