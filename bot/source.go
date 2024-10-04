package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Source interface {
	GetImage() (tgbotapi.RequestFileData, error)
}
