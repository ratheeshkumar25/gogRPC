package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	UserID uint
	Phone  string `json:"username"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
