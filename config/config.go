package config

import (
	"github.com/joho/godotenv"
	"log"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
