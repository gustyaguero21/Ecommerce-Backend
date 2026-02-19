package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("error loading .env file")
	}
	return &Config{
		HTTPConfig: HTTPConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
		DB: DBConfig{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
	}, nil
}
