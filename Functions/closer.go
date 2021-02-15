package main

import "fmt"

func main() {
	// x := 10

	// fmt.Println(x)

	// x++

	// {
	// 	y := 20
	// 	fmt.Println(y)
	// 	fmt.Println(x)
	// }

	// fmt.Println(x)
	a := incrementor()
	b := incrementor()

	fmt.Println(a(), b())
	fmt.Println(a(), b())
	fmt.Println(a(), b())

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
