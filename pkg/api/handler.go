package api

import (
	"encoding/json"
	"net/http"
	"payment-platform-solid/internal/model"
	"payment-platform-solid/pkg/payment"
)

type PaymentHandler struct {
	Service payment.PaymentProcessor
}

func (h *PaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var payment model.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.Service.ProcessPayment(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func (h *PaymentHandler) QueryPaymentDetails(w http.ResponseWriter, r *http.Request) {
	// Assume que o ID do pagamento é passado como parâmetro na URL
	paymentID := r.URL.Query().Get("id")
	if paymentID == "" {
		http.Error(w, "Payment ID is required", http.StatusBadRequest)
		return
	}

	payment, err := h.Service.QueryPaymentDetails(paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(payment)
}

func (h *PaymentHandler) ProcessRefund(w http.ResponseWriter, r *http.Request) {
	var refund model.Refund
	if err := json.NewDecoder(r.Body).Decode(&refund); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.Service.ProcessRefund(refund)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

// Implement other handlers (QueryPaymentDetails, ProcessRefund) in a similar fashion
