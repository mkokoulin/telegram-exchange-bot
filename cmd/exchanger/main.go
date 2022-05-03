package main

import (
	"fmt"

	"github.com/KokoulinM/telegram-exchange-bot/cmd/exchanger/configs"
	"github.com/KokoulinM/telegram-exchange-bot/internal/app/client"
)

func main() {
	cfg := configs.New()

	c := client.New(cfg.BaseURL, cfg.ApiKey)

	data, err := c.Convert("AMD", "RUB", "5000")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.Result)
}
