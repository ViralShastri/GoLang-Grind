package main

import "fmt"

func main() {
	var a int
	var p *int

	a = 19
	p = &a

	fmt.Printf("%v = %T\n", a, a)
	fmt.Printf("%v = %T\n", &a, &a)
	fmt.Printf("%v = %T\n", p, p)
	fmt.Printf("%v = %T\n", *p, *p)
	// fmt.Println(*&a)
	// fmt.Println(*&p)
	// fmt.Println(&p)
	*p++

	fmt.Println(*p)
	fmt.Println(a)

	var b int = *p
	fmt.Println(b)

}
