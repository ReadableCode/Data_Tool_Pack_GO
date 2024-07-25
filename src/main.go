package main

import (
	"fmt"
	"log"

	"data_tool_pack_go/src/utils"
)

func main() {
	// Declare variables
	var (
		data          [][]interface{}
		spreadsheetId string
		sheetName     string
		readRange     string
		writeRange    string
        err           error
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
	
    // Example data to write
    values := [][]interface{}{
        {"A", "B", "C", "D"},
        {1, 2, 3, 4},
        {5, 6, 7, 8},
    }

    // Call to write data
    writeRange = "A10:D12"
    err = utils.WriteGoogleSheet(spreadsheetId, sheetName, writeRange, values)
    if err != nil {
        log.Fatalf("Unable to write data to sheet: %v", err)
    } else {
        fmt.Println("Data successfully written to the sheet.")
    }
}
