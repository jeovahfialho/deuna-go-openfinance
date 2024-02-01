
# Payment Platform API

This API provides a platform for processing payments, querying payment details, and handling refunds. It includes JWT authentication for secure access to endpoints.

## Getting Started

To start using this API, run the server and use an API client like Postman to make requests.

## Authentication

Before accessing the protected routes, obtain a JWT token by authenticating through the login endpoint.

### Login Endpoint

- **URL**: `/login`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "username": "user",
    "password": "password"
  }
  ```
- **Response**: JWT Token

## Protected Endpoints

Use the JWT token obtained from the login endpoint as a Bearer token in the Authorization header for the following requests.

### Process Payment

- **URL**: `/process-payment`
- **Method**: `POST`
- **Headers**:
  - `Authorization: Bearer <JWT_TOKEN>`
- **Body**: 
  ```json
  {
    "amount": 100.0,
    "description": "Payment description"
  }
  ```

### Query Payment Details

- **URL**: `/query-payment`
- **Method**: `GET`
- **Headers**:
  - `Authorization: Bearer <JWT_TOKEN>`
- **Query Parameters**:
  - `id`: Payment ID

### Process Refund

- **URL**: `/process-refund`
- **Method**: `POST`
- **Headers**:
  - `Authorization: Bearer <JWT_TOKEN>`
- **Body**: 
  ```json
  {
    "paymentId": "payment123",
    "amount": 50.0
  }
  ```

### Protected Route Example

- **URL**: `/protected-route`
- **Method**: `GET`
- **Headers**:
  - `Authorization: Bearer <JWT_TOKEN>`

## Running the Server

To run the server, use the command `go run main.go` in the root directory of the project. Ensure all dependencies are installed.
