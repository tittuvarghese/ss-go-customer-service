# Customer Service for E-Commerce Application

[![GoDoc](https://pkg.go.dev/badge/github.com/tittuvarghese/ss-go-customer-service)](https://pkg.go.dev/github.com/tittuvarghese/ss-go-customer-service)
[![Build Status](https://travis-ci.org/tittuvarghese/ss-go-customer-service.svg?branch=main)](https://travis-ci.org/tittuvarghese/ss-go-customer-service)

The **Customer Service** is a microservice responsible for managing user authentication and profile details in the e-commerce platform. It provides endpoints for user registration, login, and profile retrieval. This service communicates via gRPC, and it is designed to be accessed by the **Gateway Service** and other microservices.

## API Overview

The **Customer Service** exposes the following gRPC methods:

### 1. **Register**  
- **RPC Method**: `Register`
- **Request Type**: `RegisterRequest`
- **Response Type**: `RegisterResponse`
- **Description**: Registers a new user with the provided details.

#### Request (RegisterRequest)
```proto
message RegisterRequest {
  string firstname = 1;
  string lastname = 2;
  string username = 3;
  string password = 4;
  string type = 5; // e.g., "customer", "admin"
}
```

#### Response (RegisterResponse)
```proto
message RegisterResponse {
  string message = 1; // Success or failure message
}
```

### 2. **Login**
- **RPC Method**: `Login`
- **Request Type**: `LoginRequest`
- **Response Type**: `LoginResponse`
- **Description**: Authenticates a user and returns a JWT token upon successful login.

#### Request (LoginRequest)
```proto
message LoginRequest {
  string username = 1;
  string password = 2;
}
```

#### Response (LoginResponse)
```proto
message LoginResponse {
  bool status = 1;   // Indicates if the login is successful
  string token = 2;  // JWT token for further authenticated requests
}
```

### 3. **GetProfile**
- **RPC Method**: `GetProfile`
- **Request Type**: `GetProfileRequest`
- **Response Type**: `GetProfileResponse`
- **Description**: Fetches the profile information of a logged-in user using the user ID.

#### Request (GetProfileRequest)
```proto
message GetProfileRequest {
  string userid = 1;  // The ID of the user whose profile is being requested
}
```

#### Response (GetProfileResponse)
```proto
message GetProfileResponse {
  string userid = 1;
  string username = 2;
  string firstname = 3;
  string lastname = 4;
  string type = 5; // User type (e.g., "customer", "admin")
}
```

## How It Works

- The **Gateway Service** sends HTTP requests, which are forwarded to this Customer Service via gRPC.
- For user authentication and registration, the service uses the **Register**, **Login**, and **GetProfile** methods to interact with the client-side and return relevant data.
- The `Login` method returns a JWT token, which is used by clients for authenticating subsequent requests to other services in the system.
- The service securely stores user credentials and other profile data, such as first name, last name, username, and user type.

## Running the Service Locally

### Prerequisites

Before running the Customer Service locally, ensure the following:

- Go 1.18 or higher
- Protocol Buffers (Protobuf) Compiler (`protoc`)
- gRPC Go Plugin for Protobuf (`protoc-gen-go` and `protoc-gen-go-grpc`)

### Steps to Run Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/tittuvarghese/ss-go-customer-service.git
   cd customer-service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Generate gRPC code from the `proto` file:
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/customer.proto
   ```

4. Start the Customer Service:
   ```bash
   go run cmd/main.go
   ```

The service will be up and running and listening for gRPC requests.

## Architecture

The **Customer Service** is part of a microservice-based architecture. It handles the user-related logic and communicates with other services using gRPC. Below is a basic overview of the architecture:

- **gRPC Protocol**: Handles communication between the **Gateway Service** and **Customer Service**.
- **JWT Authentication**: Secure authentication mechanism used for protecting endpoints and enabling secure communication.
- **Database**: Stores user information (e.g., usernames, passwords, profile details).
