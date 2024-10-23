# Project Overview
This project demonstrates how to build RESTful APIs using the Go programming language's standard library. Go's standard library provides robust and efficient packages for HTTP handling, JSON encoding/decoding, and other common tasks required for API development without the need for external libraries.

## Key Features
HTTP Handling: The net/http package is used to define API routes, handle incoming HTTP requests, and send appropriate responses.
JSON Encoding/Decoding: The encoding/json package is utilized to marshal and unmarshal JSON data, ensuring seamless communication between the server and the client.
URL Parameters & Query Handling: URL path parameters and query parameters are managed through the standard library's functionalities, simplifying input parsing.
Error Handling: Proper error handling mechanisms are incorporated using Go’s built-in error types and status codes.
API Endpoints

The project provides the following API endpoints:

* GET /: Retrieves all tasks.
* POST /: Creates a new task.
* GET /id: Retrieves a specific task by its ID.
* PUT /id: Updates an existing task.

## Why Use Go Standard Library?
The decision to use Go's standard library for building APIs ensures:

* Simplicity: By avoiding third-party dependencies, the code remains simple and easy to maintain.
* Performance: The Go standard library is highly optimized for performance, which makes it suitable for high-performance applications.
* Portability: The code remains lightweight and portable, as it does not rely on external libraries or frameworks.
