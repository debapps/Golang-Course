package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	EmployeeID, FirstName, LastName, AddressLine1, AddressLine2, City, State, ZipCode string
}

// This function shows employee details.
func (e *Employee) showEmp() {
	log.Printf("\nEmployee Details:\nEmployee ID - %s \nName - %s %s \nAddress Line 1 - %s \nAddress Line 2 - %s \nCity - %s \nState - %s \nZip Code - %s\n", e.EmployeeID, e.FirstName, e.LastName, e.AddressLine1, e.AddressLine2, e.City, e.State, e.ZipCode)
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

// This function validates all the fields.
// 1. Emp ID should be 10 characters.
// 2. Emp ID should be digit only.
// 3. Zip Code should be 6 characters.
// 4. Zip Code should be digit only.
// 5. Name, Address, City and State fields should not have any special character.
func (e *Employee) validateEmpFields() bool {
	// Checks the fields contains any special characters or not.
	isValidFName := !checkSpclChar(e.FirstName)
	isValidLName := !checkSpclChar(e.LastName)
	isValidAddr1 := !checkSpclChar(e.AddressLine1)
	isValidAddr2 := !checkSpclChar(e.AddressLine2)
	isValidCity := !checkSpclChar(e.City)
	isValidState := !checkSpclChar(e.State)

	isValidEmp := e.validateEmpId() && e.validateZip() && isValidFName && isValidLName && isValidAddr1 && isValidAddr2 && isValidCity && isValidState

	return isValidEmp
}

// This function returns the string slice of the employee details.
func (e *Employee) formatEmp() []string {
	return []string{e.EmployeeID, e.FirstName, e.LastName, e.AddressLine1, e.AddressLine2, e.City, e.State, e.ZipCode}
}

func main() {
	log.Print("\n********** CSV File Handling **********\n\n")
	var inputFilePath string = "./emp_data.csv"
	var outputFilePath string = "./valid_emp_data.csv"

	// Open the Output file.
	outFile, err := os.Create(outputFilePath)
	checkError(err)

	// Close the Output file at the end.
	defer outFile.Close()

	// Create CSV writer object.
	writer := csv.NewWriter(outFile)

	// Read input CSV File.
	empList := readEmpCSV(inputFilePath)
	if empList == nil {
		panic("Empty File.")
	}

	// Traverse Employee List. Checks for valid employee record. Writes the valid records into output file.
	for idx := range empList {
		var emp Employee = empList[idx]
		if emp.validateEmpFields() {
			log.Printf("\n\nWritting Valid Employee - %v into File - %s ", emp.EmployeeID, outputFilePath)
			writeEmpCSV(emp, writer)
		} else {
			log.Println("\n\nSkipping invalid Employee -  ")
			emp.showEmp()
		}
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
	log.Println("File Headers - ", header)

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

// This function writes the employee details into output CSV file.
func writeEmpCSV(emp Employee, writer *csv.Writer) {
	err := writer.Write(emp.formatEmp())
	checkError(err)
	writer.Flush()
}

// This function checks if there is a special character in the input string.
func checkSpclChar(inStr string) bool {
	// List of special characters.
	var specialChars = []rune{',', ';', ':', '"', '\'', '\\', '/', '{', '}', '|', '[', ']', '+',
		'=', '-', '_', '(', ')', '*', '&', 'ˆ', '%', '$', '#', '@', '!', '˜', '`'}

	var isContainSpecial bool
	for _, char := range specialChars {
		isContainSpecial = strings.ContainsRune(inStr, char)
		if isContainSpecial {
			break
		}
	}

	return isContainSpecial
}

// This function checks the error and raise if it exists.
func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}
