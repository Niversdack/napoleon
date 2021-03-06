package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       uint
	Username string
	jwt.StandardClaims
}

func GenerateJWT(secretKey string, username string, id uint) string {
	var jwtKey = []byte(secretKey)

	claims := &Claims{
		Username:       username,
		ID:             id,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return ""
	}

	return tokenString

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
