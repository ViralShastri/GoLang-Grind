package main

import "fmt"

func main() {

	bar := foo()
	x := bar()
	fmt.Println(x)

	// fmt.Println(foo()())
}

func foo() func() int {
	return func() int {
		return 19
	}
}
