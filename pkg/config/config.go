package config

import (
	"os"
)

// ? Config holds all configuration for our program
type Config struct {
	Port string
}

// ? Load returns a Config struct with values loaded from the environment
func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
	}
}

// ? Simple helper function to read an environment variable or return a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
