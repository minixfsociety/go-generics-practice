package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println("Стек чисел:", intStack.elements)

	strStack := Stack[string]{}
	strStack.Push("Jump")
	strStack.Push("Run")
	fmt.Println("Стек строк:", strStack.elements)
}
