// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import "fmt"

// ============= ####### =============
// ============ structs =============
// ============= ####### =============

// Student represents a basic student structure
type Student struct {
	Name string
}

// ============= ####### =============
// ============== main ==============
// ============= ####### =============

func main() {
	// ============= ####### =============
	// ====== regular struct usage ======
	// ============= ####### =============

	// Create a new Student instance
	student := Student{
		Name: "GO",
	}
	fmt.Println("Student: ", student)

	// ============= ####### =============
	// ======== pointer creation ========
	// ============= ####### =============

	// Get the memory address of student
	studentPtr := &student // & operator returns the memory address
	fmt.Println("Student Pointer: ", studentPtr)

	// ============= ####### =============
	// ======= pointer dereference ======
	// ============= ####### =============

	// Access the value at the pointer's address
	studentPtrDereferenced := *studentPtr // * operator dereferences the pointer
	fmt.Println("Student Pointer Dereferenced: ", studentPtrDereferenced)
	fmt.Println("Student Pointer Dereferenced Name: ", studentPtrDereferenced.Name)
}

/* Additional Notes on Pointers in Structs:
   When defining structs, you can use pointers to make fields optional:

   type Student struct {
       Name string
       Runner *Runner  // Optional field using pointer
   }

   type Runner struct {
       Name string
   }
*/
