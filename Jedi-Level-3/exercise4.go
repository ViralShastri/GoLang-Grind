package main

import "fmt"

func main() {
	var bd int = 1999
	for {
		if bd < 2021 {
			fmt.Println(bd)
			bd++
		} else {
			break
		}
	}
}
