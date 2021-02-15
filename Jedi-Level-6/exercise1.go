package main

import "fmt"

func main() {

	x := foo()

	a, b := bar()

	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
}

func foo() int {
	return 19
}

func bar() (int, string) {
	return 19, "Viral"
}
