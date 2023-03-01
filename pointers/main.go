package main

import "fmt"

func main() {

	var ptr *int
	var value int = 14

	ptr = &value
	fmt.Printf("Variable Memory Address - %v", ptr)
	fmt.Printf("\nPointer Value - %v\nValue - %v", *ptr, value)

	*ptr = *ptr * 2

	fmt.Printf("\nPointer Value - %v\nValue - %v", *ptr, value)

	valueFloat := 3.14
	floatPtr := &valueFloat
	fmt.Printf("\nPointer Value - %v\nValue - %v", *floatPtr, valueFloat)

}
