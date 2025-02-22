package main

import (
	"fmt"
	"hello_go/math"
	"hello_go/pets"
)

func main() {
	fmt.Println("Hello Go!")

	// ============= #### =============
	// ============= PETS =============
	// ============= #### =============

	fmt.Printf("\n============= PETS =============\n\n")

	pet := pets.Dog{
		Name:  "Fast Doggy",
		Breed: "Strange",
	}

	fmt.Println(pet.Feed("Bread."))   // Calls the exported Feed method
	fmt.Println(pet.GivenAttention()) // Calls the exported GivenAttention method

	// ============= #### =============
	// ============= MATH =============
	// ============= #### =============

	fmt.Printf("\n============= MATH =============\n\n")

	rectangle := math.Rectangle{
		Width:  1.2,
		Length: 2.2,
	}

	fmt.Println(rectangle.GetArea())
}
