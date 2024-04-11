package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env, err: %v", err)
	}
}
