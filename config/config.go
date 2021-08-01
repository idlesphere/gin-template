package config

import (
	"os"

	"github.com/joho/godotenv"
)

// default get from env, then .env
func Get(key string) string {
	_ = godotenv.Load("config/.env")

	return os.Getenv(key)
}
