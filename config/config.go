package config

import (
	"os"

	"github.com/joho/godotenv"
)

// default get from env, then .env
func Get(key string) string {
	_ = godotenv.Load()

	return os.Getenv(key)
}
