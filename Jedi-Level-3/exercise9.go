package main

import "fmt"

func main() {
	var favSport string = "Cricket"

	switch favSport {
	case "Cricket":
		fmt.Print(favSport)
	case "VolleyBall":
		fmt.Print(favSport)
	case "FootBall":
		fmt.Print(favSport)
	default:
		fmt.Print("No Favorite")
	}
}
