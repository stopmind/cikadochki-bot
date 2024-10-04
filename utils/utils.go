package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Utils struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) Utils {
	return Utils{
		bot: bot,
	}
}
