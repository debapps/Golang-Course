package main

import (
	"bufio"
	"fmt"
	"os"
)

const aConst float64 = 3.14

func main() {
	// Variable declarations.
	var myText string = "Welcome to Go Programming!"
	showString(myText)

	// Initializing Variable. Type of the variable defined implecitly.
	newText := "Hello World!"
	showString(newText)

	// Numeric Variable declarations.
	var num1 int = 10
	// type casting
	showInt(int64(num1))

	var num2 = 99.93456
	showFloat(num2)

	// Showing global constant.
	showFloat(aConst)

	// Getting input from the console.
	name := getConsoleInput("Enter Your Name: ")
	fmt.Println("Hello,", name)

}

func showString(item string) {
	fmt.Println(item)
	fmt.Printf("Data Type - %T\n\n", item)
}

func showInt(item int64) {
	fmt.Println(item)
	fmt.Printf("Data Type - %T\n\n", item)
}

func showFloat(item float64) {
	fmt.Printf("%.3f", item)
	fmt.Printf("\nData Type - %T\n\n", item)
}

func getConsoleInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return input
}
