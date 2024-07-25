package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"data_tool_pack_go/src/utils"
)


func main() {
    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	// Declare variables
    var (
        data          [][]interface{}
        spreadsheetId string
        sheetName     string
        readRange     string
    )

	// Call the function to read data from the Google Sheet
    spreadsheetId = "1pvmIGeanVd0mjIO4-y53OY-z-ueLIY1AF7e-KZGAMzI"
    sheetName = "rust_test"
    readRange = "A1:D"
    data, err = utils.ReadGoogleSheet(spreadsheetId, sheetName, readRange)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Print the data
	if len(data) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data:")
		for _, row := range data {
			fmt.Println(row)
		}
	}

	// Call the function to read data from the Google Sheet
    spreadsheetId = "1pvmIGeanVd0mjIO4-y53OY-z-ueLIY1AF7e-KZGAMzI"
    sheetName = "rust_test"
    readRange = "A1:C"
    data, err = utils.ReadGoogleSheet(spreadsheetId, sheetName, readRange)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Print the data
	if len(data) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data:")
		for _, row := range data {
			fmt.Println(row)
		}
	}
}
