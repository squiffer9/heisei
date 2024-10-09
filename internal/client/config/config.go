package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server     ServerConfig
	UI         UIConfig
	Connection ConnectionConfig
}

type ServerConfig struct {
	URL string
}

type UIConfig struct {
	Language       string
	RefreshRate    time.Duration
	MaxThreadsShow int
}

type ConnectionConfig struct {
	Timeout       time.Duration
	RetryAttempts int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		// It's okay if .env file is missing, we'll use default values
	}

	return &Config{
		Server: ServerConfig{
			URL: getEnv("SERVER_URL", "http://localhost:8080"),
		},
		UI: UIConfig{
			Language:       getEnv("UI_LANGUAGE", "en"),
			RefreshRate:    time.Duration(getEnvAsInt("UI_REFRESH_RATE", 5)) * time.Second,
			MaxThreadsShow: getEnvAsInt("UI_MAX_THREADS", 20),
		},
		Connection: ConnectionConfig{
			Timeout:       time.Duration(getEnvAsInt("CONNECTION_TIMEOUT", 10)) * time.Second,
			RetryAttempts: getEnvAsInt("CONNECTION_RETRY_ATTEMPTS", 3),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
