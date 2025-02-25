package main

import "fmt"

var (
	firsName, lastName string = "GO", "Lang" //unexported or private, accessible within the package

	Number int //exported or public, accessible outside the package
)

func main() {
	// // ============= ####### =============
	// // ============== var ================
	// // ============= ####### =============

	var c, python, java bool //default value false
	var i int                //default value 0
	var j, k int = 9, 99     //type inference

	fmt.Printf("Unexported firstName: %s, lastName: %s, Exported Number: %d \n", firsName, lastName, Number)

	fmt.Println(i, j, k, c, python, java)

	// // ============= ####### =============
	// // ===== := (short assignment) =======
	// // ============= ####### =============

	name := "GO"

	fmt.Println(name)

	// // ============= ####### =============
	// // ============= const  ==============
	// // ============= ####### =============

	const isGoodToLean = true

	fmt.Println(isGoodToLean)

	// isGoodToLean = false
	// cost are immutable values
	// that cannot be changed after they are set.
}
