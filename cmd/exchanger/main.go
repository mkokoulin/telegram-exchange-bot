package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/KokoulinM/telegram-exchange-bot/internal/client"
	"github.com/KokoulinM/telegram-exchange-bot/internal/configs"
)

func main() {
	cfg := configs.New()

	c := client.New(cfg.ExchangeURL, cfg.ExchangeToken)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		command := strings.Split(update.Message.Text, " ")

		сurrencyFrom := command[0]
		сurrencyTo := command[1]
		amount := command[2]

		cr := client.ConvertRequest{
			From:   сurrencyFrom,
			To:     сurrencyTo,
			Amount: amount,
		}

		data, err := c.Convert(cr)
		if err != nil {
			var convertErr *client.ErrorWithConvert

			if errors.As(err, &convertErr) && err.(*client.ErrorWithConvert).Title == "InvalidRequestFormat" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid request format. The correct format is FROM TO AMOUNT -> GBP USD 3000")
				bot.Send(msg)
				continue
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%f", data.Result))
		bot.Send(msg)
	}
}
