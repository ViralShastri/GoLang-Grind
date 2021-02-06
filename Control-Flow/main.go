package main

import (
	"fmt"
)

func main() {
	loop()
	nestedLoop()
	whileFor()
	forRange()
	ascii()
	switchCondition()
}

func loop() {
	println()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func nestedLoop() {
	println()
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			print("* ")
		}
		println()
	}
}

func whileFor() {
	println()
	i := 1
	for {
		if i == 5 {
			print(i, " Continue ")
			i++
			continue
		}

		if i == 10 {
			print(i, " Break")
			break
		}
		print(i, " ")
		i++
	}
}

func forRange() {
	arr := [5]string{"A", "B", "C", "D", "E"}
	println()
	for _, val := range arr {
		print(val, " ")
	}
}

func ascii() {
	for i := 33; i <= 126; i++ {
		fmt.Printf("\n%d\t%c\t%#U", i, i, i)
	}
}

func switchCondition() {
	println()
	i := 5
	switch {
	case i == 5:
		print("Case 1")
		fallthrough
	case i == 6:
		print(" Case 2")
	case i == 7:
		print("Case 3")
	case i == 8:
		print("Case 4")
	}
}
