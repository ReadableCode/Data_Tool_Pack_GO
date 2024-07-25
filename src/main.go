package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	// import from utils/google_toools.go
	"data_tool_pack_go/src/utils"
)


func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Define the spreadsheet ID, sheet name, and range
	spreadsheetId := "1pvmIGeanVd0mjIO4-y53OY-z-ueLIY1AF7e-KZGAMzI"
	sheetName := "rust_test"
	readRange := "A1:D"

	// Call the function to read data from the Google Sheet
	data, err := utils.ReadGoogleSheet(spreadsheetId, sheetName, readRange)
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
