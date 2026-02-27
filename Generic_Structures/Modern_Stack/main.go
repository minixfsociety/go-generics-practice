package main

import (
	"fmt"
	"sync/atomic"
)

type Stack[T any] struct {
	elements []T
	ops      atomic.Int64
}

func (s *Stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
	s.ops.Add(1)
}

func (s *Stack[T]) Pop() (T, bool) {
	s.ops.Add(1)
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	lastIndex := len(s.elements) - 1
	item := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return item, true
}

func main() {
	s := Stack[int]{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Стек после пушей:", s.elements)
	val, ok := s.Pop()
	if ok {
		fmt.Printf("Мы достали: %d\n", val)
	}
	fmt.Println("Стек после удаления последнего:", s.elements)
	fmt.Printf("Всего операций было: %d\n", s.ops.Load())
}
