package token

import (
	"errors"
	"fmt"
	"os"

	//"os/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/model"
)

// GenerateToken generates a token for 5 hours with given data
func GenerateToken(phone string, userid uint) (string, error) {
	// Get the secret key from environment variable
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET_KEY not set in environment")
	}

	// Create the claims
	claims := &model.UserClaims{
		UserID: userid,
		Phone:  phone,
		Role:   "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(), // Token expires after 5 hours
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}
