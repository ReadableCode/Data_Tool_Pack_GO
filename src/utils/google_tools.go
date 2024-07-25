// src/utils/google_sheets.go
package utils

import (
    "context"
    "fmt"
    "os"

    "golang.org/x/oauth2/google"
    "google.golang.org/api/sheets/v4"
)

// ReadGoogleSheet initializes the Sheets service and fetches data from the specified range
func ReadGoogleSheet(spreadsheetId, sheetName, readRange string) ([][]interface{}, error) {
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
