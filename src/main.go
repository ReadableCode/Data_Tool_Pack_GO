package main

import (
	"fmt"
    // "context"
    // "encoding/json"
    "log"
    "os"

    "github.com/joho/godotenv"
    // "golang.org/x/oauth2/google"
    // "google.golang.org/api/sheets/v4"
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
}