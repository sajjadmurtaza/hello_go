package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go!")

	// numbers := []int{1, 5, 3, 9, 2, 8}
	// max_number := find_max(numbers) // uncomment find_max
	// println(max_number)

	// ============= #### =============
	// ============= PETS =============
	// ============= #### =============

	// fmt.Printf("\n============= PETS =============\n\n")

	// pet := pets.Dog{
	// 	Name:  "Fast Doggy",
	// 	Breed: "Strange",
	// }

	// fmt.Println(pet.Name)
	// fmt.Println(pet.Breed)

	// // fmt.Println(pet.Feed("Bread."))   // Calls the exported Feed method
	// // fmt.Println(pet.GivenAttention()) // Calls the exported GivenAttention method

	// sleepTime := time.Now()

	// // dog := pets.GuardDog("Rudi", "GermanShepherd", sleepTime)
	// var animals []pets.Pet

	// animals = append(animals, pets.Bulldog("Rudi", "GermanShepherd", sleepTime))
	// animals = append(animals, pets.AngryCat("AngryCat", "LoLCat"))

	// for _, pet := range animals {
	// 	fmt.Printf("Type: %s \n\n", reflect.TypeOf(pet))

	// 	switch animal := pet.(type) {
	// 	case *pets.Dog:
	// 		fmt.Printf("Animal: Dog, Name: %s\n", animal.Name)
	// 	case *pets.Cat:
	// 		fmt.Printf("Animal: Cat, Name: %s\n", animal.Name)
	// 	default:
	// 		fmt.Printf("Stuipido Animalo Here...")
	// 	}

	// 	if pet.IsHungry() {
	// 		fmt.Println(pet.Feed("KFC Chicken."))
	// 	} else {
	// 		fmt.Println("Your Animal is not Hungry.")

	// 		time.Sleep(5 * time.Second)
	// 		fmt.Print(pet.Feed("sushi."))
	// 	}

	// 	fmt.Println(pet.GivenAttention())
	// 	fmt.Printf("\n")
	// }

	// ============= #### =============
	// ============= MATH =============
	// ============= #### =============

	// fmt.Printf("\n============= MATH =============\n\n")

	// rectangle := math.Rectangle{
	// 	Width:  1.2,
	// 	Length: 2.2,
	// }

	// fmt.Printf("Rectangle area(Width: %.2f, Length: %.2f): %.2f\n", rectangle.Width, rectangle.Length, rectangle.GetArea())
}

// func find_max(numbers []int) int {
// 	if len(numbers) == 0 {
// 		panic("Empty Array!")
// 	}

// 	max := numbers[0]

// 	for _, number := range numbers {
// 		if number > max {
// 			max = number
// 		}
// 	}

// 	return max
// }
