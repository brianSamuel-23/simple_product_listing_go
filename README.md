# Simple Product Listing Application

This guide will help you start the Simple Product Listing application using Go and Fiber as the HTTP handler.

## Prerequisites

- Go installed on your machine (version 1.17+)
- Git installed on your machine
- Set JWT secret on your machine as an environment variable:

    ```sh
    JWT_SECRET=your_secret_here
    ```

## Installation

1. Install the dependencies:

    ```sh
    go mod tidy
    ```

## Running the Application

1. Navigate to the `cmd` directory:

    ```sh
    cd cmd
    ```

2. Build the application:

    ```sh
    go build main.go
    ```
   
3. Run the application:

    ```sh
    ./main
    ```
   
4. The application should now be running on `http://localhost:3000`.

