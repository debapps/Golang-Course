package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("\n***** Text File Handling *****")
	var poem string = `
		On a dark desert highway,
		Cool wind in my hair.
		Warm smell of collitus,
		Rising up through the air...

		Up ahead in the distance,
		I saw a smearing light,
		My head go heavy and sight go deem,
		I had to stop for the night.

		Welcome to the Hotel California!
		Such a lovely place, such a lovely place,
		Such a Lovely Face!
	`

	var fileName string = "./Song.txt"
	writeTextFile(fileName, poem)
	fileContent := readTextFile(fileName)
	fmt.Println("\nThe File content of", fileName)
	fmt.Println(fileContent)
}

// This function writes the content into the file name and show the length of the file.
func writeTextFile(fileName string, content string) {
	// Create the file.
	file, err := os.Create(fileName)
	checkError(err)

	// Writes into the file.
	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("\n%s is written with %d characters. \n", fileName, length)

	defer file.Close()

}

// This function reads the content of the file name and return the content.
func readTextFile(fileName string) string {
	// Read the content from the file in bytes
	data, err := os.ReadFile(fileName)
	checkError(err)
	content := string(data)
	return content
}

// This function checks the error and raise exception.
func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
