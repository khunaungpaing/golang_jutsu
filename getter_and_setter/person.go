package main

import "fmt"

type Person struct {
	name   string
	age    uint8
	gender string
}

func (p Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func main() {

	p := Person{"mgmg", 12, "male"}

	fmt.Println(p.GetName())

	p.SetName("Aung Aung")

	fmt.Println(p.GetName())

}
