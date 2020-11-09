package handlers

import (
	"github.com/dgrijalva/jwt-go"
)

type tokenField string

var userID tokenField = "user_id"

func GenerateJWT(uid, verifyKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		string(userID): uid,
	})

	tokenString, err := token.SignedString([]byte(verifyKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
