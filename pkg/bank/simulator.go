package bank

import (
	"payment-platform-solid/internal/model"
)

// BankSimulator simulates bank operations
type BankSimulator struct{}

func (b *BankSimulator) ProcessPayment(payment model.Payment) (model.PaymentResponse, error) {
	// Simulate bank processing
	// This is a placeholder for actual bank integration
	return model.PaymentResponse{
		ID:     payment.ID,
		Status: "Success",
	}, nil
}

func (b *BankSimulator) ProcessRefund(refund model.Refund) (model.RefundResponse, error) {
	// Simulate refund processing
	// This is a placeholder for actual bank integration
	return model.RefundResponse{
		Status: "Success",
	}, nil
}
