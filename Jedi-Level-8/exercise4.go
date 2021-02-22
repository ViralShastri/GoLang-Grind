package main

import (
	"fmt"
	"sort"
)

// User Structure
type User struct {
	name  string
	age   int
	marks []int
}

func main() {

	x := []int{5, 4, 3, 2, 1}
	y := []string{"e", "d", "c", "b", "a"}

	fmt.Println("Not Sorted:", x)
	sort.Ints(x)
	fmt.Println("Sorted:", x)

	fmt.Println("Not Sorted:", y)
	sort.Strings(y)
	fmt.Println("Sorted:", y)

}
