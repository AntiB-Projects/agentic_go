package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GeminiModel  string
	GeminiAPIKey string
	DatabaseURL  string
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf(
			"error loading .env file: %w; make sure you copy .env.example to .env",
			err,
		)
	}

	apiKey, err := mustEnv("GEMINI_API_KEY")
	if err != nil {
		return nil, err
	}

	dbURL, err := mustEnv("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	model := getEnvWithDefault("GEMINI_MODEL", "gemini-1.5-flash")

	cfg := &Config{
		GeminiModel:  model,
		GeminiAPIKey: apiKey,
		DatabaseURL:  dbURL,
	}
	return cfg, nil
}

func mustEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("required env var %q not set", key)
	}
	return val, nil
}

func getEnvWithDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}
