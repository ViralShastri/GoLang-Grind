package main

import "fmt"

func main() {
	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", array)
	for index, value := range array {
		fmt.Printf("array[%d] = %d\n", index, value)
	}
}
