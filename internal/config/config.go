package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env  string
	Port string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		Env:  getEnv("APP_ENV", "development"),
		Port: getEnv("APP_PORT", "8080"),
	}

	log.Printf("Environment: %s", cfg.Env)
	return cfg
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
