package hello

import "rsc.io/quote/v3"

// Hello World
func Hello() string {
	return quote.HelloV3()
}

// Proverb Function
func Proverb() string {
	return quote.Concurrency()
}
