package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int
	Env  string
}

func LoadConfig() *Config {
	// Load .env file (only in local)
	_ = godotenv.Load()

	portStr := getEnv("PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid PORT value: %v", err)
	}

	env := getEnv("ENV", "local")

	return &Config{
		Port: port,
		Env:  env,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
