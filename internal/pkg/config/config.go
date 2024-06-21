package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type RunningEnvironment string

const (
	Local      RunningEnvironment = "local"
	Production RunningEnvironment = "production"
)

type Config struct {
	PostgresConnectionString string             `split_words:"true" validate:"required"`
	RunningEnvironment       RunningEnvironment `split_words:"true" validate:"required"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
