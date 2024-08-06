package token

import (
	"errors"
	"fmt"
	"os"
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
		Phone:  phone,
		UserID: userid,
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

// func ValidateToken(tokenString string) (*model.UserClaims, error) {
// 	// Get the secret key from environment variable
// 	secretKey := os.Getenv("JWT_SECRET_KEY")
// 	if secretKey == "" {
// 		return nil, errors.New("JWT_SECRET_KEY not set in environment")
// 	}

// 	// Parse the token
// 	token, err := jwt.ParseWithClaims(tokenString, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(secretKey), nil
// 	})

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse token: %w", err)
// 	}

// 	// Validate the token and return the claims
// 	if claims, ok := token.Claims.(*model.UserClaims); ok && token.Valid {
// 		return claims, nil
// 	}

// 	return nil, errors.New("invalid token")
// }
