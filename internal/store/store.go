package store

import (
	"errors"
	"payment-platform-solid/internal/model"
	"sync"
)

// Store simula um armazenamento de dados simples em memória.
type Store struct {
	payments map[string]model.Payment
	refunds  map[string]model.Refund
	mu       sync.RWMutex
}

// NewStore cria e retorna uma nova instância de Store.
func NewStore() *Store {
	return &Store{
		payments: make(map[string]model.Payment),
		refunds:  make(map[string]model.Refund),
	}
}

// SavePayment armazena um pagamento.
func (s *Store) SavePayment(payment model.Payment) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.payments[payment.ID]; exists {
		return errors.New("payment already exists")
	}

	s.payments[payment.ID] = payment
	return nil
}

// GetPayment retorna um pagamento pelo ID.
func (s *Store) GetPayment(id string) (model.Payment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	payment, exists := s.payments[id]
	if !exists {
		return model.Payment{}, errors.New("payment not found")
	}

	return payment, nil
}

// SaveRefund armazena um reembolso.
func (s *Store) SaveRefund(refund model.Refund) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.refunds[refund.PaymentID]; exists {
		return errors.New("refund already exists for this payment")
	}

	s.refunds[refund.PaymentID] = refund
	return nil
}

// GetRefund retorna um reembolso pelo ID do pagamento.
func (s *Store) GetRefund(paymentID string) (model.Refund, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	refund, exists := s.refunds[paymentID]
	if !exists {
		return model.Refund{}, errors.New("refund not found for this payment")
	}

	return refund, nil
}

// UpdatePaymentStatus atualiza o status de um pagamento existente.
func (s *Store) UpdatePaymentStatus(paymentID, newStatus string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Verifica se o pagamento existe
	payment, exists := s.payments[paymentID]
	if !exists {
		return errors.New("payment not found")
	}

	// Atualiza o status do pagamento
	payment.Status = newStatus

	// Salva o pagamento atualizado de volta no mapa
	s.payments[paymentID] = payment

	return nil
}
