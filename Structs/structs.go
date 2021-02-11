package main

import "fmt"

// Person Structure
type Person struct {
	firstname string
	lastname  string
}

// SecretAgent Structure
type SecretAgent struct {
	Person
	ltk bool
}

func main() {
	var person = Person{
		firstname: "Viral",
		lastname:  "Shastri",
	}

	fmt.Println(person)
	fmt.Println(person.firstname)
	fmt.Println(person.lastname)

	var secretAgent = SecretAgent{
		Person: Person{
			firstname: "Shastri",
			lastname:  "Viral",
		},
		ltk: true,
	}

	fmt.Println(secretAgent)
	fmt.Println(secretAgent.Person)
	fmt.Println(secretAgent.lastname)
	fmt.Println(secretAgent.firstname)
	fmt.Println(secretAgent.ltk)

	var student = struct {
		firstname string
		lastname  string
	}{
		firstname: "Viral",
		lastname:  "Shastri",
	}

	fmt.Println(student)
	fmt.Println(student.firstname)
	fmt.Println(student.lastname)

}
