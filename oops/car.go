package main

import "fmt"

// struct type definition.
type Car struct {
	Make  string
	Model string
	Year  int
}

// This function shows the car.
// (c *Car) => Receiver: *Car pointer is passed for reduced memory usage.
func (c *Car) ShowCarDetails() {
	fmt.Println("\n***** Car Details *****")
	fmt.Printf("\n Make - %s\n Model - %s\n Year - %d\n", c.Make, c.Model, c.Year)
}

func main() {

	// Object instantiation.
	car1 := Car{Make: "Toyota", Model: "Camry", Year: 2012}
	car2 := Car{"Dodge", "Charger", 2023}

	// fmt.Printf("Car 1: %+v \n", car1)
	// fmt.Println("Car 2: ", car2)
	car1.ShowCarDetails()
	car2.ShowCarDetails()

	// Changing Properties/Fields.
	car1.Year = 2022
	car2.Model = "Challanger"

	// fmt.Printf("Car 1: %+v \n", car1)
	// fmt.Println("Car 2: ", car2)
	car1.ShowCarDetails()
	car2.ShowCarDetails()
}
