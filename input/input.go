package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This function gets console input as string.
func GetStringInput(prompt string) string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	strValue, _ := reader.ReadString('\n')
	return strings.TrimSpace(strValue)

}

// This function gets console input as int64.
func GetInt64Input(prompt string) int64 {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	strValue, _ := reader.ReadString('\n')
	intValue, err := strconv.ParseInt(strings.TrimSpace(strValue), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return intValue

}

// This function gets console input as float64.
func GetFloat64Input(prompt string) float64 {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	strValue, _ := reader.ReadString('\n')
	floatValue, err := strconv.ParseFloat(strings.TrimSpace(strValue), 64)
	if err != nil {
		panic(err.Error())
	}
	return floatValue

}
