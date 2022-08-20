package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"web-standalone-template/pkg/configs"
)

func Verify(signedToken string) (int, error) {
	token, err := jwt.Parse(signedToken, keyFunc)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	id, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("invalid userID")
	}

	return int(id), nil
}

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.JWTSecretKey))
}

func keyFunc(token *jwt.Token) (any, error) {
	return []byte(configs.JWTSecretKey), nil
}
