
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

## Considerations

### Execution of the Solution

- **Setup**: Ensure Go is installed on your system. Clone the repository and navigate to the project directory.
- **Running the API**: Execute `go run main.go` to start the server. The API will be accessible on `localhost:8080`.
- **Dependencies**: This project requires the `github.com/dgrijalva/jwt-go` package for JWT handling. Use `go get` to install this package.

### Assumptions

- User authentication is simulated; in a production environment, integrate with a real user database.
- JWT tokens are used for simplicity in securing API endpoints.
- The payment processing logic is a placeholder and should be replaced with actual payment gateway integration.

### Areas for Improvement

- **Security**: Implement more robust user authentication and authorization mechanisms.
- **Database Integration**: Connect to a real database for user and payment data management.
- **Error Handling**: Enhance error handling and logging for better debugging and traceability.
- **Scalability**: Optimize the application for scalability and high availability.

### Cloud Technologies

- **Hosting**: The API can be hosted on cloud platforms like AWS, Azure, or GCP for scalability.
- **Database**: Consider using cloud-based databases like AWS RDS or Google Cloud SQL.
- **Reasoning**: Cloud technologies offer scalability, reliability, and distributed computing capabilities, making them suitable for online payment platforms.
