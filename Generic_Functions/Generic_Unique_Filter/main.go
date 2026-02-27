package main

import (
	"fmt"
	"time"
)

func Unique[T comparable](input []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 4, 4, 4, 1345, 444, 222, 222, 12, 12, 12, 33}
	words := []string{"go", "java", "go", "javascript", "c++", "java", "c#", "rust"}
	fmt.Println("Numbers:", Unique(numbers))
	time.Sleep(1 * time.Second)
	fmt.Println("Words:", Unique(words))
}
