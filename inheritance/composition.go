package main

import "fmt"

// Parent struct
type Person struct {
	name string
	age  int
}

// Child struct embedded the Parent struct
type Employee struct {
	Person
	jobTitle string
}

func main() {
	// Create an instance of Employee
	emp := Employee{
		Person:   Person{name: "John", age: 30},
		jobTitle: "Developer",
	}

	// Access fields from both Person and Employee
	fmt.Println("Name:", emp.name)
	fmt.Println("Age:", emp.age)
	fmt.Println("Job Title:", emp.jobTitle)
}
