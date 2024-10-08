// src/utils/google_sheets.go
package utils

import (
    "context"
    "fmt"
    "os"
    "sync"
	"path/filepath"
	"log"
    "runtime"

	"github.com/joho/godotenv"

    "golang.org/x/oauth2/google"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)

var (
    srv   *sheets.Service
    once  sync.Once
    initErr   error
)

func initializeService() {
    fmt.Println("#####################\nInitializing Google Sheets service\n#####################")

    // Get the directory of the source file
    _, b, _, _ := runtime.Caller(0)
    basepath := filepath.Dir(b)

    // Construct the path to the .env file relative to the source file
    envPath := filepath.Join(basepath, "../..", ".env")
    err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

    // Get the service account key from the environment variable
    key := os.Getenv("GOOGLE_SERVICE_ACCOUNT")
    if key == "" {
        initErr = fmt.Errorf("GOOGLE_SERVICE_ACCOUNT not set in .env file")
        return
    }

    // Decode the key into a google.Config object
    config, err := google.JWTConfigFromJSON([]byte(key), sheets.SpreadsheetsScope)
    if err != nil {
        initErr = fmt.Errorf("unable to parse client secret file to config: %v", err)
        return
    }

    // Create a new client
    ctx := context.Background()
    client := config.Client(ctx)

    // Create a new Sheets service
    srv, err = sheets.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
        initErr = fmt.Errorf("unable to retrieve Sheets client: %v", err)
    }
}

// ReadGoogleSheet fetches data from the specified range
func ReadGoogleSheet(spreadsheetId, sheetName, readRange string) ([][]interface{}, error) {
    once.Do(initializeService)
    if initErr != nil {
        return nil, initErr
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


// WriteGoogleSheet writes data to the specified range
func WriteGoogleSheet(spreadsheetId, sheetName, writeRange string, values [][]interface{}) error {
    once.Do(initializeService)
    if initErr != nil {
        return initErr
    }

    // Specify the full range including the sheet name
    fullRange := fmt.Sprintf("%s!%s", sheetName, writeRange)

    // Prepare the data to be written
    valueRange := &sheets.ValueRange{
        Values: values,
    }

    // Call the Sheets API
    _, err := srv.Spreadsheets.Values.Update(spreadsheetId, fullRange, valueRange).ValueInputOption("RAW").Do()
    if err != nil {
        return fmt.Errorf("unable to write data to sheet: %v", err)
    }

    return nil
}