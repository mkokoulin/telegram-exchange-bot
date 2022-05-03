package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type config struct {
	BaseURL string `env:"BASE_URL"`
	ApiKey  string `env:"API_KEY"`
}

func New() config {
	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Println("%+v\n", err)
	}

	return cfg
}
