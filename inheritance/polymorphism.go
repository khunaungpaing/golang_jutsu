package main

import "fmt"

// Animal interface
type Animal interface {
	Speak() string
}

// Dog type implementing the Animal interface
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat type implementing the Animal interface
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	// Create instances of Dog and Cat
	dog := Dog{}
	cat := Cat{}

	// Use the Animal interface to call Speak method
	animals := []Animal{dog, cat}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
