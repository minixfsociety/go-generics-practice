package main

import (
	"fmt"
	"time"
)

func containts[T comparable](list []T, target T) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	names := []string{"Valentino", "Mario", "Artem", "Max"}
	numbers := []int{10, 2, 5, 6, 8, 2, 44, 88, 44, 22, 7657, 643, 11}
	fmt.Println(containts(names, "Bob"))
	time.Sleep(1 * time.Second)
	fmt.Println(containts(numbers, 10))
}
