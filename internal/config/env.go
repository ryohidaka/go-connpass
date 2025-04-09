package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetAPIKey() string {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Println("API_KEY is not set")
	}
	return apiKey
}
