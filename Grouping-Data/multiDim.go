package main

import "fmt"

func main() {
	x := []string{"Viral", "Shastri"}
	y := []string{"Shastri", "Viral"}

	z := [][]string{x, y}

	fmt.Print(z)

}
