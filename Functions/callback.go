package main

import "fmt"

func main() {
	fmt.Println(evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func evenSum(sum func(x ...int) int, nums ...int) int {
	xi := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	return sum(xi...)
}
