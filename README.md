# Data_Tool_Pack_GO

## Setup Go

* Follow instructions in [setup_go.md](./docs/setup_go.md)

## Run Directly from Source

* cd into the src diretory where the main.go file is located

* Run the program by executing the following command:

  ```bash
  go run main.go -action=read1
  ```

## Compiling from source

* To build it and run the executable:

  * cd into the src directory where the main.go file is located

  ```bash
  go build main.go
  ```

  * Running on Linux
  
    ```bash
    chmod +x main
    ./main -action=read1
    ```
  
  * Running on Windows
  
    ```bash
    .\main.exe -action=read1
    ```
