package utils

import (
	"headliner-be/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(config.Jwt_secret)

func CreateToken(id uint, email string, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"email":    email,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
