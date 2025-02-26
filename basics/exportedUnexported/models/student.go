package models

import "fmt"

type Student struct {
	Name    string
	Address string
	Age     int
	gender  string
}

func (s Student) GetName() string {
	s.printName()
	return s.Name
}

func (s Student) printName() {
	fmt.Println("Name (printName): ", s.Name)
}
