package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID       string
	Username string
	jwt.StandardClaims
}

func GenerateJWT(secretKey string, expiration time.Duration, username string) string {
	var jwtKey = []byte(secretKey)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
		},
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
