package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Println("Rectangle Area: ", rectangle.Area())

	circle := Circle{radius: 5}
	fmt.Println("Circle Area: ", circle.Area())
}
