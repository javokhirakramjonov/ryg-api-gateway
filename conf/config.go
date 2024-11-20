package conf

import (
	"os"
)

type Config struct {
	ApiGatewayPort string
	UserService    struct {
		Url string
	}
	TaskService struct {
		Url string
	}
	ChallengeService struct {
		Url string
	}
}

func LoadConfig() *Config {
	return &Config{
		ApiGatewayPort: os.Getenv("API_GATEWAY_PORT"),
		UserService: struct {
			Url string
		}{
			Url: os.Getenv("USER_SERVICE_URL"),
		},
		TaskService: struct {
			Url string
		}{
			Url: os.Getenv("TASK_SERVICE_URL"),
		},
		ChallengeService: struct {
			Url string
		}{
			Url: os.Getenv("CHALLENGE_SERVICE_URL"),
		},
	}
}
