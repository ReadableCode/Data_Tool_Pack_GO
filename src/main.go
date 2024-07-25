package main

import (
	"fmt"
    "context"
    // "encoding/json"
    "log"
    "os"

    "github.com/joho/godotenv"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/sheets/v4"
)

func main() {
	fmt.Println("Hello, World!")

	// Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	
	// print the key
	fmt.Println(os.Getenv("GOOGLE_SERVICE_ACCOUNT"))

    // Get the service account key from the environment variable
    key := os.Getenv("GOOGLE_SERVICE_ACCOUNT")
    if key == "" {
        log.Fatalf("GOOGLE_SERVICE_ACCOUNT not set in .env file")
    }

    // Decode the key into a google.Config object
    config, err := google.JWTConfigFromJSON([]byte(key), sheets.SpreadsheetsReadonlyScope)
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }

    // Create a new client
    ctx := context.Background()
    client := config.Client(ctx)

    // Create a new Sheets service
    srv, err := sheets.New(client)
    if err != nil {
        log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }

    // Specify the spreadsheet ID and range
    spreadsheetId := "1pvmIGeanVd0mjIO4-y53OY-z-ueLIY1AF7e-KZGAMzI"
    readRange := "rust_test!A1:D"

    // Call the Sheets API
    resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data from sheet: %v", err)
    }

    // Print the data
    if len(resp.Values) == 0 {
        fmt.Println("No data found.")
    } else {
        fmt.Println("Data:")
        for _, row := range resp.Values {
            fmt.Println(row)
        }
    }
}