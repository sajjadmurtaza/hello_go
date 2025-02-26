package main

import "fmt"

type Student struct {
	Name string
}

func main() {

	// // ============= ####### =============
	// // ============= regular  ============
	// // ============= ####### =============

	student := Student{
		Name: "GO",
	}

	fmt.Println("Student: ", student)

	// // ============= ####### =============
	// // ============= address  ============
	// // ============= ####### =============

	studentPtr := &student // memory address of a student

	// studentPtr is called a pointer because it points to the memory location of studentPtr
	fmt.Println("Student Pointer: ", studentPtr)

	// // ============= ####### =============
	// // ============= dereference  ============
	// // ============= ####### =============

	studentPtrDereferenced := *studentPtr //access the value stored at a student pointer's address.

	fmt.Println("Student Pointer Dereferenced: ", studentPtrDereferenced)
	fmt.Println("Student Pointer Dereferenced Name: ", studentPtrDereferenced.Name)

}

// type Student struct {
// 	Name string
//
//  we can use * point to make optional fields in struct
// Runner *Runner
// }

// type Runner struct {
// 	Name string
// }
