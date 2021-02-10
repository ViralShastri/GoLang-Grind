package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9, 10}
	fmt.Println(x)
	fmt.Println(y)
	z := append(x, y...)
	fmt.Println(z)
	fmt.Println(newPrint(1, 2, 3, 4, 5))
}

func newPrint(nums ...int) []int {
	slice := []int{}
	return append(slice, nums...)
}
