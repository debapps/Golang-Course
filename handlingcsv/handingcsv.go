package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	EmployeeID, FirstName, LastName, AddressLine1, AddressLine2, City, State, ZipCode string
}

// This function shows employee detailes.
func (e *Employee) showEmp() {
	fmt.Printf("\nEmployee Details:\nEmployee ID - %s \nName - %s %s \nAddress Line 1 - %s \nAddress Line 2 - %s \nCity - %s \nState - %s \nZip Code - %s\n", e.EmployeeID, e.FirstName, e.LastName, e.AddressLine1, e.AddressLine2, e.City, e.State, e.ZipCode)
}

// This function validates Employee ID based on following criteria:
// 1. Should be 10 characters.
// 2. Should be digit only.
func (e *Employee) validateEmpId() bool {
	var validEmpID, isNumeric, is10Char bool

	if a, _ := strconv.Atoi(e.EmployeeID); a > 0 {
		isNumeric = true
	}

	if len(e.EmployeeID) == 10 {
		is10Char = true
	}

	if isNumeric && is10Char {
		validEmpID = true
	}

	return validEmpID
}

// This function validates Zip code based on following criteria:
// 1. Should be 6 characters.
// 2. Should be digit only.
func (e *Employee) validateZip() bool {
	var validEmpID, isNumeric, is10Char bool

	if a, _ := strconv.Atoi(e.ZipCode); a > 0 {
		isNumeric = true
	}

	if len(e.ZipCode) == 6 {
		is10Char = true
	}

	if isNumeric && is10Char {
		validEmpID = true
	}

	return validEmpID
}

func main() {
	fmt.Print("\n********** CSV File Handling **********\n\n")
	var inputFilePath string = "./emp_data.csv"

	// Read input CSV File.
	empList := readEmpCSV(inputFilePath)
	if empList == nil {
		panic("Empty File.")
	}

	// Traverse Employee List.
	for idx := range empList {
		var emp Employee = empList[idx]
		emp.showEmp()
		fmt.Printf("Valid EmpID? %v \n", emp.validateEmpId())
		fmt.Printf("Valid Zip? %v \n", emp.validateZip())
	}
}

// This function reads the employee CSV file and returns the Employee object list.
func readEmpCSV(filePath string) []Employee {
	// Read the file from the path and gets the bytes of data
	data, err := os.ReadFile(filePath)
	checkError(err)

	// Creates the CSV reader object from the dara bytes.
	reader := csv.NewReader(strings.NewReader(string(data)))

	// Get the all the records into slice.
	header, err := reader.Read()
	if err == io.EOF {
		return nil
	}
	checkError(err)

	// Prints the Header.
	fmt.Println("File Headers - ", header)

	// Create an empty employee list.
	empList := []Employee{}

	// Traverse all the records.
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		checkError(err)

		// Create the Employee object.
		emp := Employee{record[0], record[1], record[2], record[3],
			record[4], record[5], record[6], record[7]}

		// Append the employee object to the list.
		empList = append(empList, emp)

	}

	return empList
}

// This function checks the error and raise if it exists.
func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}
