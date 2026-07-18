package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	RedisAddr      string
}

func LeadConfig() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No env file found")
		log.Println("Error loading .env:", err)
	}

	//loads .env and validate
	required := []string{
		"APP_PORT",
		"REDIS_ADDR",
		"BREVO_FROM_EMAIL",
		"BREVO_FROM_NAME",
		"BREVO_API_KEY",
	}
	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Missing required env variable: %s", key)
		}
	}
	//return the hole struct
	return &Config{
		AppPort:   os.Getenv("APP_PORT"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
	}
}
