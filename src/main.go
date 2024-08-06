package main

import (
	"flag"
	"fmt"
	"log"

	"data_tool_pack_go/src/utils"
)

func main() {
	// Define CLI flags
	action := flag.String("action", "", "Action to perform: read1, read2, write")
	flag.Parse()

	// Declare variables
	var (
		data          [][]interface{}
		spreadsheetId string
		sheetName     string
		readRange     string
		writeRange    string
		err           error
	)

	// Google Sheet details
	spreadsheetId = "1pvmIGeanVd0mjIO4-y53OY-z-ueLIY1AF7e-KZGAMzI"
	sheetName = "rust_test"

	switch *action {
	case "read1":
		readRange = "A1:D"
		data, err = utils.ReadGoogleSheet(spreadsheetId, sheetName, readRange)
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}
		printData(data)

	case "read2":
		readRange = "A1:C"
		data, err = utils.ReadGoogleSheet(spreadsheetId, sheetName, readRange)
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}
		printData(data)

	case "write":
		// Example data to write
		values := [][]interface{}{
			{"A", "B", "C", "D"},
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		}
		writeRange = "A10:D12"
		err = utils.WriteGoogleSheet(spreadsheetId, sheetName, writeRange, values)
		if err != nil {
			log.Fatalf("Unable to write data to sheet: %v", err)
		} else {
			fmt.Println("Data successfully written to the sheet.")
		}

	default:
		fmt.Println("Invalid action. Please use 'read1', 'read2', or 'write'.")
	}
}

func printData(data [][]interface{}) {
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
