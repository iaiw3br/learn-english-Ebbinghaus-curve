package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"tg-bot-learning-english/internal/config"
)

func Run(cfg *config.Config) error {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "Create set":
			msg.Text = "Write a new name of set:"
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		default:

		}

		if _, err := bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}
