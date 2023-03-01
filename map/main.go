package main

import "fmt"

func main() {
	// Map Declarations.
	// var techDict = map[string]string{
	// 	"HTTP": "HyperText Transfer Protocol",
	// 	"FTP":  "File Transfer Protocol",
	// 	"SMTP": "Simple Message Transfer Protocol",
	// 	"WWW":  "World Wide Web",
	// }
	var techDict = make(map[string]string)

	// Value Assignments.
	techDict["HTTP"] = "HyperText Transfer Protocol"
	techDict["FTP"] = "File Transfer Protocol"
	techDict["SMTP"] = "Simple Message Transfer Protocol"
	techDict["WWW"] = "World Wide Web"

	// Display Map
	// fmt.Println(techDict)
	// fmt.Println(techDict["SMTP"])

	for key, val := range techDict {
		fmt.Printf("\nKey - %v, Value - %v", key, val)
	}

	// Geting the value of the specific key.
	k := "TCP"
	if value, ok := techDict[k]; ok {
		fmt.Printf("\n\nThe value of key [%v] : %v", k, value)
	} else {
		fmt.Printf("\n\nThe key [%v] is not found!", k)

	}

	// Get the Keys in a slice.
	var keys []string

	for key := range techDict {
		keys = append(keys, key)
	}

	fmt.Println("\n\nKeys - ", keys)
}
