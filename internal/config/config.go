package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WebhookURL     string
	RedisHost      string
	RedisPort      string
	PostgresHost   string
	PostgresPort   string
	PostgresUser   string
	PostgresPass   string
	PostgresDBName string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	cfg := &Config{
		WebhookURL:     os.Getenv("WEBHOOK_URL"),
		RedisHost:      os.Getenv("REDIS_HOST"),
		RedisPort:      os.Getenv("REDIS_PORT"),
		PostgresHost:   os.Getenv("POSTGRES_HOST"),
		PostgresPort:   os.Getenv("POSTGRES_PORT"),
		PostgresUser:   os.Getenv("POSTGRES_USER"),
		PostgresPass:   os.Getenv("POSTGRES_PASSWORD"),
		PostgresDBName: os.Getenv("POSTGRES_DB"),
	}
	if cfg.WebhookURL == "" || cfg.RedisHost == "" || cfg.PostgresHost == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}
