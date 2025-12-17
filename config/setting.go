package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Port string

func init() {
	configpath := Initenv()

	err := godotenv.Load(configpath)
	if err != nil {
		log.Fatalf("Problem loading .env file: %v", err)
		os.Exit(-1)
	}
	Port = os.Getenv("PORT")
}
