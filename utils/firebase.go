package utils

import (
	"context"
	"log"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v3"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client

func InitFirebase() error {
	opt := option.WithCredentialsFile("./etc/secrets/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return err
	}

	AuthClient = client
	return nil
}

func FirebaseAuth(c fiber.Ctx) error {
	// 1. ดึง Header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Authorization header is required")
	}

	// 2. เช็ค Bearer Format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
	}

	// 3. ตรวจสอบกับ Firebase Singleton
	if AuthClient == nil {
    	log.Println("ERROR: Firebase AuthClient is not initialized")
    	return fiber.NewError(fiber.StatusInternalServerError, "Internal server configuration error")
	}

	token, err := AuthClient.VerifyIDToken(context.Background(), parts[1])
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}
	// 4. ฝากข้อมูลไว้ใน Locals
	c.Locals("GOOGLE_AUTH_TOKEN", token)
	
	return c.Next()
}