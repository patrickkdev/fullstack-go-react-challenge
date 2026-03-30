package configs

import (
	"os"

	"api/internal/infrastructure/db"

	"github.com/joho/godotenv"
)

var (
	Environment   string
	IsDevelopment bool
	DBConfig      db.ConnectionConfig
)

func init() {
	_ = godotenv.Load()

	Environment = getEnv("APP_ENV", "development")
	IsDevelopment = Environment != "production"

	DBConfig = db.ConnectionConfig{
		DBName: getEnv("DB_NAME", "recruiting"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", "postgres"),
		DBPort: getEnv("DB_PORT", "5432"),
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
