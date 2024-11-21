package conf

import (
	"os"
)

type Config struct {
	ApiGatewayUrl string
	UserService   struct {
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
		ApiGatewayUrl: os.Getenv("RYG_API_GATEWAY_URL"),
		UserService: struct {
			Url string
		}{
			Url: os.Getenv("RYG_USER_SERVICE_URL"),
		},
		TaskService: struct {
			Url string
		}{
			Url: os.Getenv("RYG_TASK_SERVICE_URL"),
		},
		ChallengeService: struct {
			Url string
		}{
			Url: os.Getenv("RYG_TASK_SERVICE_URL"),
		},
	}
}
