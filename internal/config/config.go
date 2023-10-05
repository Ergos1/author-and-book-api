package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func NewConfig() *Config {
	return &Config{
		Database: PostgresConfig{
			User:     getEnv("POSTGRES_USER", "test"),
			Password: getEnv("POSTGRES_PASSWORD", "test"),
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnvAsInt("POSTGRES_PORT", 5433),
			Dbname:   getEnv("POSTGRES_DB", "test"),
		},
		Server: ServerConfig{
			Address: getEnv("SERVER_ADDRESS", ":9000"),
		},
	}
}
