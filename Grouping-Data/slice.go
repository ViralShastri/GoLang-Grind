package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] = %d\n", i, slice[i])
	}
	fmt.Println()
	for i, v := range slice {
		fmt.Printf("slice[%d] = %d\n", i, v)
	}
}
