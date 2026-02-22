package main

import (
	"fmt"
	"time"
)

type Ordered interface {
	int | float64 | string
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println("Winner between 'Golang' and 'C++' =", Max("C++", "Golang"))
	time.Sleep(1 * time.Second)
	fmt.Println("Winner between 457 and 649 =", Max(457, 649))
}
