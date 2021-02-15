package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}

	sum := foo(x...)

	fmt.Println(sum)

}

func foo(x ...int) int {
	var sum int

	for _, v := range x {
		sum += v
	}
	return sum

}
