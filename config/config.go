package config

import (
	"fmt"
	"net/url"
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

	// Read individual database environment variables
	dbHost, err := mustEnv("DATABASE_HOST")
	if err != nil {
		return nil, err
	}
	dbPort, err := mustEnv("DATABASE_PORT")
	if err != nil {
		return nil, err
	}
	dbName, err := mustEnv("DATABASE_NAME")
	if err != nil {
		return nil, err
	}
	dbUser, err := mustEnv("DATABASE_USER")
	if err != nil {
		return nil, err
	}
	dbPassword, err := mustEnv("DATABASE_PASSWORD")
	if err != nil {
		return nil, err
	}

	// Construct the DatabaseURL
	// The format for PostgreSQL is:
	// postgresql://user:password@host:port/dbname?sslmode=disable
	// For other databases, the scheme and parameters might differ.
	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		url.QueryEscape(dbUser), // URL-encode user and password
		url.QueryEscape(dbPassword),
		dbHost,
		dbPort,
		dbName,
	)

	model := getEnvWithDefault("GEMINI_MODEL", "gemini-1.5-flash")

	cfg := &Config{
		GeminiModel:  model,
		GeminiAPIKey: apiKey,
		DatabaseURL:  databaseURL,
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
