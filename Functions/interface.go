package main

import "fmt"

// Human Interface
type Human interface {
	speak()
	eat()
}

// Person Structure
type Person struct {
	name string
	age  int
}

func main() {
	person := Person{
		name: "Viral Shastri",
		age:  22,
	}
	person.speak()
}

func (person Person) speak() {
	fmt.Println(person.name, person.age)
}

func eat() {
	fmt.Println("Eating...")
}
