package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

var _ Animal = Cat{}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

type Parrot struct{}

func (p Parrot) Speak() string {
	return "Squawk"
}

type Snake struct{}

func (s Snake) Speak() string {
	return "Hiss"
}

func main() {
	cat := Cat{}
	dog := Dog{}
	parrot := Parrot{}
	snake := Snake{}

	fmt.Println("Cat: ", cat.Speak())
	fmt.Println("Dog: ", dog.Speak())
	fmt.Println("Parrot: ", parrot.Speak())
	fmt.Println("Snake: ", snake.Speak())
}
