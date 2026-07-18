package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	DBUrl          string
	RedisAddr      string
	JWTSecret      string
	RazorpayKey    string
	RazorpaySecret string
}

func LeadConfig() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No env file found")
		log.Println("Error loading .env:", err)
	}

	redisUrl :=  os.Getenv("REDIS_ADDR")
	//production case (neon)
	if redisUrl ==""{

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
	}
	//return the hole struct
	return &Config{
		AppPort:   os.Getenv("APP_PORT"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
	}
}
