package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(sum(nums...))
	fmt.Println(evenSum(sum, nums...))

}

func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func evenSum(sum func(nums ...int) int, nums ...int) int {
	evenNums := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			evenNums = append(evenNums, v)
		}
	}
	return sum(evenNums...)
}
