package main

import (
	"debmodules/input"
	"fmt"
	"math"
)

func main() {
	fmt.Println("******* Circle Area and Perimeter ********")
	var radius float64 = input.GetFloat64Input("Enter Circle Radius: ")

	var perimeter float64 = 2 * math.Pi * radius
	var area float64 = math.Pi * radius * radius
	fmt.Printf("\nRadius - %.2f \nPerimeter - %.2f \nArea - %.2f \n", radius, perimeter, area)

}
