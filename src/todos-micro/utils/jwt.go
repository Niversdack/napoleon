package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       uint
	Username string
	jwt.StandardClaims
}

func CheckJwtValid(secretKey string, accessToken string) bool {
	var jwtKey = []byte(secretKey)

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false
	}

	if tkn.Valid {
		return true
	} else {
		return false
	}
}
func GetClaims(secretKey string, accessToken string) *Claims {
	var jwtKey = []byte(secretKey)

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return &Claims{}
	}

	if tkn.Valid {
		return claims
	}
	return &Claims{}
}
