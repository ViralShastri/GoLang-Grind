package main

import "fmt"

type Human interface {
	speak()
}

type Person struct {
	name string
	age  int
}

func saySomething(human Human) {
	human.speak()
}

func (person *Person) speak() {
	fmt.Println(person.name, person.age)
}

func main() {

	person := Person{"Viral", 22}

	saySomething(&person)
	// saySomething(person) // Has to pass poniter beacuse spaek has pointer type as a recivier

}
