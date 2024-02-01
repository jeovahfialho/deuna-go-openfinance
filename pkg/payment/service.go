package payment

import (
	"errors"
	"fmt"
	"payment-platform-solid/internal/model"
	"payment-platform-solid/internal/store"
	"payment-platform-solid/pkg/bank"
	"time"
)

// Service struct that implements the PaymentProcessor interface.
type Service struct {
	Store         *store.Store
	BankSimulator bank.BankSimulator
	AuditLogFunc  func(entry model.AuditLogEntry) // Function for audit trail logging
}

// NewService creates a new instance of the payment service.
func NewService(store *store.Store, bankSimulator bank.BankSimulator, auditLogFunc func(model.AuditLogEntry)) *Service {
	return &Service{
		Store:         store,
		BankSimulator: bankSimulator,
		AuditLogFunc:  auditLogFunc, // Assign the audit log function
	}
}

// ProcessPayment handles the logic to process a payment.
func (s *Service) ProcessPayment(payment model.Payment) (model.PaymentResponse, error) {

	// Inside ProcessPayment before processing the payment:
	err := ValidatePaymentDetails(payment)
	if err != nil {
		return model.PaymentResponse{}, err
	}

	// Simulate processing the payment with the bank.
	response, err := s.BankSimulator.ProcessPayment(payment)
	if err != nil {
		return model.PaymentResponse{}, err
	}

	// Assuming success, save the payment details.
	payment.Status = "Success"
	err = s.Store.SavePayment(payment)
	if err != nil {
		return model.PaymentResponse{}, err
	}

	// Record the payment processing in the audit trail
	s.AuditLogFunc(model.AuditLogEntry{
		Timestamp:   time.Now(),
		UserID:      payment.CustomerID,
		Action:      "ProcessPayment",
		Description: "Processed payment of amount " + fmt.Sprintf("%f", payment.Amount),
	})

	// Return a successful response.
	return response, nil
}

// QueryPaymentDetails retrieves details of a previously made payment.
func (s *Service) QueryPaymentDetails(id string) (model.Payment, error) {
	payment, err := s.Store.GetPayment(id)
	if err != nil {
		// Payment not found or other error.
		return model.Payment{}, err
	}

	return payment, nil
}

// ProcessRefund handles the logic to process a refund.
func (s *Service) ProcessRefund(refund model.Refund) (model.RefundResponse, error) {
	// First, verify the original payment exists and is valid for a refund.
	originalPayment, err := s.Store.GetPayment(refund.PaymentID)
	if err != nil {
		return model.RefundResponse{}, err
	}

	err = ValidateRefundRequest(refund, originalPayment)
	if err != nil {
		return model.RefundResponse{}, err
	}

	// Simulate processing the refund with the bank.
	refundResponse, err := s.BankSimulator.ProcessRefund(refund)
	if err != nil {
		return model.RefundResponse{}, err
	}

	// Assuming success, save the refund details.
	refund.Status = "Success"
	err = s.Store.SaveRefund(refund)
	if err != nil {
		return model.RefundResponse{}, err
	}

	// Update the original payment status if necessary.
	// Atualiza o status do pagamento original, se necessário.
	err = s.Store.UpdatePaymentStatus(refund.PaymentID, "Refunded")
	if err != nil {
		return model.RefundResponse{}, err
	}

	// Return a successful refund response.
	return refundResponse, nil
}

// ValidateRefundRequest checks if the refund request is valid based on the original payment.
func ValidateRefundRequest(refund model.Refund, originalPayment model.Payment) error {
	if originalPayment.Status != "Success" {
		return errors.New("the original payment was not successful and cannot be refunded")
	}
	if refund.Amount <= 0 {
		return errors.New("refund amount must be greater than zero")
	}
	if refund.Amount > originalPayment.Amount {
		return errors.New("refund amount cannot be greater than the original payment amount")
	}
	// Check if the payment has already been refunded, if necessary.
	// This requires keeping track of refunded amounts or refund statuses.
	// Add any other necessary validations here.
	return nil
}

// ValidatePaymentDetails checks if the payment details are valid.
func ValidatePaymentDetails(payment model.Payment) error {
	if payment.Amount <= 0 {
		return errors.New("payment amount must be greater than zero")
	}
	if payment.MerchantID == "" {
		return errors.New("merchant ID is required")
	}
	if payment.CustomerID == "" {
		return errors.New("customer ID is required")
	}
	// Add any other necessary validations here.
	return nil
}
