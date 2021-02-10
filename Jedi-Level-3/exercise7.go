package main

import "fmt"

func main() {
	name := "Viral Shastri"
	if name == "Viral Shastri" {
		fmt.Print("You are ", name)
	} else if name == "Viral" || name == "Shastri" {
		fmt.Print("You are ", name)
	} else {
		fmt.Print("You are not ", name)
	}
}
