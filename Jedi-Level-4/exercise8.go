package main

import "fmt"

func main() {
	students := map[string]int{
		"Viral":   22,
		"Shastri": 50,
	}

	fmt.Println(students)

	for key, value := range students {
		fmt.Printf("Map students['%s'] = %d\n", key, value)
	}

}
