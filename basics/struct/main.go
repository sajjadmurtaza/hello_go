package main

import "fmt"

type Student struct {
	Name    string
	Address string
	Age     int

	AdditionalInfo *AdditionalInfo // Optional field
}

type AdditionalInfo struct {
	Height int
	Weight int
}

func main() {
	// // ============= ####### =============
	// // =========== positional  ============
	// // ============= ####### =============

	student := Student{
		"GO",
		"2007 Google.com",
		18,
		nil,
	}

	fmt.Println("Student with Positional Arguments: \n", student)

	// // ============= ####### =============
	// // ============= Named  ==============
	// // ============= ####### =============

	goodStudent := Student{
		Name:    "Rust",
		Address: "Somewhere",
		Age:     20,
	}

	fmt.Println("Good Student with Named Arguments: \n", goodStudent)

	// // ============= ####### =============
	// // ============ Embedded  ============
	// // ============= ####### =============

	// Create an instance of AdditionalInfo
	additionalInfo := AdditionalInfo{
		Height: 180,
		Weight: 70,
	}

	fmt.Println("Additional Info: \n", additionalInfo)

	// // ============= ####### =============
	// // ========== with Methods  ==========
	// // ============= ####### =============

	student.PrintInfo()
	PrintInfo(student)

	fmt.Println("Before Update: \n", student)
	student.UpdateAge()
	fmt.Println("After Update: \n", student)
}

func (s Student) PrintInfo() {
	fmt.Println("Student Info: \n", s)
}

func PrintInfo(s Student) {
	fmt.Println("Student Info: \n", s)
}

func (s Student) UpdateAge() {
	s.Age++

	// remove pointer *
	// fmt.Println("Updated Age: \n", s.Age)
}
