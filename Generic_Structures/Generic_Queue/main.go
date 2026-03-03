package main

import "fmt"

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Enqueue(v T) {
	q.elements = append(q.elements, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	val := q.elements[0]
	q.elements = q.elements[1:]
	return val, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.elements[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) Clear() {
	q.elements = nil
	fmt.Println(" Очередь полностью очищена")
}

func (q *Queue[T]) Print() {
	if q.IsEmpty() {
		fmt.Println("Очередь пуста: []")
		return
	}
	fmt.Printf("Очередь (размер %d): %v\n", q.Size(), q.elements)
}

func main() {
	prices := &Queue[float64]{}
	prices.Enqueue(10.50)
	prices.Enqueue(25.99)
	prices.Enqueue(5.00)
	prices.Print()
	if p, ok := prices.Peek(); ok {
		fmt.Printf(" Первый в очереди: %.2f\n", p)
	}
	val, _ := prices.Dequeue()
	fmt.Printf(" Обслужили (Dequeue): %.2f\n", val)
	prices.Print()
	fmt.Printf(" Сейчас элементов: %d\n", prices.Size())
	prices.Clear()
	prices.Print()
}
