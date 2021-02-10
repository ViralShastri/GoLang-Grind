package main

import "fmt"

func main() {
	name := "Shastr"
	switch {
	case name == "Viral":
		fmt.Print(name)
	case name == "Shastri":
		fmt.Print(name)
	default:
		fmt.Print("No name")
	}
}
