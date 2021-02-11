package main

import "fmt"

// Vehicle Type
type Vehicle struct {
	doors int
	color string
}

// Truck Type
type Truck struct {
	Vehicle
	fourWheel bool
}

// Sedan Type
type Sedan struct {
	Vehicle
	luxury bool
}

func main() {

	truck := Truck{
		Vehicle: Vehicle{
			doors: 4,
			color: "Black",
		},
		fourWheel: true,
	}

	fmt.Println(truck)
	fmt.Println(truck.Vehicle)
	fmt.Println(truck.doors)
	fmt.Println(truck.color)
	fmt.Println(truck.fourWheel)

	sedan := Sedan{
		Vehicle: Vehicle{
			doors: 2,
			color: "White",
		},
		luxury: true,
	}

	fmt.Println(sedan)
	fmt.Println(sedan.Vehicle)
	fmt.Println(sedan.doors)
	fmt.Println(sedan.color)
	fmt.Println(sedan.luxury)
}
