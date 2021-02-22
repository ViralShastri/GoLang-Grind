package main

import (
	"encoding/json"
	"fmt"
)

// Person Structure
type Person struct {
	Name string `json:"person_name"`
	Age  int    `json:"person_age"`
}

func main() {

	bytes := `[{"person_name":"Viral","person_age":22},{"person_name":"Shastri","person_age":50}]`

	var people []Person

	json.Unmarshal([]byte(bytes), &people)

	fmt.Println(people)

}
