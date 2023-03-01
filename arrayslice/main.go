package main

import (
	"fmt"
	"sort"
)

func main() {
	//Array
	fmt.Println("\nArray")

	var colors [3]string
	colors[0] = "Orange"
	colors[1] = "White"
	colors[2] = "Green"

	fmt.Println(colors)
	fmt.Println("Count of colors - ", len(colors))

	var numbers = [5]int{5, 3, 4, 10, 6}
	fmt.Println(numbers)

	//Slice.
	fmt.Println("\nSlice")
	// Declaration
	var colorList = []string{"Red", "Green", "Blue"}
	fmt.Println(colorList)
	fmt.Println("Count of colors - ", len(colorList))

	// Adding Elements
	fmt.Println("\nAdding more colors ..")
	colorList = append(colorList, "Orange", "Purple", "Yellow")
	fmt.Println(colorList)
	fmt.Println("Count of colors - ", len(colorList))

	// Removing Elements.
	var idx int = 2
	color := colorList[idx]
	fmt.Printf("\nDeleting color %v at %v ..\n", color, idx)
	colorList = append(colorList[:idx], colorList[idx+1:]...)
	fmt.Println(colorList)
	fmt.Println("Count of colors - ", len(colorList))

	// Sorting.
	fmt.Println("\nSorting.")
	numberList := []int{346, 23, 65, 567, 90, 14, 1987, 8}
	fmt.Println("Original Slice - ", numberList)
	sort.Ints(numberList)
	fmt.Println("Sorted Slice - ", numberList)

	// Declare Slice with make() function
	// make(type, length, capacity)
	fmt.Println("\nLenght & Capacity.")
	var list = make([]int, 2, 3)
	fmt.Printf("List - %v\nLength - %d | Capacity - %d\n", list, len(list), cap(list))

	list = append(list, 1)
	fmt.Printf("List - %v\nLength - %d | Capacity - %d\n", list, len(list), cap(list))

	list = append(list, 2)
	fmt.Printf("List - %v\nLength - %d | Capacity - %d\n", list, len(list), cap(list))

	list = append(list, 3)
	fmt.Printf("List - %v\nLength - %d | Capacity - %d\n", list, len(list), cap(list))

	list = append(list, 4)
	fmt.Printf("List - %v\nLength - %d | Capacity - %d\n", list, len(list), cap(list))

	list = append(list, 5)
	fmt.Printf("List - %v\nLength - %d | Capacity - %d\n", list, len(list), cap(list))

	list = append(list, 6)
	fmt.Printf("List - %v\nLength - %d | Capacity - %d\n", list, len(list), cap(list))

	list = append(list, 7)
	fmt.Printf("List - %v\nLength - %d | Capacity - %d\n", list, len(list), cap(list))

	// Slice/Array traversals.
	fmt.Println("\nSlice/Array traversals.")
	for idx, val := range list {
		fmt.Printf("Index - %v | Value - %v\n", idx, val)
	}

}
