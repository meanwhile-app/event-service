package utilities

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = []byte("meanwhile-jwt-secret")

func DecodeJwt(token string) (jwt.MapClaims, error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok && tokenObj.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func EncodeJwt(clams *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)

	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
