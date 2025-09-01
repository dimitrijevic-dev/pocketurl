package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	godotenv.Load()
	return os.Getenv(key)
}