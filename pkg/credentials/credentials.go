package credentials

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetCredentials() Credentials {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Credentials{
		X_API_TOKEN: os.Getenv("X_API_TOKEN"),
		From:        os.Getenv("FROM"),
		DB:          os.Getenv("DB"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASS:     os.Getenv("DB_PASS"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_DIALECT:  os.Getenv("DB_DIALECT"),
	}
}
