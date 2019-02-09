package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

// GetEnvValue get .env value via key
func GetEnvValue(key string) string {
	err := godotenv.Load()

	Error(err)

	return os.Getenv(key)
}
