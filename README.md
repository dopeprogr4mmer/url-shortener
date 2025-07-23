# Go URL Shortener

A simple, production-ready URL shortener built using Go and SQLite. This project allows users to shorten URLs and redirect to the original URLs using a short code.

## Features

- **Shorten URLs**: Submit a URL and get a short code.
- **Redirect**: Visit the short code URL to get redirected to the original URL.
- **SQLite Database**: Stores URL mappings locally.
- **Logging**: Provides basic logging for server operations and requests.
- **Go Modules**: Uses Go modules for dependency management.


## Requirements

- Go 1.11 or higher
- SQLite database
- `.env` file for configuration (e.g., DB path, port)

## Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```

2. **Install Go dependencies**:
Run the following command to install required dependencies:
    ```bash
    go mod tidy
    ```

3. **Configure your environment**:
Create a .env file in the root of the project with the following contents:
PORT=8080
DB_PATH=shortener.db

4. **Run the application**:
To start the server, run:
    ```bash
    go run main.go
    ```
The server will be available at http://localhost:8080.


## API Endpoints

### Shorten URL
- **POST** `/shorten`

  **Request Body**:
  ```json
  {
    "url": "https://example.com"
  }

Response:

{
  "short_url": "http://localhost:8080/<short-code>"
}

### Redirect to Original URL
- **GET** /<short-code>
This endpoint will redirect to the original URL for the provided short code.

## Example Usage

Shorten a URL:
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'
Redirect using the short URL:
Navigate to http://localhost:8080/<short-code> to be redirected to the original URL.

## Example Usage

Shorten a URL:
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'
Redirect using the short URL:
Navigate to http://localhost:8080/<short-code> to be redirected to the original URL.

## Logging

The server logs all incoming requests and database operations.
Logs are printed in the following format:
[INFO] Shortening URL: https://example.com
[INFO] Generated short code: <short-code>
[INFO] Redirecting <short-code> => https://example.com


## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.
