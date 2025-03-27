// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import (
	"fmt"
)

// ============= ####### =============
// ========== interfaces ===========
// ============= ####### =============

// Animal defines the behavior expected from any animal
type Animal interface {
	Speak() string
}

// ============= ####### =============
// ======== animal structs =========
// ============= ####### =============

// Cat represents a feline animal
type Cat struct{}

// Speak returns the sound a cat makes
func (c Cat) Speak() string {
	return "Meow"
}

// Compile-time verification that Cat implements Animal
var _ Animal = Cat{}

// Dog represents a canine animal
type Dog struct{}

// Speak returns the sound a dog makes
func (d Dog) Speak() string {
	return "Woof"
}

// Parrot represents a bird animal
type Parrot struct{}

// Speak returns the sound a parrot makes
func (p Parrot) Speak() string {
	return "Squawk"
}

// Snake represents a reptile animal
type Snake struct{}

// Speak returns the sound a snake makes
func (s Snake) Speak() string {
	return "Hiss"
}

// ============= ####### =============
// ============== main =============
// ============= ####### =============

func main() {
	// Create instances of different animals
	cat := Cat{}
	dog := Dog{}
	parrot := Parrot{}
	snake := Snake{}

	// Demonstrate polymorphism through the Speak interface
	fmt.Println("Cat: ", cat.Speak())
	fmt.Println("Dog: ", dog.Speak())
	fmt.Println("Parrot: ", parrot.Speak())
	fmt.Println("Snake: ", snake.Speak())
}
