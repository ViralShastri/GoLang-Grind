package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	y := []string{"e", "d", "c", "b", "a"}

	fmt.Println("Unsorted:", x)
	sort.Ints(x)
	fmt.Println("Sorted:", x)

	fmt.Println("Unsorted:", y)
	sort.Strings(y)
	fmt.Println("Sorted:", y)

}
