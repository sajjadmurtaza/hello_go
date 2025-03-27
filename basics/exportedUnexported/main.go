package main

import (
	"fmt"
	"hello_go/basics/exportedUnexported/models"
)

func main() {
	student := models.Student{
		Name:    "John",
		Address: "123 Main St",
		Age:     20,

		// unknown field Gender in struct literal
		// Gender:  "Male", // this will not work because Gender is unexported
	}

	fmt.Println("Student: ", student)
	fmt.Println("Student Name: ", student.GetName())

	// student.PrintName undefined (type models.Student has no field or method PrintName, but does have unexported method printName)
	// student.PrintName()

	fmt.Println(models.FirstName)
}
