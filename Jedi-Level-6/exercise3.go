package main

import "fmt"

func main() {
	defer log()
	fmt.Println("Process 1")
	fmt.Println("Process 2")
	fmt.Println("Process 3")
}

func log() {
	fmt.Println("Create Logs After Execution of The Process")
}
