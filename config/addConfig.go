package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	Dsn              string //data seouce name
	AppSecret        string
	TwilioAccountSid string
	TwilioAuthToken  string
	TwilioFromPhone  string
}

func SetupEnv() (cfg Config, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("Port")

	if len(httpPort) < 1 {
		return Config{}, errors.New("environment variable port not found")
	}

	Dsn := os.Getenv("DSN")

	if len(Dsn) < 1 {
		return Config{}, errors.New("environment variable dsn not found")
	}

	appSercet := os.Getenv("APP_SECRET")
	if len(appSercet) < 1 {
		return Config{}, errors.New("environment variable appsecret not found")
	}

	return Config{Port: httpPort, Dsn: Dsn, AppSecret: appSercet,
		TwilioAccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioFromPhone:  os.Getenv("TWILIO_FROM_PHONE"),
	}, nil

}
