package main

import "fmt"

// Person Type
type Person struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	var person = Person{
		firstname: "Viral",
		lastname:  "Shastri",
		age:       22,
	}

	fmt.Println(person)
	fmt.Println(person.firstname)
	fmt.Println(person.lastname)
	fmt.Println(person.age)

	var personMap = map[string]Person{
		person.lastname: person,
	}

	fmt.Println(personMap)

}
