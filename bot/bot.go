package bot

import (
	"cikadochki-bot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"slices"
)

type Bot struct {
	utils.Utils
	bot              *tgbotapi.BotAPI
	updateConfig     tgbotapi.UpdateConfig
	commandsHandlers map[string]func(message *tgbotapi.Message)
	data             data
	dataPath         string
	Sources          []Source
}

func (b *Bot) Run() {
	updates := b.bot.GetUpdatesChan(b.updateConfig)

	for update := range updates {

		if update.MyChatMember != nil {
			if update.MyChatMember.NewChatMember.CanPostMessages {
				if !slices.Contains(b.data.Channels, update.MyChatMember.Chat.ID) {
					b.data.Channels = append(b.data.Channels, update.MyChatMember.Chat.ID)
				}
			} else {
				index := slices.Index(b.data.Channels, update.MyChatMember.Chat.ID)
				if index != -1 {
					b.data.Channels = slices.Delete(b.data.Channels, index, index)
				}
			}
			continue
		}
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		handler, ok := b.commandsHandlers[update.Message.Command()]
		if ok {
			handler(update.Message)
		}
	}
}

func (b *Bot) Stop() error {
	if err := b.data.write(b.dataPath); err != nil {
		return err
	}

	return nil
}

func (b *Bot) SendImages() error {
	for _, channel := range b.data.Channels {
		source := b.Sources[rand.Intn(len(b.Sources))]
		image, err := source.GetImage()
		if err != nil {
			return err
		}

		media := tgbotapi.NewInputMediaPhoto(image)
		media.Caption = "Цикадочки"

		_, err = b.bot.SendMediaGroup(tgbotapi.NewMediaGroup(channel, []interface{}{media}))
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) onStart(message *tgbotapi.Message) {
	if !slices.Contains(b.data.Channels, message.Chat.ID) {
		b.data.Channels = append(b.data.Channels, message.Chat.ID)
	}

}

func (b *Bot) onSend(message *tgbotapi.Message) {
	err := b.SendImages()
	if err != nil {
		panic(err)
	}
}

func New(token string, dataPath string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	d, _ := tryReadData(dataPath)

	result := &Bot{
		bot:              bot,
		updateConfig:     updateConfig,
		dataPath:         dataPath,
		data:             d,
		commandsHandlers: make(map[string]func(message *tgbotapi.Message)),
	}

	result.commandsHandlers["start"] = result.onStart
	result.commandsHandlers["send"] = result.onSend

	return result, nil
}
