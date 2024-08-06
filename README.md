# Data_Tool_Pack_GO

## Setup on Linux

- Using apt

  - Open terminal and run the following commands:

  ```bash
  sudo apt update
  sudo apt install golang
  ```

## Setup on Windows

- Using WinGet

  - Open powershell as administrator and run:
  
  ```bash
  winget install -e --id GoLang.Go
  ```

## Testing and Finishing Installation

- If using VSCode, install the Go extension by searching for `@id:golang.go` in the extensions tab.

- Close and reopen the terminal to make sure installation is successful and then run the folling commands to verify the version of Go:

  ```bash
  go version
  ```

## Run Directly from Source

- Run the program by executing the following command:

  ```bash
  go run main.go -action=read1
  ```

## Compiling from source

- To build it and run the executable:

  - cd to directory where the main.go file is located

  ```bash
  go build main.go
  ```

  - Running on Linux
  
    ```bash
    chmod +x main
    ./main -action=read1
    ```
  
  - Running on Windows
  
    ```bash
    ./main.exe -action=read1
    ```
