package utils

import (
	"fmt"
	"headliner-be/config"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func IsExist(c fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	parts := strings.Split(auth, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        return fiber.NewError(fiber.StatusUnauthorized, "invalid token format")
    }
    tokenString := parts[1]
	
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.Jwt_secret), nil
    })
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	claim := token.Claims.(jwt.MapClaims)
	if claim["exp"] == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	c.Locals("userID", claim["id"])
	c.Locals("email", claim["email"])
	fmt.Println("User ID:", claim["id"], "Email:", claim["email"])
	return c.Next()
}
