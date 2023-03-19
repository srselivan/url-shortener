package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP
		Postgres
	}

	HTTP struct {
		Port string `env-required:"true" env:"HTTP_PORT" env-default:"8080"`
	}

	Postgres struct {
		URL         string `env-required:"true" env:"POSTGRES_URL"`
		UsePostgres bool   `env:"USE_POSTGRES" env-default:"false"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
