package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const signedSecretKey = "supersecret"

func GenerateJWTToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    userId,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(signedSecretKey))
}
