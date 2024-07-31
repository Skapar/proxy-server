# Proxy Server

## Introduction

This project is an HTTP proxy server written in Go. It accepts JSON requests to proxy HTTP requests to external services, returns the responses back to the client, and logs the requests and responses locally.

## Features

- Accepts HTTP requests with method, URL, and headers specified in the JSON body.
- Proxies the request to the specified URL and returns the response.
- Logs each request and response with a unique ID.
- Supports validation of incoming request data.
- Provides a health check endpoint.
- Includes Swagger documentation.

## Prerequisites

- Docker
- Docker Compose
- Make

## Installation

1. **Clone the repository**

   git clone <https://github.com/Skapar/proxy-server.git>
   cd yourproject

## Usage

### Building and Running

To build the Docker image:

make build

To run the Docker container:

make run
Usage
Start the server

Ensure the server is running by executing the above commands. The server will be listening on port 8080.

## Sending a POST Request

You can send a POST request to the proxy server using tools like Postman or curl. Example curl command:

```bash
curl --location --request POST 'http://localhost:8080/proxy' \
--header 'Content-Type: application/json' \
--data '{
  "method": "GET",
  "url": "http://example.com",
  "headers": {
    "Authorization": "Bearer token"
  }
}'
```

## The server will return a JSON response containing

```bash
id: Unique identifier for the request.
status: HTTP status code from the external service.
headers: Headers from the response of the external service.
length: Length of the response body.
Example response:
```

```bash
{
  "id": "unique-request-id",
  "status": 200,
  "headers": { 
    "Content-Type": "text/html; charset=UTF-8" 
    // other headers
    
  "length": 1256
}
```

## Health Check

Check the health of the server by sending a GET request to /health:

```bash
curl --location --request GET 'http://localhost:8080/health'
This should return OK with a 200 status code if the server is healthy.
```

## Swagger Documentation

Access the Swagger documentation at <http://localhost:8080/swagger/> to view and interact with the API endpoints.

## Project Structure

myproxy/
  ├── cmd/
  │   └── main.go
  ├── docs/
  ├── internal/
  │   └── proxy/
  │       ├── handler.go
  │       ├── health.go
  │       └── storage.go
  ├── swagger.yaml
  ├── docker-compose.yml
  ├── Dockerfile
  ├── go.mod
  ├── go.sum
  ├── Makefile
  └── README.md

## Files

cmd/main.go: Entry point for the application.
internal/proxy/handler.go: Contains the logic for handling and validating incoming requests.
internal/proxy/health.go: Contains the health check endpoint.
internal/proxy/storage.go: Manages the storage of requests and responses.
swagger.yaml: Defines the Swagger documentation for the API.
Dockerfile: Defines the Docker image.
docker-compose.yml: Manages Docker services.
Makefile: Automates the build and run processes.

## Validation

The server performs validation on the incoming request data to ensure that the required fields are present. If any fields are missing or invalid, the server returns a 400 Bad Request response with a detailed error message.
