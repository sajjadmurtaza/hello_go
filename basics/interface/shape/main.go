// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import (
	"fmt"
	"math"
)

// ============= ####### =============
// ========== interfaces ===========
// ============= ####### =============

// Shape interface defines the behavior for geometric shapes
type Shape interface {
	Area() float64
}

// ============= ####### =============
// ============ structs ============
// ============= ####### =============

// Circle represents a circular shape
type Circle struct {
	radius float64
}

// Rectangle represents a rectangular shape
type Rectangle struct {
	width, height float64
}

// ============= ####### =============
// ========= shape methods =========
// ============= ####### =============

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// ============= ####### =============
// ============== main =============
// ============= ####### =============

func main() {
	// Create and test a rectangle
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Println("Rectangle Area: ", rectangle.Area())

	// Create and test a circle
	circle := Circle{radius: 5}
	fmt.Println("Circle Area: ", circle.Area())
}
