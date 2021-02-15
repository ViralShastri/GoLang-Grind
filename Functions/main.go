package main

import "fmt"

func main() {
	foo()

	x := make([]int, 5, 10)

	fmt.Println("Sum = ", sum(1, 2, 3))
	fmt.Println("Sum = ", sum(x...))

	fmt.Printf("len=%d cap=%d %v\n\n\n", len(x), cap(x), x)

	printType(x...)
	printType(1, 2, 3)

}

func foo() {
	fmt.Println("Hello World!")
}

func sum(x ...int) int {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

func printType(x ...int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}
