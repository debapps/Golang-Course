package main

import "fmt"

func main() {

	// Travase a slice using loop.
	var places = []string{"Kolkata", "Chennai", "Delhi", "Darjeeling", "Puri", "Simla", "Pune"}
	fmt.Println(places)

	fmt.Println("\nTravase places using loop.")
	// Method 1
	// for i := 0; i < len(places); i++ {
	// 	fmt.Println(places[i])
	// }

	// Method 2
	// for i := range places {
	// 	fmt.Println(places[i])
	// }

	// Method 3
	for _, place := range places {
		fmt.Println(place)
	}

	fmt.Println("\nBreak in loop.")
	for _, place := range places {
		fmt.Println(place)
		if place == "Darjeeling" {
			break
		}
	}

	fmt.Println("\nContinue in loop.")
	for _, place := range places {
		if place == "Darjeeling" {
			continue
		}
		fmt.Println(place)
	}

	fmt.Println("\nGoto in loop.")
	for _, place := range places {
		if place == "Darjeeling" {
			goto endOfLoop
		}
		fmt.Println(place)
	}

endOfLoop:
	fmt.Println("End of Loop!")
}
