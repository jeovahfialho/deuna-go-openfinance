package model

import "time"

type Payment struct {
	ID         string
	MerchantID string
	CustomerID string
	Amount     float64
	Status     string
	CreatedAt  time.Time
}

type PaymentResponse struct {
	ID     string
	Status string
}

type Refund struct {
	PaymentID string
	Amount    float64
	Status    string // Status of the refund like 'Success', 'Pending', 'Failed'
}

type RefundResponse struct {
	Status string
}
