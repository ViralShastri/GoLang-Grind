package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%T\n", array)
	for index, value := range array {
		fmt.Printf("array[%d] = %d\n", index, value)
	}
}
