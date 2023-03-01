package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// The Tour.
type Tour struct {
	PackageTitle, Name, Blurb, Price string
}

func (t Tour) ShowTour() {
	fmt.Printf("\n\n********************* %v *********************\n\n", t.PackageTitle)
	fmt.Println(" Name:", t.Name)
	fmt.Println(" Synopsis:", t.Blurb)
	fmt.Println("\n Price: $", t.Price)
}

func main() {
	var url string = "http://services.explorecalifornia.org/json/tours.php"
	bodyContent := getWebContent(url)
	// fmt.Println(bodyContent)
	tours := getToursFromJSON(bodyContent)

	for _, tour := range tours {
		tour.ShowTour()
	}

}

// This function gets the web content of the HTTP GET and returns the string content.
func getWebContent(url string) string {
	// HTTP GET.
	res, err := http.Get(url)
	checkError(err)

	// Close the entire body.
	defer res.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(res.Body)
	checkError(err)

	// Convert the body into string and return.
	content := string(body)
	return content

}

// This function parse the web body content into the Tour slice.
func getToursFromJSON(content string) []Tour {
	// Declare the Tour Slice.
	tourList := make([]Tour, 0, 30)

	// Initialize the decoder.
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	// Convert the content into Tour object.
	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		// Append Tour into the tour list.
		tourList = append(tourList, tour)
	}

	return tourList
}

// This function checks the error and raise exception.
func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
