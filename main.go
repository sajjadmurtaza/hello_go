package main

import (
	"fmt"
	"hello_go/math"
	"hello_go/pets"
	"time"
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

	fmt.Println(pet.Name)
	fmt.Println(pet.Breed)

	// fmt.Println(pet.Feed("Bread."))   // Calls the exported Feed method
	// fmt.Println(pet.GivenAttention()) // Calls the exported GivenAttention method

	sleepTime := time.Now()

	dog := pets.GuardDog("Rudi", "GermanShepherd", sleepTime)

	if dog.IsHungry() {
		fmt.Println(dog.Feed("KFC Chicken."))
	} else {
		fmt.Println("Your Dog is not Hungry.")

		time.Sleep(5 * time.Second)
		fmt.Print(dog.Feed("sushi."))
	}

	fmt.Println(dog.GivenAttention())

	// ============= #### =============
	// ============= MATH =============
	// ============= #### =============

	fmt.Printf("\n============= MATH =============\n\n")

	rectangle := math.Rectangle{
		Width:  1.2,
		Length: 2.2,
	}

	fmt.Printf("Rectangle area(Width: %.2f, Length: %.2f): %.2f\n", rectangle.Width, rectangle.Length, rectangle.GetArea())
}
