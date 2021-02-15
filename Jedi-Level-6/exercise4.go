package main

import "fmt"

// Person Structure
type Person struct {
	name string
	age  int
}

func main() {
	person := Person{
		name: "Viral",
		age:  22,
	}
	person.speak()
}

func (person Person) speak() {
	fmt.Println(person.name, person.age)
}
