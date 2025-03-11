package main

import "fmt"

type generic interface {
	int | float64 | string
}

func max[T generic](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T generic](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max(1, 2))
	fmt.Println(max(1.0, 2.0))
	fmt.Println(max("hello", "world"))

	fmt.Println(min(1, 2))
	fmt.Println(min(1.0, 2.0))
	fmt.Println(min("hello", "world"))
}
