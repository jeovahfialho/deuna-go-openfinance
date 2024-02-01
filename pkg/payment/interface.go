package payment

import "payment-platform-solid/internal/model"

type PaymentProcessor interface {
	ProcessPayment(payment model.Payment) (model.PaymentResponse, error)
	QueryPaymentDetails(id string) (model.Payment, error)
	ProcessRefund(payment model.Refund) (model.RefundResponse, error)
}
