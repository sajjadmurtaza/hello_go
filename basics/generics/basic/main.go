// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import "fmt"

// ============= ####### =============
// ============== funcs =============
// ============= ####### =============

// max is a generic function that finds the maximum of two values
// It works with int, float64, and string types
// This demonstrates basic generic type constraints in Go
func max[T int | float64 | string](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// ============= ####### =============
// ============== main ==============
// ============= ####### =============

func main() {
	// ============= ####### =============
	// ======= generic function use ======
	// ============= ####### =============

	// Demonstrate max with different types
	fmt.Println(max(1, 2))             // integers
	fmt.Println(max(1.0, 2.0))         // float64
	fmt.Println(max("hello", "world")) // strings
}
