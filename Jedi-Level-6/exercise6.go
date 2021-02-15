package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World!!")
	}()

	func(x int) {
		fmt.Println(x)
	}(19)
}
