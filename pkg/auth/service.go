package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthService struct to handle authentication logic
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
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns its claims
func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return token, nil
}
