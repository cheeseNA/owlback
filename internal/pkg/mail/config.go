package mail

import (
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	ResendApiKey string `split_words:"true" validate:"required"`
}

func loadConfig() (*config, error) {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
