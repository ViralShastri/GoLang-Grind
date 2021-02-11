package main

import "fmt"

func main() {
	states := make([]string, 3, 5)

	states[0] = "Viral"
	states[1] = "Shastri"
	states[2] = "Manish"

	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Printf("states[%d] = %s\n", i, states[i])
	}

}
