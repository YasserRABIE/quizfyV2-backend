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

	BasePath string
)

func InitConfig() {
	DBUser = GetEnv("DB_USER", "postgres")
	DBPassword = GetEnv("DB_PASSWORD", "password")
	DBName = GetEnv("DB_NAME", "postgres")
	DBHost = GetEnv("DB_HOST", "localhost")
	DBPort = GetEnv("DB_PORT", "5432")

	BasePath = GetEnv("QUIZ_UPLOAD_PATH", "uploads/quiz/")
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
