package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️ .env не найден, читаю переменные окружения напрямую")
	}

	cfg := &Config{}

	// Database configuration
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "postgres")
	cfg.Database.Password = getEnv("DB_PASSWORD", "postgres")
	cfg.Database.Name = getEnv("DB_NAME", "lms_system")
	cfg.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	// Server configuration
	cfg.Server.Port = getEnv("PORT", "8080")

	// Logger configuration
	cfg.Logger.Level = getEnv("LOG_LEVEL", "info")
	cfg.Logger.Format = getEnv("LOG_FORMAT", "json")

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
