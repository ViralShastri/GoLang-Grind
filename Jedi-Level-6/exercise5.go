package main

import (
	"fmt"
	"math"
)

// Shape Interface
type Shape interface {
	area() float32
}

// Square Structure
type Square struct {
	length float32
}

// Circle Structure
type Circle struct {
	radius float32
}

func (circle Circle) area() float32 {
	area := math.Pi * (circle.radius * circle.radius)
	return area

}

func (square Square) area() float32 {
	area := square.length * square.length
	return area
}

func info(shape Shape) {
	fmt.Println(shape.area())
}

func main() {
	square := Square{
		length: 10,
	}
	circle := Circle{
		radius: 5,
	}
	info(square)
	info(circle)
}
