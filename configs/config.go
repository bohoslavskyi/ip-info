package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const envFilePath = ".env"

type Config struct {
	ServerPort int `envconfig:"SERVER_PORT" default:"80"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(envFilePath); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
