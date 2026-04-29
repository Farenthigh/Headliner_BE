package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Jwt_secret string
)

func InitAuth() {
	configpath := Initenv()

	err := godotenv.Load(configpath)
	if err != nil {
		log.Fatalf("Problem loading .env file: %v", err)
		os.Exit(-1)
	}
	Jwt_secret = os.Getenv("JWT_SECRET")
}
