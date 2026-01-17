package config

import "os"

type Config struct {
	Port         string
	GeminiAPIKey string
}

func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
