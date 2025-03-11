package main

import "fmt"

func max[T int | float64 | string](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max(1, 2))
	fmt.Println(max(1.0, 2.0))
	fmt.Println(max("hello", "world"))
}
