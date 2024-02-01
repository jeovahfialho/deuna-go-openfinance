package middleware

import (
	"context" // Necessary for using context
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JWTMiddleware structure to hold the secret key for JWT validation
type JWTMiddleware struct {
	SecretKey string
}

// NewJWTMiddleware creates a new instance of JWTMiddleware
func NewJWTMiddleware(secretKey string) *JWTMiddleware {
	return &JWTMiddleware{SecretKey: secretKey}
}

// Handler is the middleware function to handle JWT validation
func (m *JWTMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		if tokenString == "" {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(m.SecretKey), nil
		})

		// Check for parsing errors
		if err != nil {
			http.Error(w, fmt.Sprintf("Unauthorized: %v", err), http.StatusUnauthorized)
			return
		}

		// Check if the token claims are valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, ok := claims["user_id"].(string)
			if !ok {
				http.Error(w, "Unauthorized: Invalid user_id in token", http.StatusUnauthorized)
				return
			}

			// Add userID to the request context
			ctx := context.WithValue(r.Context(), "userID", userID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		}
	})
}

// extractToken extracts the token from the Authorization header
func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
