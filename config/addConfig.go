package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func SetupEnv() (cfg Config, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("Port")
	if len(httpPort) < 1 {
		return Config{}, errors.New("environment variable not found")
	}
	return Config{Port: httpPort}, nil

}
