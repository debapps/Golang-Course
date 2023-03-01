package main

import (
	"debmodules/input"
	"fmt"
	"math"
)

// Main function.
func main() {

	fmt.Println("********** Simple Calculator ************")

	num1 := input.GetFloat64Input("\nEnter Number 1: ")
	num2 := input.GetFloat64Input("\nEnter Number 2: ")
	operator := input.GetStringInput("\nChoose the operation (+, -, *, /): ")

	switch operator {
	case "+":
		fmt.Printf("\nThe Result - %v \n", addValues(num1, num2))
	case "-":
		fmt.Printf("\nThe Result - %v \n", minusValues(num1, num2))
	case "*":
		fmt.Printf("\nThe Result - %v \n", multiplyValues(num1, num2))
	case "/":
		fmt.Printf("\nThe Result - %v \n", devideValues(num1, num2))
	default:
		fmt.Println("\nWrong Choice !!")
	}
}

func addValues(num1, num2 float64) float64 {
	result := math.Round((num1+num2)*100) / 100
	return result
}

func minusValues(num1, num2 float64) float64 {
	result := math.Round((num1-num2)*100) / 100
	return result
}

func multiplyValues(num1, num2 float64) float64 {
	result := math.Round((num1*num2)*100) / 100
	return result
}

func devideValues(num1, num2 float64) float64 {
	result := math.Round((num1/num2)*100) / 100
	return result
}
