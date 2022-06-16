package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type config struct {
	ExchangeURL   string `env:"EXCHANGE_URL"`
	ExchangeToken string `env:"EXCHANGE_TOKEN"`
	TelegramToken string `env:"TELEGRAM_TOKEN"`
	Debug         bool   `env:"DEBUG"`
}

func New() config {
	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Println("%+v\n", err)
	}

	return cfg
}
