package config

import (
	"errors"
	"os"
)

type Config struct {
	AppEnv      string
	Port        string
	DatabaseURL string
	JwtSecret   string
}

func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("DATABASE_URL is required")
	}

	return &Config{
		AppEnv:      os.Getenv("APP_ENV"),
		Port:        port,
		DatabaseURL: dbURL,
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}, nil
}
