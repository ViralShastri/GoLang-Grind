package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person Structure
type Person struct {
	Name string `json:"person_name"`
	Age  int    `json:"person_age"`
}

func main() {

	person1 := Person{"Viral", 22}
	person2 := Person{"Shastri", 50}

	people := []Person{person1, person2}

	bytes, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(bytes)

}
