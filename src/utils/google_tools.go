// src/utils/google_sheets.go
package utils

import (
    "context"
    "fmt"
    "os"
    "sync"

    "golang.org/x/oauth2/google"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)

var (
    srv   *sheets.Service
    once  sync.Once
    err   error
)

func initializeService() {
	fmt.Println("################\nInitializing Google Sheets service\n################")

    // Get the service account key from the environment variable
    key := os.Getenv("GOOGLE_SERVICE_ACCOUNT")
    if key == "" {
        err = fmt.Errorf("GOOGLE_SERVICE_ACCOUNT not set in .env file")
        return
    }

    // Decode the key into a google.Config object
    config, err := google.JWTConfigFromJSON([]byte(key), sheets.SpreadsheetsReadonlyScope)
    if err != nil {
        err = fmt.Errorf("unable to parse client secret file to config: %v", err)
        return
    }

    // Create a new client
    ctx := context.Background()
    client := config.Client(ctx)

    // Create a new Sheets service
    srv, err = sheets.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
        err = fmt.Errorf("unable to retrieve Sheets client: %v", err)
    }
}

// ReadGoogleSheet fetches data from the specified range
func ReadGoogleSheet(spreadsheetId, sheetName, readRange string) ([][]interface{}, error) {
    once.Do(initializeService)
    if err != nil {
        return nil, err
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
