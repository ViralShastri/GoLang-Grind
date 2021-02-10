package main

import "fmt"

func main() {
	x := make([]int, 2, 5)
	x[0] = 1
	x[1] = 2
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 3, 4, 5, 6)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
