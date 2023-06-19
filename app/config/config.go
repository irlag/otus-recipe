package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	BindAddress string `envconfig:"optional"`
	Port        string `envconfig:"PORT"`
	DB          *DBConfig
	Migrate     Migrate `envconfig:"MIGRATE"`
	Debug       bool    `envconfig:"DEBUG"`
}

func NewConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")

	if env != "" {
		godotenv.Load(".env." + env + ".local")
	}

	if "test" != env {
		godotenv.Load(".env.local")
	}
	if env != "" {
		godotenv.Load(".env." + env)
	}
	godotenv.Load() // The Original .env

	config := &Config{
		Port:  "8000",
		Debug: false,
	}

	if err := envconfig.Init(config); err != nil {
		return nil, err
	}

	return config, nil
}
