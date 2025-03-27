// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import "fmt"

// ============= ####### =============
// =========== interfaces ===========
// ============= ####### =============

// generic is an interface that constrains types to those that support comparison
// It allows int, float64, and string types
type generic interface {
	int | float64 | string
}

// ============= ####### =============
// ============== funcs =============
// ============= ####### =============

// max returns the larger of two values
// Uses the generic interface constraint to work with multiple types
func max[T generic](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// min returns the smaller of two values
// Uses the generic interface constraint to work with multiple types
func min[T generic](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

// ============= ####### =============
// ============== main ==============
// ============= ####### =============

func main() {
	// ============= ####### =============
	// ======= max function tests =======
	// ============= ####### =============

	// Test max with different types
	fmt.Println(max(1, 2))             // integers
	fmt.Println(max(1.0, 2.0))         // float64
	fmt.Println(max("hello", "world")) // strings

	// ============= ####### =============
	// ======= min function tests =======
	// ============= ####### =============

	// Test min with different types
	fmt.Println(min(1, 2))             // integers
	fmt.Println(min(1.0, 2.0))         // float64
	fmt.Println(min("hello", "world")) // strings
}
