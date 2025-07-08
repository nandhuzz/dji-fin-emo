package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DatabaseURL  string
	JWTSecretKey string
}

func Load() *Config {
	_ = godotenv.Load(".env") // load from .env if exists

	cfg := &Config{
		Port:         getEnv("PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		JWTSecretKey: getEnv("JWT_SECRET_KEY", "supersecret"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
