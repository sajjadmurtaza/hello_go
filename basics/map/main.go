package main

import "fmt"

func main() {
	fmt.Println("=========================================")
	fmt.Println("============ Map ========================")

	students := map[string]int{
		"Jhon": 10,
		"Kim":  12,
	}

	for student, age := range students {
		fmt.Printf("Student: %s, Age: %d\n", student, age)
	}

	fmt.Println("=== Get single student's age ===========")
	fmt.Println("Student Jhon's age:", students["Jhon"])

	fmt.Println("=========================================")
	fmt.Println("==== Update Existing student's age ======")

	students["Kim"] = 13

	fmt.Println("Student Kim's age:", students["Kim"])

	fmt.Println("=========================================")
	fmt.Println("=== Add New student's age ===============")

	students["Habibi"] = 14

	fmt.Println("Student Habibi's age:", students["Habibi"])

	fmt.Println("=========================================")
	fmt.Println("==== Delete student ================")

	delete(students, "Kim")

	fmt.Println("Student Kim's age:", students["Kim"])
	fmt.Println(students)

}
