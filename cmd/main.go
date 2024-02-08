package main

import (
	"log"
	"net/http"

	"payment-platform-solid/internal/store"
	"payment-platform-solid/pkg/api"
	"payment-platform-solid/pkg/auth"
	"payment-platform-solid/pkg/bank"
	"payment-platform-solid/pkg/middleware"
	"payment-platform-solid/pkg/payment"

	"payment-platform-solid/internal/model"
)

// AuthService struct for handling JWT token generation
type AuthService struct {
	SecretKey string
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(secretKey string) *AuthService {
	return &AuthService{SecretKey: secretKey}
}

// ProtectedHandler is an example of a protected route
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access to protected route successful"))
}

// AuditLog records an entry to the audit trail
func AuditLog(entry model.AuditLogEntry) {
	// Example: log to console, but you can modify to log to a file, database, etc.
	log.Printf("Audit Trail: Timestamp: %v, UserID: %v, Action: %v, Description: %v",
		entry.Timestamp, entry.UserID, entry.Action, entry.Description)
}

func main() {
	secretKey := "huASdHobtia4g96x"
	authService := NewAuthService(secretKey)
	jwtMiddleware := middleware.NewJWTMiddleware(secretKey)

	// Initialize the storage layer, bank simulator, and payment service
	store := store.NewStore()
	bankSimulator := bank.BankSimulator{}
	paymentService := payment.NewService(store, bankSimulator, AuditLog)
	paymentHandler := api.PaymentHandler{Service: paymentService}

	// Configure all routes, applying JWT middleware to protected endpoints
	http.HandleFunc("/login", auth.LoginHandler((*auth.AuthService)(authService)))
	http.Handle("/protected-route", jwtMiddleware.Handler(http.HandlerFunc(ProtectedHandler)))
	http.Handle("/process-payment", jwtMiddleware.Handler(http.HandlerFunc(paymentHandler.ProcessPayment)))
	http.Handle("/query-payment", jwtMiddleware.Handler(http.HandlerFunc(paymentHandler.QueryPaymentDetails)))
	http.Handle("/process-refund", jwtMiddleware.Handler(http.HandlerFunc(paymentHandler.ProcessRefund)))

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
