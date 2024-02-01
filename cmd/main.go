package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"payment-platform-solid/internal/store"
	"payment-platform-solid/pkg/api"
	"payment-platform-solid/pkg/bank"
	"payment-platform-solid/pkg/middleware"
	"payment-platform-solid/pkg/payment"

	"github.com/dgrijalva/jwt-go"
)

// AuthService struct for handling JWT token generation
type AuthService struct {
	SecretKey string
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(secretKey string) *AuthService {
	return &AuthService{SecretKey: secretKey}
}

// GenerateToken generates a JWT token for a given user ID
func (s *AuthService) GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     jwt.TimeFunc().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(s.SecretKey))
}

// JWTMiddleware to protect routes
func JWTMiddleware(next http.Handler, secretKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// LoginHandler to authenticate users and issue JWT tokens
func LoginHandler(authService *AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := "testuser" // Simulated user ID
		token, err := authService.GenerateToken(userID)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	}
}

// ProtectedHandler is an example of a protected route
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access to protected route successful"))
}

func main() {
	secretKey := "huASdHobtia4g96x"
	authService := NewAuthService(secretKey)
	jwtMiddleware := middleware.NewJWTMiddleware(secretKey)

	// Initialize the storage layer, bank simulator, and payment service
	store := store.NewStore()
	bankSimulator := bank.BankSimulator{}
	paymentService := payment.NewService(store, bankSimulator)
	paymentHandler := api.PaymentHandler{Service: paymentService}

	// Configure all routes, applying JWT middleware to protected endpoints
	http.HandleFunc("/login", LoginHandler(authService))
	http.Handle("/protected-route", jwtMiddleware.Handler(http.HandlerFunc(ProtectedHandler)))
	http.Handle("/process-payment", jwtMiddleware.Handler(http.HandlerFunc(paymentHandler.ProcessPayment)))
	http.Handle("/query-payment", jwtMiddleware.Handler(http.HandlerFunc(paymentHandler.QueryPaymentDetails)))
	http.Handle("/process-refund", jwtMiddleware.Handler(http.HandlerFunc(paymentHandler.ProcessRefund)))

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
