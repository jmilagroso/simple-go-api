// env.go
// Environment key/val
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 31 2019

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
