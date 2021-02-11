package main

import "fmt"

func main() {
	person := struct {
		number int
	}{
		number: 15,
	}

	fmt.Println(person)
	fmt.Println(person.number)

}
