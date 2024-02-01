# Online Payment Platform

This project implements a basic online payment platform using Go, adhering to SOLID principles. It provides APIs for processing payments, querying payment details, and processing refunds, along with a bank simulator for the transaction process.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.17 or later

### Installing

Clone the repository to your local machine:

```bash
git clone https://yourrepositoryurl/payment-platform-solid.git
cd payment-platform-solid
```

Initialize the module (if not already done):

```bash
go mod init payment-platform-solid
```

Run the application:

```bash
go run cmd/main.go
```

The server will start on port 8080.

## API Endpoints

The following API endpoints are available:

### Process Payment

- **URL**: `/process-payment`
- **Method**: `POST`
- **Data Params**

```json
{
  "ID": "payment123",
  "MerchantID": "merchant123",
  "CustomerID": "customer123",
  "Amount": 100.50
}
```

- **Success Response**:

```json
{
  "ID": "payment123",
  "Status": "Success"
}
```

### Query Payment Details

- **URL**: `/query-payment?id=payment123`
- **Method**: `GET`
- **URL Params**: `id=[string]`

- **Success Response**:

```json
{
  "ID": "payment123",
  "MerchantID": "merchant123",
  "CustomerID": "customer123",
  "Amount": 100.50,
  "Status": "Success",
  "CreatedAt": "2024-02-01T15:04:05Z"
}
```

### Process Refund

- **URL**: `/process-refund`
- **Method**: `POST`
- **Data Params**

```json
{
  "PaymentID": "payment123",
  "Amount": 100.50
}
```

- **Success Response**:

```json
{
  "Status": "Success"
}
```
