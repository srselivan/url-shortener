package config

import (
	"flag"
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
		User        string `env:"PG_USER" env-default:"postgres"`
		Password    string `env:"PG_PASSWORD"`
		Host        string `env:"PG_HOST" env-default:"localhost"`
		Port        string `env:"PG_PORT" env-default:"5432"`
		DBName      string `env:"PG_DB" env-default:"postgres"`
		SSLMode     string `env:"PG_SSLMODE" env-default:"disable"`
		UsePostgres bool   `env:"USE_POSTGRES" env-default:"false"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	flag.BoolVar(&cfg.Postgres.UsePostgres, "d", false, "use postgres")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
