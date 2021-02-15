package main

import "fmt"

func main() {
	bar := foo()
	x := bar()
	fmt.Print(x)
	fmt.Print(foo()())
}

func foo() func() int {
	return func() int {
		return 19
	}
}
