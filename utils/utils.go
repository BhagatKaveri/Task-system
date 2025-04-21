package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("diputskoolima")

func GenerateToken(role string, UserID string, email string) (token string, err error) {
	tokenExpirationTime := time.Now().Add(time.Hour * 6)
	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": UserID,
		"email":   email,
		"exp":     tokenExpirationTime.Unix(),
	})
	token, err = tokenObject.SignedString(SecretKey)
	return
}
