package main

import (
	"fmt"
	"time"
)

type Numbers interface {
	int64 | float64
}

func sum[T Numbers](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}
func main() {
	ints := []int64{1, 6, 8, 4, 6, 7, 5, 3, 5, 6, 8, 3}
	floats := []float64{1.5, 3.6, 7.4, 6.3, 7.3, 2.4, 2.1}
	fmt.Println("Ints:", sum(ints))
	time.Sleep(1 * time.Second)
	fmt.Println("Floats:", sum[float64](floats))
}
