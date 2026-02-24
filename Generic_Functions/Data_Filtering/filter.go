package main

import "fmt"

func Filter[T any](input []T, check func(T) bool) []T {
	var result []T
	for _, v := range input {
		if check(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := Filter(ints, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evenNumbers)
	names := []string{"Max", "Valentino", "Bob", "Artem"}
	longNames := Filter(names, func(s string) bool {
		return len(s) > 3
	})
	fmt.Println("Long names:", longNames)
}
