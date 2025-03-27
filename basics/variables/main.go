// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import "fmt"

// ============= ####### =============
// ============== vars ==============
// ============= ####### =============

// Package-level variables
var (
	// Private variables (unexported)
	firstName, lastName string = "GO", "Lang"

	// Public variable (exported)
	Number int
)

var age int

// ============= ####### =============
// ============== funcs =============
// ============= ####### =============

func main() {
	// ============= ####### =============
	// ============== var ==============
	// ============= ####### =============

	// Multiple variable declarations with default values
	var c, python, java bool // default value: false
	var i int                // default value: 0
	var j, k int = 9, 99     // type inference

	// Print default values
	fmt.Println("age:", age)
	fmt.Printf("Unexported firstName: %s, lastName: %s, Exported Number: %d\n",
		firstName, lastName, Number)
	fmt.Println(i, j, k, c, python, java)

	// ============= ####### =============
	// ===== := (short assignment) =======
	// ============= ####### =============

	// Short variable declaration
	name := "GO"
	fmt.Println(name)

	// ============= ####### =============
	// ============= const  ==============
	// ============= ####### =============

	// Constants are immutable values
	const isGoodToLearn = true
	fmt.Println(isGoodToLearn)

	// isGoodToLearn = false // This would cause a compilation error
	// Constants cannot be changed after they are set
}
