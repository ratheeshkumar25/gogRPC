package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	UserID uint
	Phone  string `json:"phone"`
	jwt.StandardClaims
}
