package main

import "fmt"

func main() {
	var slice = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice = append(slice, 52)
	fmt.Println(slice) // [42 43 44 45 46 47 48 49 50 51 52]

	slice = append(slice, 53, 54, 55)
	fmt.Println(slice) // [42 43 44 45 46 47 48 49 50 51 52 53 54 55]

	slice = append(slice, []int{56, 57, 58, 59, 60}...)
	fmt.Println(slice) // [42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60]
}
