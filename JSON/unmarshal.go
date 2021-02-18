package main

import (
	"encoding/json"
	"fmt"
)

// Student Structure
type Student struct {
	Name  string
	Age   int
	Marks []int
}

func main() {
	var jsonBlob = []byte(`[
		{"Name":"Viral","Age":22,"Marks":[90,99,95]},
		{"Name":"Shastri","Age":22,"Marks":[90,99,95]}
	]`)

	var students []Student

	err := json.Unmarshal(jsonBlob, &students)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(students)

}
