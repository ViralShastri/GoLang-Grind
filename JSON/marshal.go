package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Student Structure
type Student struct {
	Name  string
	Age   int
	Marks []int
}

func main() {

	student1 := Student{
		Name:  "Viral",
		Age:   22,
		Marks: []int{90, 99, 95},
	}

	student2 := Student{
		Name:  "Shastri",
		Age:   22,
		Marks: []int{90, 99, 95},
	}

	students := []Student{student1, student2}

	b, err := json.Marshal(students)

	if err != nil {
		fmt.Println("Error:", err)
	}

	// fmt.Println(string(b))
	os.Stdout.Write(b)

}
