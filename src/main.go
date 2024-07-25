package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// readGoogleSheet initializes the Sheets service and fetches data from the specified range
func readGoogleSheet(spreadsheetId, sheetName, readRange string) ([][]interface{}, error) {
	// Get the service account key from the environment variable
	key := os.Getenv("GOOGLE_SERVICE_ACCOUNT")
	if key == "" {
		return nil, fmt.Errorf("GOOGLE_SERVICE_ACCOUNT not set in .env file")
	}

	// Decode the key into a google.Config object
	config, err := google.JWTConfigFromJSON([]byte(key), sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	// Create a new client
	ctx := context.Background()
	client := config.Client(ctx)

	// Create a new Sheets service
	srv, err := sheets.New(client)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	// Specify the full range including the sheet name
	fullRange := fmt.Sprintf("%s!%s", sheetName, readRange)

	// Call the Sheets API
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, fullRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	return resp.Values, nil
}


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
	data, err := readGoogleSheet(spreadsheetId, sheetName, readRange)
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
