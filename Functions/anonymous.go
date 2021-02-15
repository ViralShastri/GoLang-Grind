package main

import "fmt"

func main() {

	func() {
		fmt.Println("Anonymous Function")
	}()

	func(x int) {
		fmt.Println("Anonymous Function With Params", x)
	}(42)
}
