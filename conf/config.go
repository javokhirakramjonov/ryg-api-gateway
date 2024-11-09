package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ApiGatewayPort string
	UserService    struct {
		Url string
	}
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		ApiGatewayPort: os.Getenv("API_GATEWAY_PORT"),
		UserService: struct {
			Url string
		}{
			Url: os.Getenv("USER_SERVICE_URL"),
		},
	}
}
