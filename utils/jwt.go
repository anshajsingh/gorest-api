package utils

import (
	"errors"
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

func VerfiyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(signedSecretKey), nil
	})
	if err != nil {
		return errors.New("cannot parse token")
	}

	IsValidToken := parsedToken.Valid

	if !IsValidToken {
		return errors.New("invalid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("cannot parse claims")
	// }

	//email, userId := claims["email"].(string), claims["id"].(int64)

	//log.Println("Email from token:", email)
	//log.Println("User ID from token:", userId)

	return nil
}
