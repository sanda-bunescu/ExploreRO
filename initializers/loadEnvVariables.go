package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFiles() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
