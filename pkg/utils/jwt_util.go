package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"web-standalone-template/pkg/configs"
)

func Verify(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idStr, ok := claims["userID"].(string)
		if !ok {
			return 0, jwt.ErrInvalidKeyType
		}
		id, err := strconv.Atoi(idStr)
		return id, err
	}
	return 0, jwt.ErrTokenUnverifiable
}

func GenerateToken(userID int) (t string, err error) {
	claims := jwt.MapClaims{}
	claims["userID"] = strconv.Itoa(userID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err = token.SignedString([]byte(configs.JWTSecretKey))
	return
}

func keyFunc(token *jwt.Token) (any, error) {
	return []byte(configs.JWTSecretKey), nil
}
