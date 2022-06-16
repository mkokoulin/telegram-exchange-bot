package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/KokoulinM/telegram-exchange-bot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/KokoulinM/telegram-exchange-bot/internal/client"
	"github.com/KokoulinM/telegram-exchange-bot/internal/configs"
)

func main() {
	сurrency := models.Currency{
		From:   "",
		To:     "",
		Amount: "",
	}

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

		switch command[0] {
		case "FROM":
			сurrency.From = command[1]

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "The currency FROM is set: "+command[1])
			bot.Send(msg)

			continue
		case "TO":
			сurrency.To = command[1]

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "The currency TO is set: "+command[1])
			bot.Send(msg)

			continue
		case "CLEAR":
			сurrency.From = ""
			сurrency.To = ""

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "The FROM and TO values have been cleared")
			bot.Send(msg)

			continue
		default:
			if len(command) == 1 {
				сurrency.Amount = command[0]
			}

			if len(command) > 2 {
				сurrency.From = command[0]
				сurrency.To = command[1]
				сurrency.Amount = command[2]
			}
		}

		cr := client.ConvertRequest{
			From:   сurrency.From,
			To:     сurrency.To,
			Amount: сurrency.Amount,
		}

		data, err := c.Convert(cr)
		if err != nil {
			var convertErr *client.ErrorWithConvert

			if errors.As(err, &convertErr) && err.(*client.ErrorWithConvert).Title == "InvalidRequestFormat" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid request format. The correct format is FROM TO AMOUNT -> GBP USD 3000. Or you need to set the values FROM and TO")
				bot.Send(msg)
				continue
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%f", data.Result))
		bot.Send(msg)
	}
}
