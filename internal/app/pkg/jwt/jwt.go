package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateToken(email, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString, secret string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
