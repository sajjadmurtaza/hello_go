// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import "fmt"

// ============= ####### =============
// ============== funcs =============
// ============= ####### =============

// printMap demonstrates how to iterate over a map
func printMap(students map[string]int) {
	for student, age := range students {
		fmt.Printf("Student: %s, Age: %d\n", student, age)
	}
}

// ============= ####### =============
// ============== main ==============
// ============= ####### =============

func main() {
	// ============= ####### =============
	// ============== init ==============
	// ============= ####### =============

	// Initialize a map with key-value pairs
	students := map[string]int{
		"Jhon": 10,
		"Kim":  12,
	}

	// ============= ####### =============
	// ============== print =============
	// ============= ####### =============

	fmt.Println("Initial students:")
	printMap(students)

	// ============= ####### =============
	// ============== get ==============
	// ============= ####### =============

	fmt.Println("\nGet single student's age:")
	fmt.Println("Student Jhon's age:", students["Jhon"])

	// ============= ####### =============
	// ============== update ============
	// ============= ####### =============

	fmt.Println("\nUpdate existing student's age:")
	students["Kim"] = 13
	fmt.Println("Student Kim's age:", students["Kim"])

	// ============= ####### =============
	// ============== add ==============
	// ============= ####### =============

	fmt.Println("\nAdd new student:")
	students["Habibi"] = 14
	fmt.Println("Student Habibi's age:", students["Habibi"])

	// ============= ####### =============
	// ============== delete ===========
	// ============= ####### =============

	fmt.Println("\nDelete student:")
	delete(students, "Kim")
	fmt.Println("After deleting Kim:")
	fmt.Println("Student Kim's age:", students["Kim"]) // Will print 0 (zero value)
	fmt.Println("Current students:", students)
}
