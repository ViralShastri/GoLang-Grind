package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	// Delete 5
	x = append(x[:4], x[5:]...)
	fmt.Println(x)
}
