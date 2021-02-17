package main

import "fmt"

// Person Structure
type Person struct {
	name string
}

func changeMe(person *Person) {
	// (*person).name = "Shastri"
	person.name = "Shastri"
}

func main() {
	person := Person{
		name: "Viral",
	}
	fmt.Println(person)
	changeMe(&person)
	fmt.Println(person)
}
