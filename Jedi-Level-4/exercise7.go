package main

import "fmt"

func main() {
	students := [][]string{{"Viral", "Shastri"}, {"Shastri", "Viral"}}

	fmt.Println(students)

	for index, value := range students {
		fmt.Println(index)
		for i, v := range value {
			fmt.Printf("\t%d ", i)
			fmt.Printf("%s\n", v)
		}
	}

}
