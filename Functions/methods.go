package main

import "fmt"

// Person Structure
type Person struct {
	name string
	age  int
}

func (person Person) speak() {
	fmt.Println("Hello, My name is", person.name, "& I'm", person.age, "Years Old.")
}

// SecretAgent Structure
type SecretAgent struct {
	Person
	ltk bool
}

func main() {
	secretAgent := SecretAgent{
		Person: Person{
			name: "Viral Shastri",
			age:  22,
		},
		ltk: true,
	}

	fmt.Println(secretAgent)
	secretAgent.speak()

}
