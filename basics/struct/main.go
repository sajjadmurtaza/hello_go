// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import "fmt"

// ============= ####### =============
// ============== types =============
// ============= ####### =============

// Student represents a student record
type Student struct {
	Name    string
	Address string
	Age     int

	AdditionalInfo *AdditionalInfo // Optional field
}

// AdditionalInfo contains additional student information
type AdditionalInfo struct {
	Height int
	Weight int
}

// ============= ####### =============
// ============== funcs =============
// ============= ####### =============

// PrintInfo is a method on Student struct
func (s Student) PrintInfo() {
	fmt.Println("Student Info:", s)
}

// PrintInfo is a regular function that takes a Student
func PrintInfo(s Student) {
	fmt.Println("Student Info:", s)
}

// UpdateAge demonstrates value receiver (doesn't modify the original struct)
func (s Student) UpdateAge() {
	s.Age++
	// Note: This won't modify the original struct because it's a value receiver
}

// ============= ####### =============
// ============== main ==============
// ============= ####### =============

func main() {
	// ============= ####### =============
	// =========== positional  ============
	// ============= ####### =============

	// Create student using positional arguments
	student := Student{
		"GO",
		"2007 Google.com",
		18,
		nil,
	}
	fmt.Println("Student with Positional Arguments:", student)

	// ============= ####### =============
	// ============= Named  ==============
	// ============= ####### =============

	// Create student using named arguments
	goodStudent := Student{
		Name:    "Rust",
		Address: "Somewhere",
		Age:     20,
	}
	fmt.Println("Good Student with Named Arguments:", goodStudent)

	// ============= ####### =============
	// ============ Embedded  ============
	// ============= ####### =============

	// Create additional info
	additionalInfo := AdditionalInfo{
		Height: 180,
		Weight: 70,
	}
	fmt.Println("Additional Info:", additionalInfo)

	// ============= ####### =============
	// ========== with Methods  ==========
	// ============= ####### =============

	// Demonstrate method calls
	student.PrintInfo()
	PrintInfo(student)

	// Demonstrate value receiver behavior
	fmt.Println("Before Update:", student)
	student.UpdateAge()
	fmt.Println("After Update:", student)
	// Note: Age won't change because UpdateAge uses a value receiver
}
