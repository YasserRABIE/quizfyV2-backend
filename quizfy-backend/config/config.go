package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
)

func InitConfig() {
	DBUser = getEnv("DB_USER", "postgres")
	DBPassword = getEnv("DB_PASSWORD", "password")
	DBName = getEnv("DB_NAME", "postgres")
	DBHost = getEnv("DB_HOST", "localhost")
	DBPort = getEnv("DB_PORT", "5432")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
